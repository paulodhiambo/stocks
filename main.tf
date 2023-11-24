locals {
  vpc_id           = "vpc-07a589482a6a7e802"
  subnet_id        = "subnet-0fc4d8f4c5c174038"
  ssh_user         = "ubuntu"
  key_name         = "devops"
  private_key_path = "~/Downloads/devops.pem"


}
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }
  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-west-2"
}
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}
resource "aws_security_group" "nginx" {
  name   = "nginx_access"
  vpc_id = local.vpc_id

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
resource "aws_instance" "app_server" {
  ami                         = data.aws_ami.ubuntu.id
  # subnet_id                   = local.subnet_id
  associate_public_ip_address = true
  instance_type               = "t2.micro"
  vpc_security_group_ids      = [aws_security_group.nginx.id]
  key_name                    = local.key_name

  provisioner "remote-exec" {
    inline = ["echo 'Wait until ssh is ready'"]

    connection {
      type        = "ssh"
      user        = local.ssh_user
      private_key = file(local.private_key_path)
      host        = self.public_ip
    }
  }
  provisioner "local-exec" {
    command = "ansible-playbook -i ${aws_instance.app_server.public_ip}, --private-key ${local.private_key_path} nginx.yaml"
  }

  tags = {
    Name = "Stock Instance"
  }

}
output "nginx_ip" {
  value = aws_instance.app_server.public_ip
}
