terraform {
  cloud {
    # your organization here
    organization = "danieladamce"
    workspaces {
      # your workspace here
      name = "go-lambda" 
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }
  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-east-1"
}