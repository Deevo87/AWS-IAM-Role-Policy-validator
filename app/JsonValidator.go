package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"zadanie_remitly/app/structures"
	"zadanie_remitly/app/validators"
)

type ValidatorService struct {
}

func NewValidatorService() *ValidatorService {
	return &ValidatorService{}
}

func (v *ValidatorService) Validate(file multipart.File) (bool, error) {
	var policy structures.Policy
	err := json.NewDecoder(file).Decode(&policy)
	fmt.Println(policy)
	if err != nil {
		return false, errors.New("not valid JSON structure")
	}

	// here could be function, which defines what variant of validation we want to use
	return v.setUpValidation(policy)
}

// here could be case with different validators variations and setups if needed, I didn't need so there is only one
func (v *ValidatorService) setUpValidation(policy structures.Policy) (bool, error) {
	// validation for basic fields
	startingValidators := []validators.Validator{
		&validators.ValidateVersionEmpty{},
		&validators.ValidateVersionPattern{},
		&validators.ValidatePolicyNameEmpty{},
	}
	result, err := v.link(startingValidators).Execute(policy, 0)
	if !result {
		return false, err
	}

	// validation is performed for each statement
	for ind := 0; ind < len(policy.PolicyDocument.Statement); ind++ {
		result, err := v.link([]validators.Validator{
			&validators.ValidateActionEmpty{},
			&validators.ValidateResourcesEmpty{},
			&validators.ValidateEffectEmpty{},
			&validators.ValidateResources{},
			&validators.ValidateEffectValue{},
		}).Execute(policy, ind)
		if !result {
			return false, err
		}
	}
	return true, nil
}

func (v *ValidatorService) link(validators []validators.Validator) validators.Validator {
	first := validators[0]
	head := first
	for i := 1; i < len(validators); i++ {
		head.SetNext(validators[i])
		head = validators[i]
	}
	return first
}
