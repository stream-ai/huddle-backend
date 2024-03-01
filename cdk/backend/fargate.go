package backend

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"gitlab.con/stream-ai/huddle/backend/cdk/shared"
)

type FargateConstruct interface {
	constructs.Construct
	FargateService() awsecspatterns.ApplicationLoadBalancedFargateService
	HealthCheck() *awselasticloadbalancingv2.HealthCheck
}

type fargateConstruct struct {
	constructs.Construct
	fargateService awsecspatterns.ApplicationLoadBalancedFargateService
}

func (f *fargateConstruct) FargateService() awsecspatterns.ApplicationLoadBalancedFargateService {
	return f.fargateService
}

func (f *fargateConstruct) HealthCheck() *awselasticloadbalancingv2.HealthCheck {
	return f.fargateService.TargetGroup().HealthCheck()
}

func NewFargateConstruct(
	// Common construct props
	scope constructs.Construct,
	id shared.ConstructId,
	// Fargate construct props
	memoryLimitMiB float64,
	cpu float64,
	vpc awsec2.IVpc,
	zoneProvider ZoneProvider,
) FargateConstruct {
	trafficPort := 80

	// Load Balanced Fargate Service
	assetImage := awsecs.ContainerImage_FromAsset(jsii.String("/workspaces/backend/"), &awsecs.AssetImageProps{
		File: jsii.String("./service/Dockerfile"),
	})

	domainZone := zoneProvider.HostedZone(scope, id.Resource("zone"))

	loadBalancedFargateService := awsecspatterns.NewApplicationLoadBalancedFargateService(scope, id.Resource("service").String(), &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Vpc:            vpc,
		MemoryLimitMiB: jsii.Number(memoryLimitMiB),
		Cpu:            jsii.Number(cpu),
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image: assetImage,
		},
		PublicLoadBalancer: jsii.Bool(true),
		ListenerPort:       jsii.Number(trafficPort),
		DomainName:         domainZone.ZoneName(),
		DomainZone:         domainZone,
	})
	loadBalancedFargateService.TargetGroup().ConfigureHealthCheck(&awselasticloadbalancingv2.HealthCheck{
		Port: jsii.String(fmt.Sprintf("%d", trafficPort)),
		Path: jsii.String("/healthz"),
	})

	return &fargateConstruct{scope, loadBalancedFargateService}
}
