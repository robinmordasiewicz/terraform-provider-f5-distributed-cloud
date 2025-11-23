---
page_title: "f5_distributed_cloud_service_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A service_policy object consists of an unordered list of predicates and a list of service policy rules. The predicates are evaluated against a set of input fields that are extracted from or derived...
---

# f5_distributed_cloud_service_policy (Data Source)

A service_policy object consists of an unordered list of predicates and a list of service policy rules. The predicates are evaluated against a set of input fields that are extracted from or derived...

## Example Usage

```hcl
data "f5_distributed_cloud_service_policy" "example" {
  name      = "example-service_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
