package backend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
)

type ZoneProvider interface {
	HostedZone(scope constructs.Construct, id shared.ResourceId) awsroute53.IHostedZone
}

type zoneProvider struct {
	domainName string
	hostedZone awsroute53.IHostedZone
}

func (z *zoneProvider) HostedZone(scope constructs.Construct, id shared.ResourceId) awsroute53.IHostedZone {
	zone := awsroute53.HostedZone_FromLookup(scope, id.String(), &awsroute53.HostedZoneProviderProps{
		DomainName: jsii.String(z.domainName),
	})
	return zone
}

// NewZoneLookupProvider returns a ZoneProvider that looks up the HostedZone by domain name

func NewZoneLookupProvider(domainName string) ZoneProvider {
	return &zoneProvider{domainName: domainName}
}

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
	envProvider shared.EnvProvider,
	scope constructs.Construct,
	id shared.StackId,
	tags map[string]*string,
	// Backend Stack Properties
	ecsCpu float64,
	ecsMemoryLimit float64,
	vpc awsec2.IVpc,
	zoneProvider ZoneProvider,
) Stack {
	cdkStack := awscdk.NewStack(scope, id.String(), &awscdk.StackProps{
		Tags: &tags,
		Env:  envProvider.Env(),
	})

	fargateConstruct := NewFargateConstruct(
		cdkStack,
		id.Construct("Fargate"),
		ecsMemoryLimit,
		ecsCpu,
		vpc,
		zoneProvider,
	)

	awscdk.NewCfnOutput(cdkStack, id.CfnOutput("LoadBalancerDNS"), &awscdk.CfnOutputProps{Value: fargateConstruct.FargateService().LoadBalancer().LoadBalancerDnsName()})

	return &stack{fargateConstruct}
}
