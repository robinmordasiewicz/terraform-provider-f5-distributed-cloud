---
page_title: "f5xc_azure_vnet_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Azure VNet site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Azure VNet. It can be used to automatically site creation in the Azure VNet.  Vi...
---

# f5xc_azure_vnet_site (Data Source)

Azure VNet site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Azure VNet. It can be used to automatically site creation in the Azure VNet.  Vi...

## Example Usage

```hcl
data "f5xc_azure_vnet_site" "example" {
  name      = "example-azure_vnet_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
