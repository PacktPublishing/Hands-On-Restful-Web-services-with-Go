provider "aws" {
  profile = "default"
  region  = "eu-central-1"
}

resource "aws_key_pair" "api_server_key" {
  key_name   = "api-server-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCd2nPNeuU28LNSEk/+GO2MkXt7RyZY25GNdSu0ypDg3UeFIfQBdfMGICWu1HePr8zWk3/LFpTdAH6TG0oKYlFZV4crUccRg8BCOTV0Ul00tSRIXkCKt/CPSyDgy/ppNxQ2tPrbIpbBgyl/KKYFWTv+Po/6FtEFfpkEn+3k3ZIpMcfd/Tu5oP5MZGSgDH8IHFxyjZvX4+6N5RQak5PnNUH0SUvoWKf8FQFDPe5dSulbnKdTv/Ga7l1BfGG+VAqsvvFAFIHAzJWMaWeqfcZZzUAsSDEcxrSrDvSsOvmfgdK0phxaOpWStfgeX5D7jTWm4XXDVQqwLiC+I6Q8o7HyKW2X naren@Narens-MacBook-Air.local"
}

resource "aws_instance" "api_server" {
  ami           = "ami-03818140b4ac9ae2b"
  instance_type = "t2.micro"
  key_name      = aws_key_pair.api_server_key.key_name
}

