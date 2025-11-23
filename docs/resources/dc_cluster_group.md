---
page_title: "f5_distributed_cloud_dc_cluster_group Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A DC Cluster Group represents a collection of sites that can directly communicate with each other using the underlay network.
---

# f5_distributed_cloud_dc_cluster_group (Resource)

A DC Cluster Group represents a collection of sites that can directly communicate with each other using the underlay network.

## Example Usage

```hcl
resource "f5_distributed_cloud_dc_cluster_group" "example" {
  name        = "example-dc_cluster_group"
  namespace   = "system"
  description = "Example DcClusterGroup resource"
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

DcClusterGroup can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_dc_cluster_group.example namespace/name
```
