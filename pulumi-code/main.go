package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Configuration variables
		instanceType := "t2.micro"

		// Create a VPC
		vpc, err := ec2.NewVpc(ctx, "mainVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main_vpc"),
			},
		})
		if err != nil {
			return err
		}

		// Create an Internet Gateway
		igw, err := ec2.NewInternetGateway(ctx, "mainIgw", &ec2.InternetGatewayArgs{
			VpcId: vpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main_igw"),
			},
		})
		if err != nil {
			return err
		}

		// Create a Route Table
		routeTable, err := ec2.NewRouteTable(ctx, "mainRouteTable", &ec2.RouteTableArgs{
			VpcId: vpc.ID(),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("0.0.0.0/0"),
					GatewayId: igw.ID(),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main_route_table"),
			},
		})
		if err != nil {
			return err
		}

		// Create a Subnet
		subnet, err := ec2.NewSubnet(ctx, "subnet1", &ec2.SubnetArgs{
			VpcId:               vpc.ID(),
			CidrBlock:           pulumi.String("10.0.1.0/24"),
			AvailabilityZone:    pulumi.String("us-east-1a"),
			MapPublicIpOnLaunch: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("subnet1"),
			},
		})
		if err != nil {
			return err
		}

		_, err = ec2.NewRouteTableAssociation(ctx, "subnet1RouteTableAssociation", &ec2.RouteTableAssociationArgs{
			SubnetId:     subnet.ID(),
			RouteTableId: routeTable.ID(),
		})
		if err != nil {
			return err
		}

		// Retrieve the specific Amazon Linux 2023 AMI
		ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			Owners: []string{"137112412989"}, // The owner ID for Amazon AMIs
			Filters: []ec2.GetAmiFilter{
				{
					Name:   "image-id",
					Values: []string{"ami-0bb84b8ffd87024d8"}, // Specific AMI ID for Amazon Linux 2023
				},
			},
		})
		if err != nil {
			return err
		}

		// Create an IAM Role
		role, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "ec2.amazonaws.com"
                        },
                        "Action": "sts:AssumeRole"
                    }
                ]
            }`),
		})
		if err != nil {
			return err
		}

		// Create an IAM Policy
		policyJSON, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Effect":   "Allow",
					"Action":   "*",
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		policy, err := iam.NewPolicy(ctx, "godMode", &iam.PolicyArgs{
			Name:        pulumi.String("god_mode_policy"),
			Description: pulumi.String("A policy with full permissions"),
			Policy:      pulumi.String(policyJSON),
		})
		if err != nil {
			return err
		}

		// Attach the policy to the role
		_, err = iam.NewRolePolicyAttachment(ctx, "rolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      role.Name,
			PolicyArn: policy.Arn,
		})
		if err != nil {
			return err
		}

		// Create an IAM Instance Profile
		instanceProfile, err := iam.NewInstanceProfile(ctx, "instanceProfile", &iam.InstanceProfileArgs{
			Role: role.Name,
		})
		if err != nil {
			return err
		}

		// Create a Security Group
		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
			VpcId: vpc.ID(),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol:   pulumi.String("tcp"),
					FromPort:   pulumi.Int(22),
					ToPort:     pulumi.Int(22),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort:   pulumi.Int(0),
					ToPort:     pulumi.Int(0),
					Protocol:   pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("ssh_access"),
			},
		})
		if err != nil {
			return err
		}

		// Create an EC2 Instance
		_, err = ec2.NewInstance(ctx, "exampleInstance", &ec2.InstanceArgs{
			Ami:                 pulumi.String(ami.Id),
			InstanceType:        pulumi.String(instanceType),
			SubnetId:            subnet.ID(),
			VpcSecurityGroupIds: pulumi.StringArray{securityGroup.ID()},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example_instance"),
			},
			IamInstanceProfile: instanceProfile.Name, // Attach IAM instance profile to EC2 instance
		})
		if err != nil {
			return err
		}

		return nil
	})
}
