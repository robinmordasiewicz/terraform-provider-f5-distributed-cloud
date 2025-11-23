---
page_title: "f5_distributed_cloud_sensitive_data_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The sensitive_data_policy is a policy defined by the user to discover the relevant compliances and data types to the user. the user can disabled predefined data types, and add custom data types. th...
---

# f5_distributed_cloud_sensitive_data_policy (Resource)

The sensitive_data_policy is a policy defined by the user to discover the relevant compliances and data types to the user. the user can disabled predefined data types, and add custom data types. th...

## Example Usage

```hcl
resource "f5_distributed_cloud_sensitive_data_policy" "example" {
  name        = "example-sensitive_data_policy"
  namespace   = "system"
  description = "Example SensitiveDataPolicy resource"
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

SensitiveDataPolicy can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_sensitive_data_policy.example namespace/name
```
