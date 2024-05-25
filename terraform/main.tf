provider "aws" {
  region = var.aws_region
}

# Retrieve the latest Amazon Linux 2 AMI
data "aws_ami" "amazon_linux" {
  most_recent = true
  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
  filter {
    name   = "owner-id"
    values = ["137112412989"]
  }
}

# VPC
resource "aws_vpc" "main" {
  cidr_block = var.vpc_cidr
  tags = {
    Name = "main_vpc"
  }
}

# Subnets
resource "aws_subnet" "subnet1" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.subnet1_cidr
  availability_zone = var.availability_zone1
  tags = {
    Name = "subnet1"
  }
}

resource "aws_subnet" "subnet2" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.subnet2_cidr
  availability_zone = var.availability_zone2
  tags = {
    Name = "subnet2"
  }
}

# IAM Policy
resource "aws_iam_policy" "god_mode" {
  name        = "god_mode_policy"
  description = "A policy with full permissions"
  policy      = file("policy.json")
}

# EC2 Instance
resource "aws_instance" "example" {
  ami           = data.aws_ami.amazon_linux.id
  instance_type = var.instance_type
  subnet_id     = aws_subnet.subnet1.id
  tags = {
    Name = "example_instance"
  }
}