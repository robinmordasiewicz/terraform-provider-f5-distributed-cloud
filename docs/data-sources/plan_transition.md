---
page_title: "f5xc_plan_transition Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Package plan transition is responsible for storing and managing requests to move from one billing plan to another.
---

# f5xc_plan_transition (Data Source)

Package plan transition is responsible for storing and managing requests to move from one billing plan to another.

## Example Usage

```hcl
data "f5xc_plan_transition" "example" {
  name      = "example-plan_transition"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
