terraform {
  backend "remote" {
    organization = "polyglotdev"

    workspaces {
      name = "dynatrace-masterclass-the-complete-guide-for-beginners"
    }
  }
}

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

# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
  tags = {
    Name = "main_igw"
  }
}

# Route Table
resource "aws_route_table" "main" {
  vpc_id = aws_vpc.main.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }
  tags = {
    Name = "main_route_table"
  }
}

# Subnets
resource "aws_subnet" "subnet1" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = var.subnet1_cidr
  availability_zone       = var.availability_zone1
  map_public_ip_on_launch = true
  tags = {
    Name = "subnet1"
  }
}

resource "aws_route_table_association" "subnet1" {
  subnet_id      = aws_subnet.subnet1.id
  route_table_id = aws_route_table.main.id
}

resource "aws_subnet" "subnet2" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = var.subnet2_cidr
  availability_zone       = var.availability_zone2
  map_public_ip_on_launch = true
  tags = {
    Name = "subnet2"
  }
}

resource "aws_route_table_association" "subnet2" {
  subnet_id      = aws_subnet.subnet2.id
  route_table_id = aws_route_table.main.id
}

# Security Group
resource "aws_security_group" "ssh_access" {
  vpc_id = aws_vpc.main.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Restrict this in a production environment
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "ssh_access"
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
  ami                    = data.aws_ami.amazon_linux.id
  instance_type          = var.instance_type
  subnet_id              = aws_subnet.subnet1.id
  vpc_security_group_ids = [aws_security_group.ssh_access.id]

  tags = {
    Name = "example_instance"
  }
}