---
page_title: "f5_distributed_cloud_terraform_parameters Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  View Terraform Parameters is set of parameters that are used by terraform scripts  to instantiate view objects external to volterra
---

# f5_distributed_cloud_terraform_parameters (Data Source)

View Terraform Parameters is set of parameters that are used by terraform scripts  to instantiate view objects external to volterra

## Example Usage

```hcl
data "f5_distributed_cloud_terraform_parameters" "example" {
  name      = "example-terraform_parameters"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
