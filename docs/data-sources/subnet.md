---
page_title: "f5_distributed_cloud_subnet Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Subnet object is used to support VMs/pods with multiple interfaces, with each one in a different subnet.
---

# f5_distributed_cloud_subnet (Data Source)

Subnet object is used to support VMs/pods with multiple interfaces, with each one in a different subnet.

## Example Usage

```hcl
data "f5_distributed_cloud_subnet" "example" {
  name      = "example-subnet"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
