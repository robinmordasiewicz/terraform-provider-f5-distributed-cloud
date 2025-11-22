---
page_title: "f5xc_user_identification Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A user_identification object consists of an ordered list of rules. The rules are evaluated against the input fields that are extracted from an request API to determine a user identifier. The identi...
---

# f5xc_user_identification (Resource)

A user_identification object consists of an ordered list of rules. The rules are evaluated against the input fields that are extracted from an request API to determine a user identifier. The identi...

## Example Usage

```hcl
resource "f5xc_user_identification" "example" {
  name        = "example-user_identification"
  namespace   = "system"
  description = "Example UserIdentification resource"
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

UserIdentification can be imported using the namespace and name:

```shell
terraform import f5xc_user_identification.example namespace/name
```
