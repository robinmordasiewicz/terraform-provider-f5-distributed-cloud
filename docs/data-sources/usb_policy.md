---
page_title: "f5_distributed_cloud_usb_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  USB policy is used to specify list of USB devices allowed to be attached to node.
---

# f5_distributed_cloud_usb_policy (Data Source)

USB policy is used to specify list of USB devices allowed to be attached to node.

## Example Usage

```hcl
data "f5_distributed_cloud_usb_policy" "example" {
  name      = "example-usb_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
