package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidateActionEmpty struct {
	next Validator
}

func (v *ValidateActionEmpty) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateActionEmpty) Execute(policy structures.Policy, ind int) (bool, error) {
	action := policy.PolicyDocument.Statement[ind].Action
	if len(action) == 0 {
		return false, errors.New("action is required, cannot be empty")
	}
	return checkNext(v.next, policy, ind)
}
