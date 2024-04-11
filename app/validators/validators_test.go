package validators

import (
	"errors"
	"testing"
	"zadanie_remitly/app/structures"
)

type Test struct {
	name        string
	testField   interface{}
	expectPass  bool
	expectedErr error
}

func TestValidateVersionPattern_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "Valid version pattern",
			testField:   "2024-04-11",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "Invalid version pattern",
			testField:   "2024/04/11",
			expectPass:  false,
			expectedErr: errors.New("couldn't match version pattern, it should like this 'yyyy-mm-dd'"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidateVersionPattern{}
			policy := structures.Policy{
				PolicyDocument: structures.PolicyDocument{
					Version: tc.testField.(string),
				},
			}
			checker(validator, policy, tc, t)
		})
	}
}

func TestValidatePolicyNameEmpty_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "'PolicyName' value is not empty",
			testField:   "całkiem spoko imię",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "'PolicyName' value is not empty",
			testField:   "całkiem spoko imię111122v212",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "'PolicyName' value is empty",
			testField:   "",
			expectPass:  false,
			expectedErr: errors.New("policy name is required, cannot be empty"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidatePolicyNameEmpty{}
			policy := structures.Policy{
				PolicyName: tc.testField.(string),
			}
			checker(validator, policy, tc, t)
		})
	}
}

func TestValidateActionEmpty_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "Valid action filled with normal values",
			testField:   []string{"iam:ListRoles", "iam:ListUsers"},
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "Invalid action being empty",
			testField:   []string{},
			expectPass:  false,
			expectedErr: errors.New("action is required, cannot be empty"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidateActionEmpty{}
			policy := structures.Policy{
				PolicyDocument: structures.PolicyDocument{
					Statement: []structures.Statement{
						{
							Action: tc.testField.([]string),
						},
					},
				},
			}

			checker(validator, policy, tc, t)
		})
	}
}

func TestValidateEffectEmpty_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "'Effect' is not empty",
			testField:   "sadasdwda",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "'Effect' is empty string",
			testField:   "",
			expectPass:  false,
			expectedErr: errors.New("effect is required, cannot be empty string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidateEffectEmpty{}
			policy := structures.Policy{
				PolicyDocument: structures.PolicyDocument{
					Statement: []structures.Statement{
						{
							Effect: tc.testField.(string),
						},
					},
				},
			}

			checker(validator, policy, tc, t)
		})
	}
}

func TestValidateEffectValue_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "Valid 'Effect' value",
			testField:   "Allow",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "Valid 'Effect' value",
			testField:   "Deny",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "Invalid 'Effect' value",
			testField:   "asdawdsda asdas sqeasd",
			expectPass:  false,
			expectedErr: errors.New("effect's value can be only 'Allow' or 'Deny'"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidateEffectValue{}
			policy := structures.Policy{
				PolicyDocument: structures.PolicyDocument{
					Statement: []structures.Statement{
						{
							Effect: tc.testField.(string),
						},
					},
				},
			}

			checker(validator, policy, tc, t)
		})
	}
}

func TestValidateResources_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "Valid resources with string array",
			testField:   []interface{}{"resource1", "resource2"},
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "Valid resources with string",
			testField:   "resource",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "Invalid resources with asterisk in string array",
			testField:   []interface{}{"resource1", "resource2", "resource*"},
			expectPass:  false,
			expectedErr: errors.New("asterisk occurred"),
		},
		{
			name:        "Invalid resources with asterisk in string",
			testField:   "resource*",
			expectPass:  false,
			expectedErr: errors.New("asterisk occurred"),
		},
		{
			name:        "Invalid resources with non-string array",
			testField:   []interface{}{1, 2, 3},
			expectPass:  false,
			expectedErr: errors.New("elements in resources array are not strings"),
		},
		{
			name:        "Invalid resources with non-string",
			testField:   123,
			expectPass:  false,
			expectedErr: errors.New("resources are nor string nor string array"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidateResources{}

			policy := structures.Policy{
				PolicyDocument: structures.PolicyDocument{
					Statement: []structures.Statement{
						{
							Resource: tc.testField,
						},
					},
				},
			}

			checker(validator, policy, tc, t)
		})
	}
}

func TestValidateResourcesEmpty_Execute(t *testing.T) {
	testCases := []Test{
		{
			name:        "'Resource' is not empty",
			testField:   []interface{}{"resource1", "resource2"},
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "'Resource' is not empty",
			testField:   "resource",
			expectPass:  true,
			expectedErr: nil,
		},
		{
			name:        "'Resource' is empty string",
			testField:   "",
			expectPass:  false,
			expectedErr: errors.New("resource is required, cannot be empty"),
		},
		{
			name:        "'Resource' is empty []string",
			testField:   []interface{}{},
			expectPass:  false,
			expectedErr: errors.New("resource is required, cannot be empty"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := &ValidateResourcesEmpty{}

			policy := structures.Policy{
				PolicyDocument: structures.PolicyDocument{
					Statement: []structures.Statement{
						{
							Resource: tc.testField,
						},
					},
				},
			}

			checker(validator, policy, tc, t)
		})
	}
}

func checker(validator Validator, policy structures.Policy, testCase Test, t *testing.T) {
	pass, err := validator.Execute(policy, 0)

	if pass != testCase.expectPass {
		t.Errorf("Test case %s failed: expected pass=%v, got pass=%v", testCase.name, testCase.expectPass, pass)
	}

	if (err == nil && testCase.expectedErr != nil) || (err != nil && testCase.expectedErr == nil) || (err != nil && testCase.expectedErr != nil && err.Error() != testCase.expectedErr.Error()) {
		t.Errorf("Test case %s failed: expected error=%v, got error=%v", testCase.name, testCase.expectedErr, err)
	}
}
