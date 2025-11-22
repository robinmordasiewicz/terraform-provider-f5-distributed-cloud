---
page_title: "f5xc_secret_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A Secret Policy defines who gets access to a secret. A secret_policy object consists of an unordered list of predicates and a list of secret policy rules. The predicates are evaluated against a set...
---

# f5xc_secret_policy (Resource)

A Secret Policy defines who gets access to a secret. A secret_policy object consists of an unordered list of predicates and a list of secret policy rules. The predicates are evaluated against a set...

## Example Usage

```hcl
resource "f5xc_secret_policy" "example" {
  name        = "example-secret_policy"
  namespace   = "system"
  description = "Example SecretPolicy resource"
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

SecretPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_secret_policy.example namespace/name
```
