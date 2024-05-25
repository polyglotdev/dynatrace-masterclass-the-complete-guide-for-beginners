output "vpc_id" {
  description = "The ID of the VPC"
  value       = aws_vpc.main.id
}

output "subnet1_id" {
  description = "The ID of subnet 1"
  value       = aws_subnet.subnet1.id
}

output "subnet2_id" {
  description = "The ID of subnet 2"
  value       = aws_subnet.subnet2.id
}

output "iam_policy_arn" {
  description = "The ARN of the IAM policy"
  value       = aws_iam_policy.god_mode.arn
}

output "instance_id" {
  description = "The ID of the EC2 instance"
  value       = aws_instance.example.id
}