package main

import (
	"cdk.tf/go/stack/ec2stack"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	stack := ec2stack.Ec2Stack(app, "CDKTF-Golang")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("MyOrg-23"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("tfc-demo-cdktf-go")),
	})

	app.Synth()
}
