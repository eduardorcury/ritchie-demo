provider "aws" {
    region = "us-east-1"
    version = "~> 2.46"
}

resource "aws_iam_user" "my_iam_user" {
    name = "rit_demo_iam_user"
}
