package ec2stack

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/instance"
	"cdk.tf/go/stack/generated/hashicorp/aws/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func Ec2Stack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	provider.NewAwsProvider(stack, jsii.String("aws"), &provider.AwsProviderConfig{
		Region: jsii.String("us-east-1"),
	})

	instance.NewInstance(stack, jsii.String("cdktfgo"), &instance.InstanceConfig{
		Ami:          jsii.String("ami-0fc5d935ebf8bc3bc"),
		InstanceType: jsii.String("t2.micro"),
	})

	return stack
}
