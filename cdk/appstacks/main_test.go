package appstacks_test

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"
	"gitlab.con/stream-ai/huddle/backend/cdk/appstacks"
)

func Test_Integrationm(t *testing.T) {
	defer jsii.Close()

	type test struct {
		// vpc props
		vpcMaxAzs float64
		// backend props
		backendCpu         float64
		backendMemoryLimit float64
		backendDomainName  string
	}

	tests := []test{
		{
			// vpc props
			vpcMaxAzs: 2,
			// backend props
			backendCpu:         256,
			backendMemoryLimit: 512,
			backendDomainName:  "test.example.com",
		},
	}

	for _, tc := range tests {
		app := awscdk.NewApp(nil)

		// get an environment from the current AWS profile

		ret := appstacks.Build(app,
			"jdibling",
			tc.vpcMaxAzs,
			tc.backendCpu,
			tc.backendMemoryLimit,
			tc.backendDomainName,
		)

		// check the vpc stack
		azs := ret.VpcStack.Vpc().AvailabilityZones()
		assert.NotEmpty(t, azs)
		assert.Len(t, *azs, int(tc.vpcMaxAzs))

		// template := assertions.Template_FromStack(ret.VpcStack.Vpc().Stack(), nil)
		// b, e := json.MarshalIndent(template.ToJSON(), "", " ")
		// require.NoError(t, e)
		// log.Printf("%s\n", b)

		// check the backend stack
		assert.NotNil(t, ret.BackendStack.FargateConstruct())
		assert.NotNil(t, ret.BackendStack.LoadBalancerDNS())
	}
}
