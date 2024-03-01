package shared

import "github.com/aws/jsii-runtime-go"

type (
	StackId     string
	ConstructId string
)

func (s StackId) String() string {
	return string(s)
}

func (s StackId) Construct(id string) ConstructId {
	return ConstructId(s.String() + id)
}

func (s StackId) CfnOutput(id string) *string {
	return jsii.String(s.String() + id)
}

func (c ConstructId) String() string {
	return string(c)
}

func (c ConstructId) Resource(id string) *string {
	return jsii.String(c.String() + id)
}
