---
page_title: "f5xc_gcp_vpc_site Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  GCP VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in GCP VPC. It can be used to automatically site creation in the GCP VPC.
---

# f5xc_gcp_vpc_site (Resource)

GCP VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in GCP VPC. It can be used to automatically site creation in the GCP VPC.

## Example Usage

```hcl
resource "f5xc_gcp_vpc_site" "example" {
  name        = "example-gcp_vpc_site"
  namespace   = "system"
  description = "Example GCPVPCSite resource"
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

GCPVPCSite can be imported using the namespace and name:

```shell
terraform import f5xc_gcp_vpc_site.example namespace/name
```
