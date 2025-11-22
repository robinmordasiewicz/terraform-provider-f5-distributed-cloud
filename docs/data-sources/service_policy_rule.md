---
page_title: "f5xc_service_policy_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A service_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 requ...
---

# f5xc_service_policy_rule (Data Source)

A service_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 requ...

## Example Usage

```hcl
data "f5xc_service_policy_rule" "example" {
  name      = "example-service_policy_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
