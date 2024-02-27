package appstacks_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.con/stream-ai/huddle/backend/cdk/appstacks"
)

var mockEnv awscdk.Environment = awscdk.Environment{
	Account: jsii.String("123456789012"),
	Region:  jsii.String("us-west-2"),
}

func Test_Build(t *testing.T) {
	type test struct {
		vpc     appstacks.VpcProps
		backend appstacks.BackendProps
	}

	tests := []test{
		{
			vpc: appstacks.VpcProps{
				Env:   mockEnv,
				MaxAz: 2,
			},
			backend: appstacks.BackendProps{
				Env:         mockEnv,
				Cpu:         256,
				MemoryLimit: 512,
			},
		},
	}

	for _, tc := range tests {
		app := awscdk.NewApp(nil)

		ret := appstacks.Build(app, "jdibling", tc.vpc, tc.backend)

		// check the vpc stack
		azs := ret.VpcStack.Vpc().AvailabilityZones()
		assert.NotEmpty(t, azs)
		assert.Len(t, *azs, int(tc.vpc.MaxAz))

		template := assertions.Template_FromStack(ret.VpcStack.Vpc().Stack(), nil)
		b, e := json.MarshalIndent(template.ToJSON(), "", " ")
		require.NoError(t, e)
		log.Printf("%s\n", b)

		// check the backend stack
		assert.NotNil(t, ret.BackendStack.FargateConstruct())
		assert.NotNil(t, ret.BackendStack.LoadBalancerDNS())

	}
}
