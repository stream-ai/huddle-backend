package vpc

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type VpcProps struct {
	MaxAzs float64
}

type VpcConstruct interface {
	constructs.Construct
	Vpc() awsec2.IVpc
}

type vpcConstruct struct {
	constructs.Construct
	vpc *awsec2.Vpc
}

func (v *vpcConstruct) Vpc() awsec2.IVpc {
	return *v.vpc
}

func NewVpcConstruct(scope constructs.Construct, id string, props *VpcProps) VpcConstruct {
	this := constructs.NewConstruct(scope, &id)

	// Create VPC and Cluster
	vpc := awsec2.NewVpc(this, jsii.String("HuddleBackendVpc"), &awsec2.VpcProps{
		MaxAzs: jsii.Number(props.MaxAzs),
	})

	return &vpcConstruct{this, &vpc}
}
