output "ec2_public_ip" {
  value = aws_instance.platform_ec2.public_ip
}
