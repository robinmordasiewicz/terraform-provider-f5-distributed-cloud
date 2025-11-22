---
page_title: "f5xc_origin_pool Resource - F5 Distributed Cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud Origin Pool.
---

# f5xc_origin_pool (Resource)

Manages an F5 Distributed Cloud Origin Pool.

Origin Pools define groups of backend servers that serve application traffic. They are used by HTTP Load Balancers to route traffic to appropriate backends.

## Example Usage

### Basic Origin Pool with DNS Name

```terraform
resource "f5xc_origin_pool" "example" {
  name              = "backend-pool"
  namespace         = "my-namespace"
  description       = "Backend server pool"
  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_name {
      dns_name = "backend.example.com"
    }
  }
}
```

### Origin Pool with Public IPs

```terraform
resource "f5xc_origin_pool" "public" {
  name              = "public-ip-pool"
  namespace         = "my-namespace"
  port              = 443
  endpoint_protocol = "https"

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }

  origin_servers {
    public_ip {
      ip = "192.0.2.2"
    }
  }
}
```

### Origin Pool with Private IPs on Site

```terraform
resource "f5xc_origin_pool" "private" {
  name              = "private-pool"
  namespace         = "my-namespace"
  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    private_ip {
      ip             = "10.0.1.10"
      site_name      = "my-site"
      inside_network = true
    }
  }
}
```

## Argument Reference

- `name` - (Required) Name of the origin pool. Changing this forces a new resource.
- `namespace` - (Required) Namespace where the origin pool is created. Changing this forces a new resource.
- `description` - (Optional) Description of the origin pool.
- `port` - (Required) Port on which origin servers listen.
- `endpoint_protocol` - (Optional) Protocol for connecting to origins: `http` or `https`. Defaults to `http`.
- `loadbalancer_algorithm` - (Optional) Load balancing algorithm: `ROUND_ROBIN`, `LEAST_REQUEST`, `RANDOM`, `SOURCE_IP_HASH`. Defaults to `ROUND_ROBIN`.
- `health_check_port` - (Optional) Port for health checks. Defaults to origin port.
- `origin_servers` - (Required) List of origin server configurations.

### origin_servers

- `public_ip` - (Optional) Public IP configuration.
  - `ip` - (Required) IPv4 or IPv6 address.
- `private_ip` - (Optional) Private IP configuration.
  - `ip` - (Required) IPv4 or IPv6 address.
  - `site_name` - (Required) Name of the site.
  - `inside_network` - (Optional) Use inside network.
- `public_name` - (Optional) DNS name configuration.
  - `dns_name` - (Required) Fully qualified domain name.
- `labels` - (Optional) Labels for the origin server.

## Attribute Reference

- `id` - The unique identifier of the origin pool.

## Import

Origin Pools can be imported using `namespace/name`:

```shell
terraform import f5xc_origin_pool.example my-namespace/backend-pool
```
