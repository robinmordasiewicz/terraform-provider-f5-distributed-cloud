# Data Model: F5 Distributed Cloud Terraform Provider

**Branch**: `001-f5xc-terraform-provider` | **Date**: 2025-11-21 | **Spec**: [spec.md](./spec.md)

## Entity Overview

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           Provider Configuration                         │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌──────────────┐   │
│  │   api_url   │  │  api_token  │  │  cert_file  │  │   key_file   │   │
│  └─────────────┘  └─────────────┘  └─────────────┘  └──────────────┘   │
└─────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────┐
│                              Namespace                                   │
│  (Organizational container for all resources)                           │
└─────────────────────────────────────────────────────────────────────────┘
        │                    │                    │                    │
        ▼                    ▼                    ▼                    ▼
┌───────────────┐  ┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│  HTTP Load    │  │  TCP Load     │  │  DNS Load     │  │  CDN Load     │
│  Balancer     │  │  Balancer     │  │  Balancer     │  │  Balancer     │
└───────────────┘  └───────────────┘  └───────────────┘  └───────────────┘
        │                    │
        ▼                    ▼
┌───────────────┐  ┌───────────────┐
│  Origin Pool  │◄─┤  Health Check │
└───────────────┘  └───────────────┘
        │
        ▼
┌───────────────┐
│  App Firewall │
└───────────────┘

