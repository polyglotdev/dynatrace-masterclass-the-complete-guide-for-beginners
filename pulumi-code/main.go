package main

// import (
//
//

// 	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )

// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		// Create a new VPC
// 		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Create a public subnet
// 		publicSubnet, err := ec2.NewSubnet(ctx, "publicSubnet", &ec2.SubnetArgs{
// 			VpcId:               vpc.ID(),
// 			CidrBlock:           pulumi.String("10.0.1.0/24"),
// 			MapPublicIpOnLaunch: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Create an Internet Gateway
// 		igw, err := ec2.NewInternetGateway(ctx, "igw", &ec2.InternetGatewayArgs{
// 			VpcId: vpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Create a route table
// 		routeTable, err := ec2.NewRouteTable(ctx, "routeTable", &ec2.RouteTableArgs{
// 			VpcId: vpc.ID(),
// 			Routes: ec2.RouteTableRouteArray{
// 				&ec2.RouteTableRouteArgs{
// 					CidrBlock: pulumi.String("0.0.0.0/0"),
// 					GatewayId: igw.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Associate the route table with the public subnet
// 		_, err = ec2.NewRouteTableAssociation(ctx, "rta", &ec2.RouteTableAssociationArgs{
// 			SubnetId:     publicSubnet.ID(),
// 			RouteTableId: routeTable.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Create a Security Group to allow SSH traffic
// 		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
// 			VpcId: vpc.ID(),
// 			Ingress: ec2.SecurityGroupIngressArray{
// 				&ec2.SecurityGroupIngressArgs{
// 					Protocol:   pulumi.String("tcp"),
// 					FromPort:   pulumi.Int(22),
// 					ToPort:     pulumi.Int(22),
// 					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Create an IAM policy with access to all services
// 		_, err = iam.NewPolicy(ctx, "allAccessPolicy", &iam.PolicyArgs{
// 			Policy: pulumi.String(`{
// 				"Version": "2012-10-17",
// 				"Statement": [
// 					{
// 						"Action": "*",
// 						"Effect": "Allow",
// 						"Resource": "*"
// 					}
// 				]
// 			}`),
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Use the provided Amazon Linux 2 AMI ID
// 		amiId := "ami-0bb84b8ffd87024d8"

// 		// Create an EC2 instance in the public subnet
// 		_, err = ec2.NewInstance(ctx, "instance", &ec2.InstanceArgs{
// 			InstanceType: pulumi.String("t2.micro"),
// 			Ami:          pulumi.String(amiId),
// 			SubnetId:     publicSubnet.ID(),
// 			VpcSecurityGroupIds: pulumi.StringArray{
// 				securityGroup.ID(),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }
