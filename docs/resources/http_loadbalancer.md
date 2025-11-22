---
page_title: "f5xc_http_loadbalancer Resource - F5 Distributed Cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud HTTP Load Balancer.
---

# f5xc_http_loadbalancer (Resource)

Manages an F5 Distributed Cloud HTTP Load Balancer.

HTTP Load Balancers distribute incoming HTTP/HTTPS traffic across origin pools based on configured rules and policies.

## Example Usage

### Basic HTTP Load Balancer

```terraform
resource "f5xc_http_loadbalancer" "example" {
  name        = "frontend-lb"
  namespace   = "my-namespace"
  description = "Frontend load balancer"
  domains     = ["app.example.com"]
  http_port   = 80

  advertise_on_public = true

  default_pool {
    name      = f5xc_origin_pool.backend.name
    namespace = "my-namespace"
  }
}
```

### HTTPS Load Balancer

```terraform
resource "f5xc_http_loadbalancer" "https" {
  name        = "secure-lb"
  namespace   = "my-namespace"
  domains     = ["secure.example.com"]
  https_port  = 443

  advertise_on_public = true

  default_pool {
    name      = f5xc_origin_pool.backend.name
    namespace = "my-namespace"
  }

  auto_cert = true
}
```

## Argument Reference

- `name` - (Required) Name of the HTTP load balancer. Changing this forces a new resource.
- `namespace` - (Required) Namespace where the load balancer is created. Changing this forces a new resource.
- `description` - (Optional) Description of the load balancer.
- `domains` - (Required) List of domains this load balancer serves.
- `http_port` - (Optional) HTTP port to listen on. Defaults to 80.
- `https_port` - (Optional) HTTPS port to listen on.
- `advertise_on_public` - (Optional) Advertise on public network. Defaults to true.
- `auto_cert` - (Optional) Enable automatic TLS certificate management.
- `default_pool` - (Required) Default origin pool configuration.

### default_pool

- `name` - (Required) Name of the origin pool.
- `namespace` - (Required) Namespace of the origin pool.

## Attribute Reference

- `id` - The unique identifier of the HTTP load balancer.

## Import

HTTP Load Balancers can be imported using `namespace/name`:

```shell
terraform import f5xc_http_loadbalancer.example my-namespace/frontend-lb
```
