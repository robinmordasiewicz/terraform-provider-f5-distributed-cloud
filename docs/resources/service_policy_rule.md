---
page_title: "f5xc_service_policy_rule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A service_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 requ...
---

# f5xc_service_policy_rule (Resource)

A service_policy_rule object consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 requ...

## Example Usage

```hcl
resource "f5xc_service_policy_rule" "example" {
  name        = "example-service_policy_rule"
  namespace   = "system"
  description = "Example ServicePolicyRule resource"
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

ServicePolicyRule can be imported using the namespace and name:

```shell
terraform import f5xc_service_policy_rule.example namespace/name
```
