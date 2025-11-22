---
page_title: "f5xc_secret_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A Secret Policy defines who gets access to a secret. A secret_policy object consists of an unordered list of predicates and a list of secret policy rules. The predicates are evaluated against a set...
---

# f5xc_secret_policy (Data Source)

A Secret Policy defines who gets access to a secret. A secret_policy object consists of an unordered list of predicates and a list of secret policy rules. The predicates are evaluated against a set...

## Example Usage

```hcl
data "f5xc_secret_policy" "example" {
  name      = "example-secret_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
