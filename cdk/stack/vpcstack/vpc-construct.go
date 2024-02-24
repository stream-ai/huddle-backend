package vpcstack

import "github.com/aws/constructs-go/constructs/v10"

type VpcProps struct {
	Tags *map[string]*string
}
type vpcConstruct struct {
	constructs.Construct
}

type VpcConstruct interface {
	constructs.Construct
}

func NewVpcConstruct(scope constructs.Construct, id string, props *VpcProps) VpcConstruct {
	this := constructs.NewConstruct(scope, &id)
	return vpcConstruct{this}
}
