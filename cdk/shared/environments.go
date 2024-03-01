package shared

type Environment interface {
	Account() string
	Region() string
	RoleArn() string
}

func NewEnvironment(account string, region string, roleArn string) Environment {
	return environment{account, region, roleArn}
}

type environment struct {
	account string
	region  string
	roleArn string
}

func (e environment) Account() string {
	return e.account
}

func (e environment) Region() string {
	return e.region
}

func (e environment) RoleArn() string {
	return e.roleArn
}

var EnvironmentMap map[string]Environment = map[string]Environment{
	"stream.core": NewEnvironment("674085691192", "us-east-1", "arn:aws:iam::674085691192:role/+huddle.cdk"),

	// "development": {
	// 	Account: "533267072195",
	// 	Region:  "us-east-1",
	// },
	// "sandbox.jdibling": {
	// 	Account: "590184032693",
	// 	Region:  "us-east-1",
	// },
	// "networking": {
	// 	Account: "674085691192",
	// 	Region:  "us-east-1",
	// },
	// "security": {
	// 	Account: "058264533892",
	// 	Region:  "us-east-1",
	// },
	// "shared-services": {
	// 	Account: "339713158493",
	// 	Region:  "us-east-1",
	// },
	// "production.1": {
	// 	Account: "905418044184",
	// 	Region:  "us-east-1",
	// },
	// "staging": {
	// 	Account: "471112991101",
	// 	Region:  "us-east-1",
	// },
}
