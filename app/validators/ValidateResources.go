package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidateResources struct {
	next Validator
}

func (v *ValidateResources) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateResources) Execute(policy structures.Policy, ind int) (bool, error) {
	resources := policy.PolicyDocument.Statement[ind].Resource
	if res, ok := resources.([]interface{}); ok {
		for _, resource := range res {
			if resStr, ok := resource.(string); ok {
				result, err := v.checkForAsterisk(resStr)
				if !result {
					return result, err
				}
			} else {
				return false, errors.New("json file is not valid, elements in resources array are not strings")
			}
		}
	} else if resStr, ok := resources.(string); ok {
		result, err := v.checkForAsterisk(resStr)
		if !result {
			return false, err
		}
	} else {
		return false, errors.New("resources are nor string nor string array")
	}
	return checkNext(v.next, policy, ind)
}

func (v *ValidateResources) checkForAsterisk(element string) (bool, error) {
	for i := 0; i < len(element); i++ {
		if element[i] == '*' {
			return false, errors.New("json file is not valid, asterisk occurred")
		}
	}
	return true, nil
}
