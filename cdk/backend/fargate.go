package backend

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateProps struct {
	MemoryLimitMiB float64
	Cpu            float64
	Vpc            awsec2.IVpc
}

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

func NewFargateConstruct(scope constructs.Construct, id string, props *FargateProps) FargateConstruct {
	name := func(in string) *string {
		return jsii.String(fmt.Sprintf("%s/%s", in, id))
	}

	trafficPort := 80

	// Load Balanced Fargate Service
	assetImage := awsecs.ContainerImage_FromAsset(jsii.String("/workspaces/backend/"), &awsecs.AssetImageProps{
		File: jsii.String("./service/Dockerfile"),
	})

	loadBalancedFargateService := awsecspatterns.NewApplicationLoadBalancedFargateService(scope, name("Service"), &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Vpc:            props.Vpc,
		MemoryLimitMiB: jsii.Number(props.MemoryLimitMiB),
		Cpu:            jsii.Number(props.Cpu),
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image: assetImage,
		},
		PublicLoadBalancer: jsii.Bool(true),
		ListenerPort:       jsii.Number(trafficPort),
	})
	loadBalancedFargateService.TargetGroup().ConfigureHealthCheck(&awselasticloadbalancingv2.HealthCheck{
		Port: jsii.String(fmt.Sprintf("%d", trafficPort)),
		Path: jsii.String("/healthz"),
	})

	return &fargateConstruct{scope, loadBalancedFargateService}
}
