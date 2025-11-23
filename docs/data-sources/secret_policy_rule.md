---
page_title: "f5_distributed_cloud_secret_policy_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Secret Policy Rule defines a rule controlling access to a secret. A secret_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set o...
---

# f5_distributed_cloud_secret_policy_rule (Data Source)

Secret Policy Rule defines a rule controlling access to a secret. A secret_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set o...

## Example Usage

```hcl
data "f5_distributed_cloud_secret_policy_rule" "example" {
  name      = "example-secret_policy_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
