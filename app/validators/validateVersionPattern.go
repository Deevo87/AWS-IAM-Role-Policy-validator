package validators

import (
	"errors"
	"regexp"
	"zadanie_remitly/app/structures"
)

type ValidateVersionPattern struct {
	next Validator
}

func (v *ValidateVersionPattern) SetNext(validator Validator) {
	v.next = validator
}

func (v *ValidateVersionPattern) Execute(policy structures.Policy, ind int) (bool, error) {
	pattern := `^\d{4}-\d{2}-\d{2}$`
	if ok, _ := regexp.MatchString(pattern, policy.PolicyDocument.Version); ok {
		return checkNext(v.next, policy, ind)
	}
	return false, errors.New("couldn't match version pattern, it should like this 'yyyy-mm-dd'")
}
