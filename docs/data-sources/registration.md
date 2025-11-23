---
page_title: "f5_distributed_cloud_registration Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  registration API(s) are used by Customer edge site to register itself with volterra edge cloud. Every node in given site is represented by it's registration. registration must have unique hostname ...
---

# f5_distributed_cloud_registration (Data Source)

registration API(s) are used by Customer edge site to register itself with volterra edge cloud. Every node in given site is represented by it's registration. registration must have unique hostname ...

## Example Usage

```hcl
data "f5_distributed_cloud_registration" "example" {
  name      = "example-registration"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
