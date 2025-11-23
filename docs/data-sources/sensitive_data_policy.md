---
page_title: "f5_distributed_cloud_sensitive_data_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The sensitive_data_policy is a policy defined by the user to discover the relevant compliances and data types to the user. the user can disabled predefined data types, and add custom data types. th...
---

# f5_distributed_cloud_sensitive_data_policy (Data Source)

The sensitive_data_policy is a policy defined by the user to discover the relevant compliances and data types to the user. the user can disabled predefined data types, and add custom data types. th...

## Example Usage

```hcl
data "f5_distributed_cloud_sensitive_data_policy" "example" {
  name      = "example-sensitive_data_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
