---
page_title: "f5_distributed_cloud_aws_tgw_site Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  AWS TGW site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC. It can be used to automatically site creation in the AWS TGW.  View will c...
---

# f5_distributed_cloud_aws_tgw_site (Resource)

AWS TGW site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC. It can be used to automatically site creation in the AWS TGW.  View will c...

## Example Usage

```hcl
resource "f5_distributed_cloud_aws_tgw_site" "example" {
  name        = "example-aws_tgw_site"
  namespace   = "system"
  description = "Example AWSTGWSite resource"
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

AWSTGWSite can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_aws_tgw_site.example namespace/name
```
