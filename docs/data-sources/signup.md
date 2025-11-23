---
page_title: "f5_distributed_cloud_signup Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Use this API to signup for F5XC service. one can signup to use volterra service as an individual/free account or as a team account more suited for enterprise customers.  for more details on what ea...
---

# f5_distributed_cloud_signup (Data Source)

Use this API to signup for F5XC service. one can signup to use volterra service as an individual/free account or as a team account more suited for enterprise customers.  for more details on what ea...

## Example Usage

```hcl
data "f5_distributed_cloud_signup" "example" {
  name      = "example-signup"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
