package backend_test

import (
	"log"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"
	"gitlab.con/stream-ai/huddle/backend/cdk/backend"
)

type mockVpcConstruct struct {
	awsec2.IVpc
}

func TestNewFargateConstruct(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := awscdk.NewStack(app, jsii.String("test"), &awscdk.StackProps{})
	vpc := awsec2.NewVpc(stack, jsii.String("test"), &awsec2.VpcProps{
		MaxAzs: jsii.Number(2),
	})
	log.Printf("stack: %+v", stack)

	// Create the FargateConstruct
	fargate := backend.NewFargateConstruct(stack, "test",
		1024,
		256,
		vpc,
	)

	// Assert that the FargateConstruct is created correctly
	if fargate == nil {
		t.Error("FargateConstruct is nil")
	}

	// Assert that the health check is on port 80 at /healthz
	healthCheck := fargate.HealthCheck()
	assert.NotNil(t, healthCheck)
	assert.NotNil(t, healthCheck.Port)
	assert.NotNil(t, healthCheck.Path)
	if *healthCheck.Port != "80" || *healthCheck.Path != "/healthz" {
		t.Errorf("Expected health check on port 80 at /healthz, got port %s and path %s", *healthCheck.Port, *healthCheck.Path)
	}
}
