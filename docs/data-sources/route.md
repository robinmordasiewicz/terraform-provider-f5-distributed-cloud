---
page_title: "f5xc_route Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  route object is used to configuring L7 routing decision. route is made of three things 1. Match condition for incoming request 2. Actions to take if match is true 3. Whether custom java script proc...
---

# f5xc_route (Data Source)

route object is used to configuring L7 routing decision. route is made of three things 1. Match condition for incoming request 2. Actions to take if match is true 3. Whether custom java script proc...

## Example Usage

```hcl
data "f5xc_route" "example" {
  name      = "example-route"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
