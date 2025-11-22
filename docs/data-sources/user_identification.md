---
page_title: "f5xc_user_identification Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A user_identification object consists of an ordered list of rules. The rules are evaluated against the input fields that are extracted from an request API to determine a user identifier. The identi...
---

# f5xc_user_identification (Data Source)

A user_identification object consists of an ordered list of rules. The rules are evaluated against the input fields that are extracted from an request API to determine a user identifier. The identi...

## Example Usage

```hcl
data "f5xc_user_identification" "example" {
  name      = "example-user_identification"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
