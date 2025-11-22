---
page_title: "f5xc_voltshare_admin_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  VoltShare Admin Policy object is an admin level policy object that restricts all secrets encrypted by author's team/tenant shared via F5XC VoltShare. There can be a maximum of *ONE* VoltShare Admin...
---

# f5xc_voltshare_admin_policy (Resource)

VoltShare Admin Policy object is an admin level policy object that restricts all secrets encrypted by author's team/tenant shared via F5XC VoltShare. There can be a maximum of *ONE* VoltShare Admin...

## Example Usage

```hcl
resource "f5xc_voltshare_admin_policy" "example" {
  name        = "example-voltshare_admin_policy"
  namespace   = "system"
  description = "Example VoltshareAdminPolicy resource"
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

VoltshareAdminPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_voltshare_admin_policy.example namespace/name
```
