terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.27.0"
    }
  }
}

provider "aws" {
  region = var.region
}


resource "aws_security_group" "platform_sg" {
  name        = "platform_sg"
  description = "Security group for platform"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "platform_ec2" {
  ami                    = "ami-00ca570c1b6d79f36"
  instance_type          = var.instance_type
  key_name               = var.key_name
  vpc_security_group_ids = [aws_security_group.platform_sg.id]

  user_data = <<-EOF
              #!/bin/bash
              sudo dnf update -y
              sudo dnf install docker -y
              sudo systemctl start docker
              sudo systemctl enable docker
              sudo usermod -a -G docker ec2-user
              sudo dnf install git -y
              EOF

  tags = {
    Name = "mini-paas-ec2"
  }
}
