---
page_title: "f5_distributed_cloud_aws_vpc_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  AWS VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC. It can be used to automatically site creation in the AWS VPC.  View will c...
---

# f5_distributed_cloud_aws_vpc_site (Data Source)

AWS VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC. It can be used to automatically site creation in the AWS VPC.  View will c...

## Example Usage

```hcl
data "f5_distributed_cloud_aws_vpc_site" "example" {
  name      = "example-aws_vpc_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
