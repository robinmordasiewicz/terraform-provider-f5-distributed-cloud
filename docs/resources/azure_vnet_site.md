---
page_title: "f5xc_azure_vnet_site Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Azure VNet site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Azure VNet. It can be used to automatically site creation in the Azure VNet.  Vi...
---

# f5xc_azure_vnet_site (Resource)

Azure VNet site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Azure VNet. It can be used to automatically site creation in the Azure VNet.  Vi...

## Example Usage

```hcl
resource "f5xc_azure_vnet_site" "example" {
  name        = "example-azure_vnet_site"
  namespace   = "system"
  description = "Example AzureVNETSite resource"
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

AzureVNETSite can be imported using the namespace and name:

```shell
terraform import f5xc_azure_vnet_site.example namespace/name
```
