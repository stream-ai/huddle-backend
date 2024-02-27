package shared

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

var EnvironmentMap map[string]awscdk.Environment = map[string]awscdk.Environment{
	"development": {
		Account: jsii.String("533267072195"),
		Region:  jsii.String("us-east-1"),
	},
	"sandbox.jdibling": {
		Account: jsii.String("590184032693"),
		Region:  jsii.String("us-east-1"),
	},
	"networking": {
		Account: jsii.String("674085691192"),
		Region:  jsii.String("us-east-1"),
	},
	"security": {
		Account: jsii.String("058264533892"),
		Region:  jsii.String("us-east-1"),
	},
	"shared-services": {
		Account: jsii.String("339713158493"),
		Region:  jsii.String("us-east-1"),
	},
	"production.1": {
		Account: jsii.String("905418044184"),
		Region:  jsii.String("us-east-1"),
	},
	"staging": {
		Account: jsii.String("471112991101"),
		Region:  jsii.String("us-east-1"),
	},
	"stream.core": {
		Account: jsii.String("674085691192"),
		Region:  jsii.String("us-east-1"),
	},
}
