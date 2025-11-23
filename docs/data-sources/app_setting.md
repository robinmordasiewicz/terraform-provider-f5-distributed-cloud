---
page_title: "f5_distributed_cloud_app_setting Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  'App Setting' controls advanced monitoring of applications defined by 'App type'. This configuration is per namespace. There is only one App Setting per namespace. It has list of references to 'App...
---

# f5_distributed_cloud_app_setting (Data Source)

'App Setting' controls advanced monitoring of applications defined by 'App type'. This configuration is per namespace. There is only one App Setting per namespace. It has list of references to 'App...

## Example Usage

```hcl
data "f5_distributed_cloud_app_setting" "example" {
  name      = "example-app_setting"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
