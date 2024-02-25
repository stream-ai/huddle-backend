package backend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type BackendStackProps struct {
	Env         awscdk.Environment
	Tags        map[string]*string
	Cpu         float64
	MemoryLimit float64
	Vpc         awsec2.IVpc
}

type backendStack struct {
	fargateConstruct FargateConstruct
	loadBalancerDNS  *string
}

func (b *backendStack) LoadBalancerDNS() *string {
	return b.loadBalancerDNS
}

func (b *backendStack) FargateConstruct() constructs.Construct {
	return b.fargateConstruct
}

type BackendStack interface {
	FargateConstruct() constructs.Construct
	LoadBalancerDNS() *string
}

func NewStack(scope constructs.Construct, id string, props *BackendStackProps) BackendStack {
	stack := awscdk.NewStack(scope, &id, &awscdk.StackProps{
		Tags: &props.Tags,
	})

	fargateConstruct := NewFargateConstruct(stack, "HuddleBackendService", &FargateProps{
		MemoryLimitMiB: props.MemoryLimit,
		Cpu:            props.Cpu,
		Vpc:            props.Vpc,
	})

	awscdk.NewCfnOutput(stack, jsii.String("LoadBalancerDNS"), &awscdk.CfnOutputProps{Value: fargateConstruct.FargateService().LoadBalancer().LoadBalancerDnsName()})

	return &backendStack{fargateConstruct, fargateConstruct.FargateService().LoadBalancer().LoadBalancerDnsName()}
}
