package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/appstacks"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// Build app stacks
	appstacks.Build(app, "jdibling",
		appstacks.VpcProps{
			Env:   shared.EnvironmentMap["sandbox.jdibling"],
			MaxAz: 2,
		},
		appstacks.BackendProps{
			Env:         shared.EnvironmentMap["sandbox.jdibling"],
			Cpu:         256,
			MemoryLimit: 512,
		},
	)

	app.Synth(nil)
}
