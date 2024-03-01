package appstacks_test

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/stretchr/testify/assert"
	"gitlab.con/stream-ai/huddle/backend/cdk/appstacks"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
)

type mockEnvironment struct {
	account string
	region  string
	roleArn string
}

func (e *mockEnvironment) Account() string {
	return e.account
}

func (e *mockEnvironment) Region() string {
	return e.region
}

func (e *mockEnvironment) RoleArn() string {
	return e.roleArn
}

var mockEnv = mockEnvironment{
	account: "123456789012",
	region:  "us-west-2",
	roleArn: "arn:aws:iam::123456789012:role/+huddle.cdk",
}

func Test_Build(t *testing.T) {
	type test struct {
		// vpc props
		vpcEnv    shared.Environment
		vpcMaxAzs float64
		// backend props
		backendEnv         shared.Environment
		backendCpu         float64
		backendMemoryLimit float64
	}

	tests := []test{
		{
			// vpc props
			vpcEnv:    nil,
			vpcMaxAzs: 2,
			// backend props
			backendEnv:         nil,
			backendCpu:         256,
			backendMemoryLimit: 512,
		},
	}

	for _, tc := range tests {
		app := awscdk.NewApp(nil)

		ret := appstacks.Build(app,
			"jdibling",
			tc.vpcEnv,
			tc.vpcMaxAzs,
			tc.backendEnv,
			tc.backendCpu,
			tc.backendMemoryLimit,
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
