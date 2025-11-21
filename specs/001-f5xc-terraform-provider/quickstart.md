# Quickstart Guide: F5 Distributed Cloud Terraform Provider

## Prerequisites

1. **F5 Distributed Cloud Account**: Active tenant with API access enabled
2. **Terraform**: Version 1.0 or later
3. **API Credentials**: Either:
   - API certificate (P12 file converted to PEM)
   - API token generated from F5 XC Console

## Installation

### From Terraform Registry (when published)

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "f5networks/f5xc"
      version = "~> 0.1"
    }
  }
}
```

### Local Development Build

```bash
# Clone repository
git clone https://github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud.git
cd terraform-provider-f5-distributed-cloud

# Build provider
go build -o terraform-provider-f5xc

# Install locally
mkdir -p ~/.terraform.d/plugins/local/f5networks/f5xc/0.1.0/darwin_arm64
mv terraform-provider-f5xc ~/.terraform.d/plugins/local/f5networks/f5xc/0.1.0/darwin_arm64/

# Configure Terraform to use local build
cat > ~/.terraformrc << EOF
provider_installation {
  dev_overrides {
    "f5networks/f5xc" = "~/.terraform.d/plugins/local/f5networks/f5xc/0.1.0/darwin_arm64"
  }
  direct {}
}
EOF
```

## Configuration

### Provider Setup

Create `main.tf`:

```hcl
terraform {
  required_providers {
    f5xc = {
      source  = "f5networks/f5xc"
      version = "~> 0.1"
    }
  }
}

# Option 1: Certificate Authentication
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  cert_file = "/path/to/your-cert.pem"
  key_file  = "/path/to/your-key.pem"
}

# Option 2: Token Authentication
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.f5xc_api_token
}
```

### Environment Variables (Recommended)

```bash
export F5XC_API_URL="https://your-tenant.console.ves.volterra.io/api"
export F5XC_API_TOKEN="your-api-token"
# Or for certificate auth:
export F5XC_CERT_FILE="/path/to/cert.pem"
export F5XC_KEY_FILE="/path/to/key.pem"
```

## First Resources

### Create a Namespace

```hcl
resource "f5xc_namespace" "demo" {
  name        = "demo-namespace"
  description = "Demo namespace for quickstart"

  labels = {
    environment = "demo"
    managed-by  = "terraform"
  }
}
```

### Create an Origin Pool

```hcl
resource "f5xc_origin_pool" "backend" {
  name      = "demo-backend"
  namespace = f5xc_namespace.demo.name

  origin_servers {
    public_name {
      dns_name = "httpbin.org"
    }
  }

  port                   = 443
  loadbalancer_algorithm = "ROUND_ROBIN"

  use_tls {
    tls_config {
      default_security = true
    }
  }
}
```

### Create an HTTP Load Balancer

```hcl
resource "f5xc_http_loadbalancer" "demo" {
  name      = "demo-lb"
  namespace = f5xc_namespace.demo.name

  domains = ["demo.example.com"]

  https_auto_cert {
    port = 443
  }

  default_route {
    origin_pools {
      pool {
        name      = f5xc_origin_pool.backend.name
        namespace = f5xc_namespace.demo.name
      }
      weight = 100
    }
  }

  advertise_on_public_default_vip = true
}

output "lb_cname" {
  description = "CNAME to configure in DNS"
  value       = f5xc_http_loadbalancer.demo.cname
}
```

## Deploy

```bash
# Initialize Terraform
terraform init

# Preview changes
terraform plan

# Apply configuration
terraform apply

# View outputs
terraform output lb_cname
```

## Verify

1. Check F5 XC Console for created resources
2. Configure DNS CNAME record pointing to output value
3. Test connectivity:
   ```bash
   curl https://demo.example.com/get
   ```

## Clean Up

```bash
terraform destroy
```

## Common Patterns

### Using Data Sources

```hcl
# Query existing namespace
data "f5xc_namespace" "existing" {
  name = "existing-namespace"
}

# Use in resource
resource "f5xc_http_loadbalancer" "example" {
  namespace = data.f5xc_namespace.existing.name
  # ...
}
```

### Importing Existing Resources

```bash
# Import existing namespace
terraform import f5xc_namespace.existing existing-namespace

# Import existing load balancer
terraform import f5xc_http_loadbalancer.existing namespace/lb-name
```

### Migrating from Proprietary Provider

1. Export current state:
   ```bash
   terraform state pull > backup.tfstate
   ```

2. Update provider source in configuration:
   ```hcl
   # Before
   provider "volterra" { ... }
   resource "volterra_namespace" "example" { ... }

   # After
   provider "f5xc" { ... }
   resource "f5xc_namespace" "example" { ... }
   ```

3. Import resources:
   ```bash
   terraform import f5xc_namespace.example my-namespace
   ```

4. Verify with plan:
   ```bash
   terraform plan
   ```

## Next Steps

- [Full Data Model](./data-model.md)
- [API Contracts](./contracts/)
- [Feature Specification](./spec.md)
- [Technical Research](./research.md)
