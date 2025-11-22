---
page_title: "f5xc_subnet Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Subnet object is used to support VMs/pods with multiple interfaces, with each one in a different subnet.
---

# f5xc_subnet (Resource)

Subnet object is used to support VMs/pods with multiple interfaces, with each one in a different subnet.

## Example Usage

```hcl
resource "f5xc_subnet" "example" {
  name        = "example-subnet"
  namespace   = "system"
  description = "Example Subnet resource"
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

Subnet can be imported using the namespace and name:

```shell
terraform import f5xc_subnet.example namespace/name
```
