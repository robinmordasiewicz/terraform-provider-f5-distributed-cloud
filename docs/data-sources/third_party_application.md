---
page_title: "f5_distributed_cloud_third_party_application Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  View will create following child objects.  * Virtual-host * API-inventory * App-type * App-setting
---

# f5_distributed_cloud_third_party_application (Data Source)

View will create following child objects.  * Virtual-host * API-inventory * App-type * App-setting

## Example Usage

```hcl
data "f5_distributed_cloud_third_party_application" "example" {
  name      = "example-third_party_application"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
