package structures

type Policy struct {
	PolicyName     string         `json:"PolicyName,omitempty"`
	PolicyDocument PolicyDocument `json:"PolicyDocument,omitempty"`
}

type Statement struct {
	Sid       string      `json:"Sid"` // optional
	Effect    string      `json:"Effect,omitempty"`
	Principal struct{}    `json:"Principal"` // optional
	Action    []string    `json:"Action,omitempty"`
	Resource  interface{} `json:"Resource,omitempty"` // I used interface{} because Resource may be either string and []string
	Condition struct{}    `json:"Condition"`
}

type PolicyDocument struct {
	Version   string      `json:"Version,omitempty"`
	Statement []Statement `json:"Statement,omitempty"`
}
