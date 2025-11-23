# Authentication Guide

This guide covers the authentication options available in the F5 Distributed Cloud Terraform provider.

## Overview

The provider supports two authentication methods:
1. **API Token** - Recommended for automation and CI/CD pipelines
2. **P12 Certificate** - Certificate-based authentication using PKCS#12 files

## API Token Authentication

API tokens are the recommended authentication method for most use cases.

### Generating an API Token

1. Log in to the F5 Distributed Cloud Console
2. Navigate to **Administration** > **Personal Management** > **Credentials**
3. Click **Create Credentials**
4. Select **API Token** as the credential type
5. Set an expiration date and save the token securely

### Using API Token

#### In Provider Configuration

```hcl
provider "f5_distributed_cloud" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.api_token
}

variable "api_token" {
  description = "F5 XC API Token"
  type        = string
  sensitive   = true
}
```

#### Via Environment Variable

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
```

```hcl
# Provider will use environment variables
provider "f5_distributed_cloud" {}
```

### API Token Security Best Practices

- Never commit tokens to version control
- Use short-lived tokens when possible
- Rotate tokens regularly
- Use environment variables or secret management systems
- Limit token scope to required permissions

## P12 Certificate Authentication

Certificate authentication uses PKCS#12 (.p12) files containing your client certificate and private key.

### Obtaining a P12 Certificate

1. Log in to the F5 Distributed Cloud Console
2. Navigate to **Administration** > **Personal Management** > **Credentials**
3. Click **Create Credentials**
4. Select **API Certificate** as the credential type
5. Download the P12 file and securely store the password

### Using P12 Certificate

#### In Provider Configuration

```hcl
provider "f5_distributed_cloud" {
  api_url      = "https://your-tenant.console.ves.volterra.io/api"
  p12_file     = var.p12_file
  p12_password = var.p12_password
}

variable "p12_file" {
  description = "Path to P12 certificate file"
  type        = string
}

variable "p12_password" {
  description = "Password for P12 certificate"
  type        = string
  sensitive   = true
}
```

#### Via Environment Variables

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_P12_FILE="/path/to/certificate.p12"
export F5XC_API_P12_PASSWORD="your-password"
```

```hcl
provider "f5_distributed_cloud" {}
```

### P12 Certificate Security Best Practices

- Store P12 files in secure locations with restricted permissions
- Never commit P12 files to version control
- Use a strong password for the P12 file
- Rotate certificates before expiration
- Consider using a secrets manager for CI/CD pipelines

## Authentication Priority

The provider evaluates authentication configuration in the following order:

1. **Explicit configuration** - Values set directly in the provider block
2. **Environment variables** - `F5XC_*` environment variables

If both API token and P12 certificate are configured, the provider will use the API token.

## CI/CD Integration

### GitHub Actions

```yaml
name: Terraform

on:
  push:
    branches: [main]

jobs:
  terraform:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Terraform
        uses: hashicorp/setup-terraform@v3

      - name: Terraform Init
        run: terraform init
        env:
          F5XC_API_URL: ${{ secrets.F5XC_API_URL }}
          F5XC_API_TOKEN: ${{ secrets.F5XC_API_TOKEN }}

      - name: Terraform Plan
        run: terraform plan
        env:
          F5XC_API_URL: ${{ secrets.F5XC_API_URL }}
          F5XC_API_TOKEN: ${{ secrets.F5XC_API_TOKEN }}
```

### GitLab CI

```yaml
variables:
  F5XC_API_URL: ${F5XC_API_URL}
  F5XC_API_TOKEN: ${F5XC_API_TOKEN}

stages:
  - validate
  - plan
  - apply

terraform_plan:
  stage: plan
  script:
    - terraform init
    - terraform plan -out=plan.tfplan
  artifacts:
    paths:
      - plan.tfplan

terraform_apply:
  stage: apply
  script:
    - terraform apply plan.tfplan
  when: manual
  only:
    - main
```

### Jenkins

```groovy
pipeline {
    agent any

    environment {
        F5XC_API_URL     = credentials('f5-distributed-cloud-api-url')
        F5XC_API_TOKEN   = credentials('f5-distributed-cloud-api-token')
    }

    stages {
        stage('Terraform Init') {
            steps {
                sh 'terraform init'
            }
        }
        stage('Terraform Plan') {
            steps {
                sh 'terraform plan'
            }
        }
        stage('Terraform Apply') {
            steps {
                sh 'terraform apply -auto-approve'
            }
        }
    }
}
```

## Secrets Management

### HashiCorp Vault

```hcl
data "vault_generic_secret" "f5_distributed_cloud" {
  path = "secret/f5-distributed-cloud"
}

provider "f5_distributed_cloud" {
  api_url   = data.vault_generic_secret.f5_distributed_cloud.data["api_url"]
  api_token = data.vault_generic_secret.f5_distributed_cloud.data["api_token"]
}
```

### AWS Secrets Manager

```hcl
data "aws_secretsmanager_secret_version" "f5_distributed_cloud" {
  secret_id = "f5-distributed-cloud-credentials"
}

locals {
  f5_distributed_cloud_creds = jsondecode(data.aws_secretsmanager_secret_version.f5_distributed_cloud.secret_string)
}

provider "f5_distributed_cloud" {
  api_url   = local.f5_distributed_cloud_creds.api_url
  api_token = local.f5_distributed_cloud_creds.api_token
}
```

## Troubleshooting

### Common Authentication Errors

#### "Missing Authentication Credentials"

Ensure either API token or P12 certificate is configured:
```bash
export F5XC_API_TOKEN="your-token"
# or
export F5XC_API_P12_FILE="/path/to/cert.p12"
export F5XC_API_P12_PASSWORD="password"
```

#### "Invalid API Token"

- Verify the token hasn't expired
- Check the token was copied correctly (no extra whitespace)
- Ensure the token has appropriate permissions

#### "Certificate Error"

- Verify the P12 file path is correct
- Check the P12 password is correct
- Ensure the certificate hasn't expired

#### "Missing API URL"

Set the API URL:
```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
```

### Debugging Authentication

Enable debug logging to troubleshoot authentication issues:

```bash
export TF_LOG=DEBUG
terraform plan
```

This will show detailed information about the authentication process.
