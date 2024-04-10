package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidateResourcesEmpty struct {
	next Validator
}

func (v *ValidateResourcesEmpty) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateResourcesEmpty) Execute(policy structures.Policy, ind int) (bool, error) {
	resources := policy.PolicyDocument.Statement[ind].Resource
	if res, ok := resources.([]interface{}); ok {
		if len(res) == 0 {
			return false, errors.New("resource is required, cannot be empty")
		}
	} else if res, ok := resources.(string); ok {
		if res == "" {
			return false, errors.New("resource is required, cannot be empty")
		}
	}
	return checkNext(v.next, policy, ind)
}
