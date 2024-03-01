package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/appstacks"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// Build app stacks
	appstacks.Build(app,
		"jdibling",
		// vpc props
		nil, // vpcEnv
		2,   // maxAzs
		// backend props
		nil,
		256, // cpu
		512, // memoryLimit
	)

	app.Synth(nil)
}
