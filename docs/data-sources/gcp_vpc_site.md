---
page_title: "f5xc_gcp_vpc_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  GCP VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in GCP VPC. It can be used to automatically site creation in the GCP VPC.
---

# f5xc_gcp_vpc_site (Data Source)

GCP VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in GCP VPC. It can be used to automatically site creation in the GCP VPC.

## Example Usage

```hcl
data "f5xc_gcp_vpc_site" "example" {
  name      = "example-gcp_vpc_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