┌─────────────────────────────────────────────────────────────────────────┐
│                          Cloud Infrastructure                            │
├───────────────┬───────────────┬───────────────┬─────────────────────────┤
│  Cloud        │  AWS VPC      │  Azure VNet   │  GCP VPC               │
│  Credentials  │  Site         │  Site         │  Site                  │
└───────────────┴───────────────┴───────────────┴─────────────────────────┘
```

## Core Entities

### 1. Provider Configuration

```hcl
provider "f5xc" {
  api_url    = "https://<tenant>.console.ves.volterra.io/api"

  # Option 1: Certificate Authentication
  cert_file  = "/path/to/cert.pem"
  key_file   = "/path/to/key.pem"

  # Option 2: API Token Authentication
  api_token  = "your-api-token"

  # Optional
  timeout    = "30s"
}
```

**Schema**:
| Attribute | Type | Required | Description |
|-----------|------|----------|-------------|
| `api_url` | string | Yes | Tenant API endpoint URL |
| `cert_file` | string | No* | Path to client certificate PEM file |
| `key_file` | string | No* | Path to client key PEM file |
| `api_token` | string | No* | API token for authentication |
| `timeout` | string | No | HTTP request timeout (default: 30s) |

*Either certificate (cert_file + key_file) or api_token required

---

### 2. Namespace

**Resource**: `f5xc_namespace`

The fundamental organizational unit in F5 XC. All other resources belong to a namespace.

```hcl
resource "f5xc_namespace" "example" {
  name        = "production"
  description = "Production environment namespace"

  labels = {
    environment = "production"
    team        = "platform"
  }

  annotations = {
    "app.kubernetes.io/managed-by" = "terraform"
  }
}
```

**Schema**:
| Attribute | Type | Required | Computed | Description |
|-----------|------|----------|----------|-------------|
| `id` | string | No | Yes | Resource identifier |
| `name` | string | Yes | No | Namespace name (DNS-1035 format) |
| `description` | string | No | No | Human-readable description |
| `labels` | map(string) | No | No | Key-value labels for organization |
| `annotations` | map(string) | No | No | Key-value annotations for metadata |
| `disable` | bool | No | No | Administratively disable namespace |

---

### 3. HTTP Load Balancer

**Resource**: `f5xc_http_loadbalancer`

Distributes HTTP/HTTPS traffic to backend origin pools.

```hcl
resource "f5xc_http_loadbalancer" "example" {
  name      = "api-gateway"
  namespace = f5xc_namespace.example.name

  domains = ["api.example.com"]

  http {
    port = 80
  }

  # Or HTTPS
  https_auto_cert {
    port = 443
  }

  default_route {
    origin_pools {
      pool {
        name      = f5xc_origin_pool.backend.name
        namespace = f5xc_namespace.example.name
      }
      weight = 100
    }
  }

  app_firewall {
    name      = f5xc_app_firewall.waf.name
    namespace = f5xc_namespace.example.name
  }

  advertise_on_public_default_vip = true
}
```

**Schema** (key attributes):
| Attribute | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | Yes | Load balancer name |
| `namespace` | string | Yes | Parent namespace |
| `domains` | list(string) | Yes | Domain names for the load balancer |
| `http` | object | No* | HTTP listener configuration |
| `https_auto_cert` | object | No* | HTTPS with auto-managed certificates |
| `https` | object | No* | HTTPS with custom certificates |
| `default_route` | object | No | Default routing configuration |
| `routes` | list(object) | No | Custom routing rules |
| `app_firewall` | object | No | WAF policy reference |
| `rate_limiter` | object | No | Rate limiting configuration |
| `advertise_on_public_default_vip` | bool | No | Advertise on F5 global network |

---

### 4. TCP Load Balancer

**Resource**: `f5xc_tcp_loadbalancer`

Distributes Layer 4 TCP traffic.

```hcl
resource "f5xc_tcp_loadbalancer" "example" {
  name      = "database-lb"
  namespace = f5xc_namespace.example.name

  domains = ["db.example.com"]

  listen_port = 5432

  origin_pools {
    pool {
      name      = f5xc_origin_pool.database.name
      namespace = f5xc_namespace.example.name
    }
    weight = 100
  }

  advertise_on_public_default_vip = true
}
```

---

### 5. Origin Pool

**Resource**: `f5xc_origin_pool`

Defines backend server endpoints for load balancers.

```hcl
resource "f5xc_origin_pool" "example" {
  name      = "api-backends"
  namespace = f5xc_namespace.example.name

  origin_servers {
    public_ip {
      ip = "192.168.1.10"
    }

    labels = {
      zone = "us-east-1a"
    }
  }

  origin_servers {
    public_name {
      dns_name = "backend.internal.example.com"
    }
  }

  port              = 8080
  endpoint_selection = "LOCAL_PREFERRED"

  healthcheck {
    name      = f5xc_healthcheck.http.name
    namespace = f5xc_namespace.example.name
  }

  loadbalancer_algorithm = "ROUND_ROBIN"
}
```

**Schema**:
| Attribute | Type | Required | Description |
|-----------|------|----------|-------------|
| `name` | string | Yes | Origin pool name |
| `namespace` | string | Yes | Parent namespace |
| `origin_servers` | list(object) | Yes | Backend server definitions |
| `port` | number | Yes | Backend port |
| `endpoint_selection` | string | No | Endpoint selection strategy |
| `healthcheck` | object | No | Health check reference |
| `loadbalancer_algorithm` | string | No | Load balancing algorithm |

---

### 6. App Firewall

**Resource**: `f5xc_app_firewall`

Web Application Firewall (WAF) policy configuration.

```hcl
resource "f5xc_app_firewall" "example" {
  name      = "standard-waf"
  namespace = f5xc_namespace.example.name

  detection_settings {
    detection_mode   = "DETECTION"
    signature_stages = ["PRODUCTION"]
  }

  blocking_page {
    response_code = 403
    body          = "Access Denied"
  }

  allow_all_response_codes = false
  default_anonymization    = true
}
```

---

### 7. Cloud Credentials

**Resource**: `f5xc_cloud_credentials`

Cloud provider credentials for site provisioning.

```hcl
# AWS Credentials - IAM Role
resource "f5xc_cloud_credentials" "aws" {
  name      = "aws-prod"
  namespace = "system"

  aws_assume_role {
    role_arn     = "arn:aws:iam::123456789012:role/f5xc-site-role"
    session_name = "f5xc-terraform"
  }
}

# AWS Credentials - Access Key
resource "f5xc_cloud_credentials" "aws_key" {
  name      = "aws-key"
  namespace = "system"

  aws_secret_key {
    access_key = "AKIAIOSFODNN7EXAMPLE"
    secret_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  }
}

