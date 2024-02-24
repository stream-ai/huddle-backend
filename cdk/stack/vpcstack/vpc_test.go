package vpcstack_test

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"
	"gitlab.con/stream-ai/huddle/backend/cdk/stack/vpcstack"
)

func TestVpcStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := vpcstack.New(app, "TestStack", &vpcstack.Props{
		Tags: &map[string]*string{
			"Environment": jsii.String("Dev"),
		},
	})

	// THEN
	template := assertions.Template_FromStack(stack, nil)
	template.HasResourceProperties(jsii.String("AWS::EC2::VPC"), map[string]any{
		"Tags": assertions.Match_AnyValue(),
	})
	tags := *stack.Tags().TagValues()
	assert.Contains(t, tags, "cdk-stack")
	assert.Equal(t, *(tags)["cdk-stack"], "vpcstack")
}
