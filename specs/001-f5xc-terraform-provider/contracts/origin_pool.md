# Resource Contract: f5xc_origin_pool

## Resource Schema

```hcl
resource "f5xc_origin_pool" "example" {
  name      = string           # Required
  namespace = string           # Required

  # Origin servers (at least one required)
  origin_servers {
    # One of: public_ip, public_name, private_ip, private_name, k8s_service, consul_service
    public_ip {
      ip = string
    }

    public_name {
      dns_name         = string
      refresh_interval = number  # Optional: seconds
    }

    private_ip {
      ip               = string
      site_locator     = object  # Site reference
      inside_network   = bool
      outside_network  = bool
    }

    labels = map(string)         # Optional
  }

  port                    = number    # Required: Backend port
  endpoint_selection      = string    # Optional: LOCAL_PREFERRED, LOCAL_ONLY, DISTRIBUTED
  loadbalancer_algorithm  = string    # Optional: ROUND_ROBIN, LEAST_ACTIVE, etc.

  # Health check reference
  healthcheck {
    name      = string
    namespace = string
  }

  # TLS settings
  use_tls {
    tls_config = object
  }
  no_tls = bool                       # Default: no TLS

  # Common
  labels      = map(string)
  annotations = map(string)
  description = string
}
```

## CRUD Operations

### Create
```
API: POST /api/config/namespaces/{namespace}/origin_pools

Preconditions:
  - Namespace exists
  - At least one origin_server defined
  - Referenced healthchecks exist
  - Port is valid (1-65535)

Postconditions:
  - Origin pool created
  - Health checks start monitoring endpoints
  - Pool available for load balancer reference

Errors:
  - 400: Invalid origin server configuration
  - 404: Referenced healthcheck not found
  - 409: Name conflict
```

### Read
```
API: GET /api/config/namespaces/{namespace}/origin_pools/{name}

Postconditions:
  - State synchronized with F5 XC
  - Health status computed attributes populated
```

### Update
```
API: PUT /api/config/namespaces/{namespace}/origin_pools/{name}

Preconditions:
  - Resource exists

Postconditions:
  - Origin servers updated
  - Load balancers referencing pool updated
  - Health checks reconfigured if changed
```

### Delete
```
API: DELETE /api/config/namespaces/{namespace}/origin_pools/{name}

Preconditions:
  - No load balancers reference this pool (or fail_if_referred=false)

Postconditions:
  - Origin pool removed
  - Health monitoring stopped

Errors:
  - 409: Pool referenced by load balancer
```

### Import
```
Import ID: {namespace}/{name}

terraform import f5xc_origin_pool.example production/api-backends
```

## Validation Rules

| Attribute | Rule |
|-----------|------|
| name | DNS-1035 format |
| port | 1-65535 |
| public_ip.ip | Valid IPv4/IPv6 address |
| public_name.dns_name | Valid hostname |
| endpoint_selection | Enum: LOCAL_PREFERRED, LOCAL_ONLY, DISTRIBUTED |
| loadbalancer_algorithm | Enum: ROUND_ROBIN, LEAST_ACTIVE, RANDOM, etc. |
