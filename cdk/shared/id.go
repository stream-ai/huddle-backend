package shared

import "github.com/aws/jsii-runtime-go"

type (
	StackId     string
	ConstructId string
	ResourceId  string
)

const Sep = "-"

func (s StackId) String() *string {
	return jsii.String(string(s))
}

func (s StackId) Construct(id string) ConstructId {
	return ConstructId(*s.String() + Sep + id)
}

func (s StackId) CfnOutput(id string) *string {
	return jsii.String(*s.String() + id)
}

func (c ConstructId) String() *string {
	return jsii.String(string(c))
}

func (c ConstructId) Resource(id string) ResourceId {
	return ResourceId(*c.String() + Sep + id)
}

func (r ResourceId) String() *string {
	return jsii.String(string(r))
}
