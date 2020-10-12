provider "aws" {
    region = var.region
    version = "~> 2.46"
}

terraform {
  backend "local" {
	
  }
}

variable "region" {

}

variable "environment" {

}

resource "aws_iam_user" "my_iam_user" {
    name = "rit_demo_iam_user_${var.environment}"
}
