---
page_title: "f5_distributed_cloud_service_policy_set Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A service_policy_set object consists of an ordered list of references to service_policy objects. The policies are evaluated in the specified order against a set of input fields that are extracted f...
---

# f5_distributed_cloud_service_policy_set (Data Source)

A service_policy_set object consists of an ordered list of references to service_policy objects. The policies are evaluated in the specified order against a set of input fields that are extracted f...

## Example Usage

```hcl
data "f5_distributed_cloud_service_policy_set" "example" {
  name      = "example-service_policy_set"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
