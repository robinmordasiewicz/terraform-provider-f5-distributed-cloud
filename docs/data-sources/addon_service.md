---
page_title: "f5_distributed_cloud_addon_service Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Basic unit of logical representation of a F5XC service. An addon service can be self serviced, partially managed or fully managed depending upon the activation requirement. A configuration object b...
---

# f5_distributed_cloud_addon_service (Data Source)

Basic unit of logical representation of a F5XC service. An addon service can be self serviced, partially managed or fully managed depending upon the activation requirement. A configuration object b...

## Example Usage

```hcl
data "f5_distributed_cloud_addon_service" "example" {
  name      = "example-addon_service"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