# Azure Credentials
resource "f5xc_cloud_credentials" "azure" {
  name      = "azure-prod"
  namespace = "system"

  azure_client_secret {
    subscription_id = "sub-id"
    tenant_id       = "tenant-id"
    client_id       = "client-id"
    client_secret   = "secret"
  }
}

# GCP Credentials
resource "f5xc_cloud_credentials" "gcp" {
  name      = "gcp-prod"
  namespace = "system"

  gcp_cred_file {
    credential_file = file("gcp-service-account.json")
  }
}
```

---

### 8. AWS VPC Site

**Resource**: `f5xc_aws_vpc_site`

F5 XC site deployment in AWS VPC.

```hcl
resource "f5xc_aws_vpc_site" "example" {
  name      = "aws-site-useast1"
  namespace = "system"

  aws_region = "us-east-1"

  aws_cred {
    name      = f5xc_cloud_credentials.aws.name
    namespace = "system"
  }

  vpc {
    new_vpc {
      name_tag  = "f5xc-vpc"
      primary_ipv4 = "10.0.0.0/16"
    }
  }

  ingress_egress_gw {
    aws_certified_hw = "aws-byol-voltmesh"

    az_nodes {
      aws_az_name = "us-east-1a"

      inside_subnet {
        subnet_param {
          ipv4 = "10.0.1.0/24"
        }
      }

      outside_subnet {
        subnet_param {
          ipv4 = "10.0.2.0/24"
        }
      }
    }
  }
}
```

---

### 9. DNS Load Balancer

**Resource**: `f5xc_dns_load_balancer`

DNS-based global load balancing.

```hcl
resource "f5xc_dns_load_balancer" "example" {
  name      = "global-dns-lb"
  namespace = f5xc_namespace.example.name

  domains = ["app.example.com"]

  dns_zone_ref {
    name      = "example-zone"
    namespace = f5xc_namespace.example.name
  }

  record_type = "A"
  ttl         = 300

  dns_lb_pool {
    pool {
      name      = f5xc_dns_lb_pool.primary.name
      namespace = f5xc_namespace.example.name
    }
    priority = 1
  }
}
```

---

## Common Nested Types

### ObjectRef
Reference to another F5 XC object:
```hcl
{
  name      = "object-name"
  namespace = "object-namespace"
  tenant    = "tenant-name"  # Optional
}
```

### Labels
Standard Kubernetes-style labels:
```hcl
labels = {
  "app.kubernetes.io/name"       = "my-app"
  "app.kubernetes.io/component"  = "frontend"
}
```

### Metadata (common to all resources)
```hcl
{
  name        = string      # Required, DNS-1035 format
  namespace   = string      # Required for most resources
  labels      = map(string) # Optional
  annotations = map(string) # Optional
  description = string      # Optional, max 1200 bytes
  disable     = bool        # Optional, default false
}
```

---

## Relationship Matrix

| Resource | Depends On | Referenced By |
|----------|-----------|---------------|
| Namespace | - | All other resources |
| Origin Pool | Namespace, Health Check | HTTP LB, TCP LB |
| App Firewall | Namespace | HTTP LB |
| HTTP Load Balancer | Namespace, Origin Pool, App Firewall | - |
| TCP Load Balancer | Namespace, Origin Pool | - |
| DNS Load Balancer | Namespace, DNS Zone, DNS LB Pool | - |
| CDN Load Balancer | Namespace, Origin Pool | - |
| Cloud Credentials | - | AWS/Azure/GCP Sites |
| AWS VPC Site | Cloud Credentials | - |
| Azure VNet Site | Cloud Credentials | - |
| GCP VPC Site | Cloud Credentials | - |

---

## State Management

### Resource Identifiers
All resources use a composite identifier:
```
{namespace}/{name}
```

### Import Syntax
```bash
terraform import f5xc_http_loadbalancer.example production/api-gateway
terraform import f5xc_namespace.example production
```

### Computed Attributes
All resources compute:
- `id` - Resource identifier
- `uid` - F5 XC internal unique ID
- `creation_timestamp` - Resource creation time
- `modification_timestamp` - Last modification time
