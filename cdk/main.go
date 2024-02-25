package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/backend"
	"gitlab.con/stream-ai/huddle/backend/cdk/vpc"
)

var EnvironmentMap map[string]awscdk.Environment = map[string]awscdk.Environment{
	"development": {
		Account: jsii.String("533267072195"),
		Region:  jsii.String("us-east-1"),
	},
	"sandbox.jdibling": {
		Account: jsii.String("590184032693"),
		Region:  jsii.String("us-east-1"),
	},
	"networking": {
		Account: jsii.String("674085691192"),
		Region:  jsii.String("us-east-1"),
	},
	"security": {
		Account: jsii.String("058264533892"),
		Region:  jsii.String("us-east-1"),
	},
	"shared-services": {
		Account: jsii.String("339713158493"),
		Region:  jsii.String("us-east-1"),
	},
	"production.1": {
		Account: jsii.String("905418044184"),
		Region:  jsii.String("us-east-1"),
	},
	"staging": {
		Account: jsii.String("471112991101"),
		Region:  jsii.String("us-east-1"),
	},
}

type vpcAppStackProps struct {
	Env   awscdk.Environment
	MaxAz float64
}

type backendAppStackProps struct {
	Env         awscdk.Environment
	Cpu         float64
	MemoryLimit float64
}

func buildAppStacks(scope constructs.Construct,
	appEnvName string,
	vpcProps vpcAppStackProps,
	backendProps backendAppStackProps,
) {
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

	backend.NewStack(scope, stackId("backend"), &backend.BackendStackProps{
		Env:         backendProps.Env,
		Tags:        stackTags("backend"),
		Cpu:         backendProps.Cpu,
		MemoryLimit: backendProps.MemoryLimit,
		Vpc:         vpc,
	})
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// Build app stacks
	buildAppStacks(app, "jdibling",
		vpcAppStackProps{
			Env:   EnvironmentMap["sandbox.jdibling"],
			MaxAz: 2,
		},
		backendAppStackProps{
			Env:         EnvironmentMap["sandbox.jdibling"],
			Cpu:         256,
			MemoryLimit: 512,
		},
	)

	app.Synth(nil)
}
