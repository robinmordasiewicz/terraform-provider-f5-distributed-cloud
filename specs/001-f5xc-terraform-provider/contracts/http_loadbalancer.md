# Resource Contract: f5_distributed_cloud_http_loadbalancer

## Resource Schema (Simplified)

```hcl
resource "f5_distributed_cloud_http_loadbalancer" "example" {
  name      = string           # Required
  namespace = string           # Required

  domains   = list(string)     # Required: At least one domain

  # Listener: One of http, https, https_auto_cert
  http {
    port = number              # Optional: default 80
  }

  https_auto_cert {
    port              = number # Optional: default 443
    add_hsts          = bool   # Optional
    tls_config        = object # Optional
  }

  https {
    port              = number # Optional: default 443
    tls_certificates  = list(object) # Required for https
  }

  # Routing
  default_route {
    origin_pools {
      pool {
        name      = string
        namespace = string
      }
      weight    = number       # Optional: default 1
      priority  = number       # Optional: default 1
    }
  }

  routes = list(object)        # Optional: Custom routes

  # Security
  app_firewall {
    name      = string
    namespace = string
  }

  rate_limiter = object        # Optional

  # Advertisement
  advertise_on_public_default_vip = bool  # Optional
  advertise_custom = object    # Optional

  # Common
  labels      = map(string)
  annotations = map(string)
  description = string
  disable     = bool
}
```

## CRUD Operations

### Create
```
API: POST /api/config/namespaces/{namespace}/http_loadbalancers

Preconditions:
  - Namespace exists
  - All referenced origin_pools exist
  - All referenced app_firewalls exist
  - domains are valid hostnames

Postconditions:
  - HTTP LB created and advertised
  - Virtual host created
  - Routes configured
  - Security policies applied

Errors:
  - 400: Invalid configuration
  - 404: Referenced resource not found
  - 409: Name conflict
```

### Read
```
API: GET /api/config/namespaces/{namespace}/http_loadbalancers/{name}

Postconditions:
  - Full state synchronized
  - Computed status fields populated
```

### Update
```
API: PUT /api/config/namespaces/{namespace}/http_loadbalancers/{name}

Preconditions:
  - Resource exists
  - name and namespace cannot change (ForceNew)

Postconditions:
  - Configuration updated
  - Traffic handling updated without interruption
```

### Delete
```
API: DELETE /api/config/namespaces/{namespace}/http_loadbalancers/{name}

Postconditions:
  - HTTP LB removed
  - Associated virtual hosts removed
  - Traffic stops being served
```

### Import
```
Import ID: {namespace}/{name}

terraform import f5_distributed_cloud_http_loadbalancer.example production/api-gateway
```

## Validation Rules

| Attribute | Rule |
|-----------|------|
| name | DNS-1035 format |
| namespace | Must exist |
| domains | Valid hostnames, non-empty list |
| http.port | 1-65535 |
| https.port | 1-65535 |
| weight | 0-100 |
| priority | 1-128 |

## Force New Triggers

These attributes trigger resource recreation:
- `name`
- `namespace`
