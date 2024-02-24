package vpcstack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Props struct {
	awscdk.StackProps
	Tags *map[string]*string
}

func New(scope constructs.Construct, id string, props *Props) awscdk.Stack {
	stackTags := map[string]*string{
		"cdk-stack": jsii.String("vpcstack"),
	}
	if props.Tags != nil {
		for k, v := range *props.Tags {
			stackTags[k] = v
		}
	}

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
		sprops.Tags = &stackTags
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awsec2.NewVpc(stack, jsii.String("HuddleVpc"), &awsec2.VpcProps{
		MaxAzs: jsii.Number(2),
	})

	return stack
}
