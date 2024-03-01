package appstacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/backend"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
	"gitlab.con/stream-ai/huddle/backend/cdk/vpc"
)

type BuildReturn struct {
	VpcStack     vpc.Stack
	BackendStack backend.Stack
}

func Build(scope constructs.Construct,
	appEnvName string,
	// vpc props
	vpcEnv shared.Environment,
	vpcMaxAzs float64,
	// backend props
	backendEnv shared.Environment,
	backendCpu float64,
	backendMemoryLimit float64,
) BuildReturn {
	// stackTags() returns a map of tags for a stack, including the "cdk-stack" tag
	stackTags := func(stackName string) map[string]*string {
		tags := make(map[string]*string)
		tags["cdk-stack"] = jsii.String(stackName)
		tags["app-environment"] = jsii.String(appEnvName)
		return tags
	}

	stackId := func(stackName string) shared.StackId {
		return shared.StackId(appEnvName + "-huddle-" + stackName)
	}

	vpcStack := vpc.NewStack(
		scope,
		stackId("vpc"),
		stackTags("vpc"),
		vpcEnv,
		vpcMaxAzs)

	vpc := vpcStack.Vpc()

	backendStack := backend.NewStack(
		scope,
		stackId("backend"),
		stackTags("backend"),
		backendEnv,
		backendCpu,
		backendMemoryLimit,
		vpc,
	)

	return BuildReturn{vpcStack, backendStack}
}
