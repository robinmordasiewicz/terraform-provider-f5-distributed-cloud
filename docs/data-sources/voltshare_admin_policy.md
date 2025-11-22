---
page_title: "f5xc_voltshare_admin_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  VoltShare Admin Policy object is an admin level policy object that restricts all secrets encrypted by author's team/tenant shared via F5XC VoltShare. There can be a maximum of *ONE* VoltShare Admin...
---

# f5xc_voltshare_admin_policy (Data Source)

VoltShare Admin Policy object is an admin level policy object that restricts all secrets encrypted by author's team/tenant shared via F5XC VoltShare. There can be a maximum of *ONE* VoltShare Admin...

## Example Usage

```hcl
data "f5xc_voltshare_admin_policy" "example" {
  name      = "example-voltshare_admin_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
