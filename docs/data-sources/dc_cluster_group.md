---
page_title: "f5_distributed_cloud_dc_cluster_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A DC Cluster Group represents a collection of sites that can directly communicate with each other using the underlay network.
---

# f5_distributed_cloud_dc_cluster_group (Data Source)

A DC Cluster Group represents a collection of sites that can directly communicate with each other using the underlay network.

## Example Usage

```hcl
data "f5_distributed_cloud_dc_cluster_group" "example" {
  name      = "example-dc_cluster_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
