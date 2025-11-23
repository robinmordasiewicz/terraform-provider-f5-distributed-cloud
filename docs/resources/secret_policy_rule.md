---
page_title: "f5_distributed_cloud_secret_policy_rule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Secret Policy Rule defines a rule controlling access to a secret. A secret_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set o...
---

# f5_distributed_cloud_secret_policy_rule (Resource)

Secret Policy Rule defines a rule controlling access to a secret. A secret_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set o...

## Example Usage

```hcl
resource "f5_distributed_cloud_secret_policy_rule" "example" {
  name        = "example-secret_policy_rule"
  namespace   = "system"
  description = "Example SecretPolicyRule resource"
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

SecretPolicyRule can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_secret_policy_rule.example namespace/name
```
