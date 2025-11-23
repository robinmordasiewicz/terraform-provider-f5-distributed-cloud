---
page_title: "f5_distributed_cloud_http_loadbalancer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  HTTP Load Balancer view defines a required parameters that can be used in CRUD, to create and manage HTTP Load Balancer. It can be used to create HTTP Load Balancer and HTTPS Load Balancer.  View w...
---

# f5_distributed_cloud_http_loadbalancer (Data Source)

HTTP Load Balancer view defines a required parameters that can be used in CRUD, to create and manage HTTP Load Balancer. It can be used to create HTTP Load Balancer and HTTPS Load Balancer.  View w...

## Example Usage

```hcl
data "f5_distributed_cloud_http_loadbalancer" "example" {
  name      = "example-http_loadbalancer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
