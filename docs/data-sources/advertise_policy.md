---
page_title: "f5xc_advertise_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. Basic concept is a service is made available to consumers in a given ...
---

# f5xc_advertise_policy (Data Source)

advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. Basic concept is a service is made available to consumers in a given ...

## Example Usage

```hcl
data "f5xc_advertise_policy" "example" {
  name      = "example-advertise_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
