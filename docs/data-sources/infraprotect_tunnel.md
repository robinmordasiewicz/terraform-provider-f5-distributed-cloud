---
page_title: "f5_distributed_cloud_infraprotect_tunnel Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit tunnel information
---

# f5_distributed_cloud_infraprotect_tunnel (Data Source)

DDoS transit tunnel information

## Example Usage

```hcl
data "f5_distributed_cloud_infraprotect_tunnel" "example" {
  name      = "example-infraprotect_tunnel"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
