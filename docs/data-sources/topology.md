---
page_title: "f5_distributed_cloud_topology Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  APIs to get topology of all the resources associated/connected to a site such as other CEs (Customer Edge), REs (Regional Edge), VPCs (Virtual Private Cloud) networks, etc., and the associated metr...
---

# f5_distributed_cloud_topology (Data Source)

APIs to get topology of all the resources associated/connected to a site such as other CEs (Customer Edge), REs (Regional Edge), VPCs (Virtual Private Cloud) networks, etc., and the associated metr...

## Example Usage

```hcl
data "f5_distributed_cloud_topology" "example" {
  name      = "example-topology"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
