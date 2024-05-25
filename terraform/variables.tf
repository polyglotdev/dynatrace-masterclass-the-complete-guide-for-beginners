variable "aws_region" {
  description = "The AWS region to deploy to"
  default     = "us-east-1"
}

variable "vpc_cidr" {
  description = "The CIDR block for the VPC"
  default     = "10.0.0.0/16"
}

variable "subnet1_cidr" {
  description = "The CIDR block for subnet 1"
  default     = "10.0.1.0/24"
}

variable "subnet2_cidr" {
  description = "The CIDR block for subnet 2"
  default     = "10.0.2.0/24"
}

variable "availability_zone1" {
  description = "The availability zone for subnet 1"
  default     = "us-east-1a"
}

variable "availability_zone2" {
  description = "The availability zone for subnet 2"
  default     = "us-east-1b"
}

variable "instance_type" {
  description = "The instance type for the EC2 instance"
  default     = "t2.micro"
}