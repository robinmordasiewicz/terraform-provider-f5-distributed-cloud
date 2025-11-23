---
page_title: "f5_distributed_cloud_cminstance Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  cminsatnce object can be used to enable connectivity between ce site and bigip central manager.
---

# f5_distributed_cloud_cminstance (Data Source)

cminsatnce object can be used to enable connectivity between ce site and bigip central manager.

## Example Usage

```hcl
data "f5_distributed_cloud_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
