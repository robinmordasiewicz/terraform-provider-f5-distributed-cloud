# Example: Cloud Credentials
# This example creates cloud credentials for AWS in F5 Distributed Cloud

# AWS Cloud Credentials
resource "f5_distributed_cloud_cloud_credentials" "aws" {
  name        = "my-aws-credentials"
  namespace   = "system"
  description = "AWS credentials for cloud site deployments"

  aws_secret_key {
    access_key = var.aws_access_key
    secret_key = var.aws_secret_key
  }
}

# Azure Cloud Credentials (alternative example)
resource "f5_distributed_cloud_cloud_credentials" "azure" {
  name        = "my-azure-credentials"
  namespace   = "system"
  description = "Azure credentials for VNET site deployments"

  azure_client_secret {
    subscription_id = var.azure_subscription_id
    tenant_id       = var.azure_tenant_id
    client_id       = var.azure_client_id
    client_secret   = var.azure_client_secret
  }
}

# GCP Cloud Credentials (alternative example)
resource "f5_distributed_cloud_cloud_credentials" "gcp" {
  name        = "my-gcp-credentials"
  namespace   = "system"
  description = "GCP credentials for VPC site deployments"

  gcp_credentials {
    credential_file = var.gcp_credentials_json
  }
}

# Variables
variable "aws_access_key" {
  description = "AWS Access Key"
  type        = string
  sensitive   = true
}

variable "aws_secret_key" {
  description = "AWS Secret Key"
  type        = string
  sensitive   = true
}

variable "azure_subscription_id" {
  description = "Azure Subscription ID"
  type        = string
}

variable "azure_tenant_id" {
  description = "Azure Tenant ID"
  type        = string
}

variable "azure_client_id" {
  description = "Azure Client ID"
  type        = string
}

variable "azure_client_secret" {
  description = "Azure Client Secret"
  type        = string
  sensitive   = true
}

variable "gcp_credentials_json" {
  description = "GCP Service Account Credentials JSON"
  type        = string
  sensitive   = true
}

# Outputs
output "aws_credentials_name" {
  description = "Name of the AWS cloud credentials"
  value       = f5_distributed_cloud_cloud_credentials.aws.name
}

output "azure_credentials_name" {
  description = "Name of the Azure cloud credentials"
  value       = f5_distributed_cloud_cloud_credentials.azure.name
}

output "gcp_credentials_name" {
  description = "Name of the GCP cloud credentials"
  value       = f5_distributed_cloud_cloud_credentials.gcp.name
}
