---
page_title: "f5xc_certified_hardware Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certified Hardware object represents physical hardware or cloud instance type that will be used to instantiate a volterra software appliance instance for the F5XC sites (Customer edge site). It has...
---

# f5xc_certified_hardware (Data Source)

Certified Hardware object represents physical hardware or cloud instance type that will be used to instantiate a volterra software appliance instance for the F5XC sites (Customer edge site). It has...

## Example Usage

```hcl
data "f5xc_certified_hardware" "example" {
  name      = "example-certified_hardware"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
