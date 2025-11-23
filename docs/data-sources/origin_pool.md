---
page_title: "f5_distributed_cloud_origin_pool Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck
---

# f5_distributed_cloud_origin_pool (Data Source)

Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck

## Example Usage

```hcl
data "f5_distributed_cloud_origin_pool" "example" {
  name      = "example-origin_pool"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
