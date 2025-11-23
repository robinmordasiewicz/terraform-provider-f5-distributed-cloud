---
page_title: "f5_distributed_cloud_service_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A service_policy object consists of an unordered list of predicates and a list of service policy rules. The predicates are evaluated against a set of input fields that are extracted from or derived...
---

# f5_distributed_cloud_service_policy (Resource)

A service_policy object consists of an unordered list of predicates and a list of service policy rules. The predicates are evaluated against a set of input fields that are extracted from or derived...

## Example Usage

```hcl
resource "f5_distributed_cloud_service_policy" "example" {
  name        = "example-service_policy"
  namespace   = "system"
  description = "Example ServicePolicy resource"
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

ServicePolicy can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_service_policy.example namespace/name
```
