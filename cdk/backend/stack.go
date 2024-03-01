package backend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
)

type stack struct {
	fargateConstruct FargateConstruct
}

func (b *stack) LoadBalancerDNS() *string {
	return b.fargateConstruct.FargateService().LoadBalancer().LoadBalancerDnsName()
}

func (b *stack) FargateConstruct() constructs.Construct {
	return b.fargateConstruct
}

type Stack interface {
	FargateConstruct() constructs.Construct
	LoadBalancerDNS() *string
}

func NewStack(
	// Common Stack Properties
	scope constructs.Construct,
	id shared.StackId,
	tags map[string]*string,
	// Backend Stack Properties
	stackEnv shared.Environment,
	ecsCpu float64,
	ecsMemoryLimit float64,
	vpc awsec2.IVpc,
	domainName string,
) Stack {
	// TODO: assume a cross-account role if `env` is not nil
	var env *awscdk.Environment
	if stackEnv != nil {
		env = stackEnv.CDK()
	}
	cdkStack := awscdk.NewStack(scope, jsii.String(string(id)), &awscdk.StackProps{
		Tags: &tags,
		Env:  env,
	})

	domainZone := awsroute53.HostedZone_FromLookup(cdkStack, id.Construct("domainZone").Resource("zone"), &awsroute53.HostedZoneProviderProps{
		DomainName: jsii.String(domainName),
	})

	fargateConstruct := NewFargateConstruct(
		cdkStack,
		id.Construct("Fargate"),
		ecsMemoryLimit,
		ecsCpu,
		vpc,
		domainZone,
	)

	awscdk.NewCfnOutput(cdkStack, id.CfnOutput("LoadBalancerDNS"), &awscdk.CfnOutputProps{Value: fargateConstruct.FargateService().LoadBalancer().LoadBalancerDnsName()})

	return &stack{fargateConstruct}
}
