---
page_title: "f5_distributed_cloud_cluster Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  cluster object represent common set of endpoints (providers of service) that can serve given route for virtual_host Cluster has list of references to Endpoint which forms the set Cluster has common...
---

# f5_distributed_cloud_cluster (Resource)

cluster object represent common set of endpoints (providers of service) that can serve given route for virtual_host Cluster has list of references to Endpoint which forms the set Cluster has common...

## Example Usage

```hcl
resource "f5_distributed_cloud_cluster" "example" {
  name        = "example-cluster"
  namespace   = "system"
  description = "Example Cluster resource"
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

Cluster can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_cluster.example namespace/name
```
