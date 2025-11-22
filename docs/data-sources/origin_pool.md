---
page_title: "f5xc_origin_pool Data Source - F5 Distributed Cloud"
subcategory: ""
description: |-
  Fetches information about an existing F5 Distributed Cloud Origin Pool.
---

# f5xc_origin_pool (Data Source)

Fetches information about an existing F5 Distributed Cloud Origin Pool.

## Example Usage

```terraform
data "f5xc_origin_pool" "existing" {
  name      = "backend-pool"
  namespace = "my-namespace"
}

output "pool_port" {
  value = data.f5xc_origin_pool.existing.port
}

output "pool_protocol" {
  value = data.f5xc_origin_pool.existing.endpoint_protocol
}

# Use in an HTTP load balancer
resource "f5xc_http_loadbalancer" "example" {
  name      = "frontend-lb"
  namespace = data.f5xc_origin_pool.existing.namespace
  domains   = ["app.example.com"]
  http_port = 80

  advertise_on_public = true

  default_pool {
    name      = data.f5xc_origin_pool.existing.name
    namespace = data.f5xc_origin_pool.existing.namespace
  }
}
```

## Argument Reference

- `name` - (Required) Name of the origin pool to look up.
- `namespace` - (Required) Namespace containing the origin pool.

## Attribute Reference

- `id` - The unique identifier of the origin pool.
- `description` - Description of the origin pool.
- `port` - Port on which origin servers listen.
- `endpoint_protocol` - Protocol for connecting to origin servers.
- `loadbalancer_algorithm` - Load balancing algorithm used.
- `health_check_port` - Port for health checks.
