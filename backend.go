package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateWithALBStackProps struct {
	awscdk.StackProps
}

func NewFargateWithALBStack(scope constructs.Construct, id string, props *FargateWithALBStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
		sprops.Tags = &map[string]*string{
			"application": jsii.String("huddle-backend"),
		}
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create VPC and Cluster
	vpc := awsec2.NewVpc(stack, jsii.String("HuddleBackendVpc"), &awsec2.VpcProps{
		MaxAzs: jsii.Number(2),
	})

	// Load Balanced Fargate Service
	loadBalancedFargateService := awsecspatterns.NewApplicationLoadBalancedFargateService(stack, jsii.String("HuddleBackendService"), &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		Vpc:            vpc,
		MemoryLimitMiB: jsii.Number(512),
		Cpu:            jsii.Number(256),
		TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
			Image: awsecs.ContainerImage_FromAsset(jsii.String("."), &awsecs.AssetImageProps{
				File: jsii.String("app/Dockerfile"),
			}),
		},
		PublicLoadBalancer: jsii.Bool(true),
	})
	loadBalancedFargateService.TargetGroup().ConfigureHealthCheck(&awselasticloadbalancingv2.HealthCheck{
		Path: jsii.String("/healthz"),
	})

	// Output the Load Balancer DNS
	awscdk.NewCfnOutput(stack, jsii.String("LoadBalancerDNS"), &awscdk.CfnOutputProps{Value: loadBalancedFargateService.LoadBalancer().LoadBalancerDnsName()})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewFargateWithALBStack(app, "FargateWithALBStack", &FargateWithALBStackProps{awscdk.StackProps{
		Env: env(),
	}})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
