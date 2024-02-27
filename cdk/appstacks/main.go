package appstacks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/backend"
	"gitlab.con/stream-ai/huddle/backend/cdk/vpc"
)

type VpcProps struct {
	Env   awscdk.Environment
	MaxAz float64
}

type BackendProps struct {
	Env         awscdk.Environment
	Cpu         float64
	MemoryLimit float64
}

type BuildReturn struct {
	VpcStack     vpc.VpcStack
	BackendStack backend.BackendStack
}

func Build(scope constructs.Construct,
	appEnvName string,
	vpcProps VpcProps,
	backendProps BackendProps,
) BuildReturn {
	// stackTags() returns a map of tags for a stack, including the "cdk-stack" tag
	stackTags := func(stackName string) map[string]*string {
		tags := make(map[string]*string)
		tags["cdk-stack"] = jsii.String(stackName)
		tags["app-environment"] = jsii.String(appEnvName)
		return tags
	}

	stackId := func(stackName string) string {
		return appEnvName + "-huddle-" + stackName
	}

	vpcStack := vpc.NewStack(scope, stackId("vpc"), &vpc.VpcStackProps{
		Env:    vpcProps.Env,
		MaxAzs: vpcProps.MaxAz,
		Tags:   stackTags("vpc"),
	})

	vpc := vpcStack.Vpc()

	backendStack := backend.NewStack(scope, stackId("backend"), &backend.BackendStackProps{
		Env:         backendProps.Env,
		Tags:        stackTags("backend"),
		Cpu:         backendProps.Cpu,
		MemoryLimit: backendProps.MemoryLimit,
		Vpc:         vpc,
	})

	return BuildReturn{vpcStack, backendStack}
}
