package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidatePolicyNameEmpty struct {
	next Validator
}

func (v *ValidatePolicyNameEmpty) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidatePolicyNameEmpty) Execute(policy structures.Policy, ind int) (bool, error) {
	if policy.PolicyName == "" {
		return false, errors.New("policy name is required, cannot be empty")
	}
	return checkNext(v.next, policy, ind)
}
