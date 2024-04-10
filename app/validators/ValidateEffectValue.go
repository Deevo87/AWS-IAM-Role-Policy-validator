package validators

import (
	"errors"
	"zadanie_remitly/app/structures"
)

type ValidateEffectValue struct {
	next Validator
}

func (v *ValidateEffectValue) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateEffectValue) Execute(policy structures.Policy, ind int) (bool, error) {
	effect := policy.PolicyDocument.Statement[ind].Effect
	if effect != "Allow" && effect != "Deny" {
		return false, errors.New("effect's value can be only 'Allow' or 'Deny'")
	}
	return checkNext(v.next, policy, ind)
}
