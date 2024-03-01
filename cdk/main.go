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
		2,                  // vpc.maxAzs
		256,                // backend.cpu
		512,                // backend.memoryLimit
		"jdibling.hudl.ai", // backend.domainName
	)

	app.Synth(nil)
}
