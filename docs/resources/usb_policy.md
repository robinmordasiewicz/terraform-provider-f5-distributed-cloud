---
page_title: "f5xc_usb_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  USB policy is used to specify list of USB devices allowed to be attached to node.
---

# f5xc_usb_policy (Resource)

USB policy is used to specify list of USB devices allowed to be attached to node.

## Example Usage

```hcl
resource "f5xc_usb_policy" "example" {
  name        = "example-usb_policy"
  namespace   = "system"
  description = "Example UsbPolicy resource"
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

UsbPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_usb_policy.example namespace/name
```
