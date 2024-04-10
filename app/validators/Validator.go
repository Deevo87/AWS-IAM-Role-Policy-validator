package validators

import (
	"zadanie_remitly/app/structures"
)

type Validator interface {
	Execute(policy structures.Policy, ind int) (bool, error)
	SetNext(validator Validator)
}

func checkNext(validator Validator, policy structures.Policy, ind int) (bool, error) {
	if validator == nil {
		return true, nil
	}
	return validator.Execute(policy, ind)
}
