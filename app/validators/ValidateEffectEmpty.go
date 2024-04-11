package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidateEffectEmpty struct {
	next Validator
}

func (v *ValidateEffectEmpty) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateEffectEmpty) Execute(policy structures.Policy, ind int) (bool, error) {
	effect := policy.PolicyDocument.Statement[ind].Effect
	if effect == "" {
		return false, errors.New("effect is required, cannot be empty string")
	}
	return checkNext(v.next, policy, ind)
}
