---
page_title: "f5xc_route Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  route object is used to configuring L7 routing decision. route is made of three things 1. Match condition for incoming request 2. Actions to take if match is true 3. Whether custom java script proc...
---

# f5xc_route (Resource)

route object is used to configuring L7 routing decision. route is made of three things 1. Match condition for incoming request 2. Actions to take if match is true 3. Whether custom java script proc...

## Example Usage

```hcl
resource "f5xc_route" "example" {
  name        = "example-route"
  namespace   = "system"
  description = "Example Route resource"
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

Route can be imported using the namespace and name:

```shell
terraform import f5xc_route.example namespace/name
```
