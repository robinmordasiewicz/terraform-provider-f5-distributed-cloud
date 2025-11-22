# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example provider configuration for F5 Distributed Cloud (F5 XC)
#
# The provider supports two authentication methods:
# 1. API Token Authentication (recommended for automation)
# 2. Certificate Authentication (P12 file)

terraform {
  required_providers {
    f5xc = {
      source = "robinmordasiewicz/f5-distributed-cloud"
    }
  }
}

# Option 1: API Token Authentication
# This is the recommended approach for CI/CD pipelines and automation.
#
# provider "f5xc" {
#   api_url   = "https://your-tenant.console.ves.volterra.io/api"
#   api_token = var.f5xc_api_token
# }

# Option 2: Certificate Authentication (P12 file)
# Use this when you have a P12 certificate file from F5 XC.
#
# provider "f5xc" {
#   api_url      = "https://your-tenant.console.ves.volterra.io/api"
#   p12_file     = "/path/to/certificate.p12"
#   p12_password = var.f5xc_p12_password
# }

# Option 3: Using Environment Variables (recommended for security)
# Set the following environment variables:
# - F5XC_API_URL: The API endpoint URL
# - F5XC_API_TOKEN: Your API token (for token auth)
# OR
# - F5XC_API_P12_FILE: Path to your P12 certificate file (for cert auth)
# - F5XC_API_P12_PASSWORD: Password for the P12 file (optional)
#
# Then configure the provider without explicit values:
provider "f5xc" {
  # All values will be read from environment variables
}

# Variables for sensitive credentials
variable "f5xc_api_token" {
  description = "API token for F5 XC authentication"
  type        = string
  sensitive   = true
  default     = ""
}

variable "f5xc_p12_password" {
  description = "Password for the P12 certificate file"
  type        = string
  sensitive   = true
  default     = ""
}

variable "f5xc_api_url" {
  description = "F5 XC API URL (e.g., https://your-tenant.console.ves.volterra.io/api)"
  type        = string
  default     = ""
}
