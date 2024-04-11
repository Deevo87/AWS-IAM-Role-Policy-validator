package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidateVersionEmpty struct {
	next Validator
}

func (v *ValidateVersionEmpty) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateVersionEmpty) Execute(policy structures.Policy, ind int) (bool, error) {
	if policy.PolicyDocument.Version == "" {
		return false, errors.New("version is required, cannot be empty")
	}
	return checkNext(v.next, policy, ind)
}
