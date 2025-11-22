---
page_title: "f5xc_app_setting Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  'App Setting' controls advanced monitoring of applications defined by 'App type'. This configuration is per namespace. There is only one App Setting per namespace. It has list of references to 'App...
---

# f5xc_app_setting (Resource)

'App Setting' controls advanced monitoring of applications defined by 'App type'. This configuration is per namespace. There is only one App Setting per namespace. It has list of references to 'App...

## Example Usage

```hcl
resource "f5xc_app_setting" "example" {
  name        = "example-app_setting"
  namespace   = "system"
  description = "Example AppSetting resource"
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

AppSetting can be imported using the namespace and name:

```shell
terraform import f5xc_app_setting.example namespace/name
```
