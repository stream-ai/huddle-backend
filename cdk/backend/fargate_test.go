package backend_test

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"
	"gitlab.con/stream-ai/huddle/backend/cdk/backend"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
)

type mockVpcConstruct struct {
	awsec2.IVpc
}

func TestNewFargateConstruct(t *testing.T) {
	defer jsii.Close()

	type test struct {
		env    shared.Environment
		memory float64
		cpu    float64
		domain string
	}

	tests := []test{
		{
			env:    shared.NewEnvironment("123456789012", "us-west-2", "arn:aws:iam::123456789012:role/+huddle.cdk"),
			memory: 1024,
			cpu:    256,
			domain: "test",
		},
	}

	for _, tc := range tests {
		app := awscdk.NewApp(nil)
		stack := awscdk.NewStack(app, jsii.String("stack"), &awscdk.StackProps{})
		vpc := awsec2.NewVpc(stack, jsii.String("vpc"), &awsec2.VpcProps{
			MaxAzs: jsii.Number(2),
		})
		domainZone := awsroute53.NewHostedZone(stack, jsii.String("zone"), &awsroute53.HostedZoneProps{
			ZoneName: jsii.String(tc.domain),
		})

		// Create the FargateConstruct
		fargate := backend.NewFargateConstruct(stack, "fargate",
			tc.memory,
			tc.cpu,
			vpc,
			domainZone,
		)

		// Assert that the FargateConstruct is created correctly
		assert.NotNil(t, fargate)
		// Assert that the health check is on port 80 at /healthz
		healthCheck := fargate.HealthCheck()
		assert.NotNil(t, healthCheck)
		assert.NotNil(t, healthCheck.Port)
		assert.NotNil(t, healthCheck.Path)
		if *healthCheck.Port != "80" || *healthCheck.Path != "/healthz" {
			t.Errorf("Expected health check on port 80 at /healthz, got port %s and path %s", *healthCheck.Port, *healthCheck.Path)
		}
	}
}
