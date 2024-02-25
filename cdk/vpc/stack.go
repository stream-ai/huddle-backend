package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

type VpcStackProps struct {
	Env    awscdk.Environment
	Tags   map[string]*string
	MaxAzs float64
}

type vpcStack struct {
	// awscdk.Stack
	vpc awsec2.IVpc
}

func (v *vpcStack) Vpc() awsec2.IVpc {
	return v.vpc
}

type VpcStack interface {
	// awscdk.Stack
	Vpc() awsec2.IVpc
}

func NewStack(scope constructs.Construct, id string, props *VpcStackProps) VpcStack {
	stack := awscdk.NewStack(scope, &id, &awscdk.StackProps{
		Tags: &props.Tags,
	})

	vpcConstruct := NewVpcConstruct(stack, "HuddleBackendVpc", &VpcProps{
		MaxAzs: props.MaxAzs,
	})

	return &vpcStack{
		// awscdk.Stack(stack),
		vpc: vpcConstruct.Vpc(),
	}
}
