---
page_title: "f5_distributed_cloud_origin_pool Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck
---

# f5_distributed_cloud_origin_pool (Resource)

Origin pool is a view to create cluster and endpoints that can be used in HTTP loadbalancer or TCP loadbalancer  It will create following child objects  * cluster * endpoints * healthcheck

## Example Usage

```hcl
resource "f5_distributed_cloud_origin_pool" "example" {
  name        = "example-origin_pool"
  namespace   = "system"
  description = "Example OriginPool resource"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace where the resource will be created.
- `description` - (Optional) Description of the resource.
- `labels` - (Optional) Labels for the resource.
- `annotations` - (Optional) Annotations for the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.

## Import

OriginPool can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_origin_pool.example namespace/name
```
