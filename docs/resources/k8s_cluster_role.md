---
page_title: "f5_distributed_cloud_k8s_cluster_role Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  K8s Cluster Role is K8s ClusterRole object, which represents set of permissions for user, group or service account to which this role is assigned.
---

# f5_distributed_cloud_k8s_cluster_role (Resource)

K8s Cluster Role is K8s ClusterRole object, which represents set of permissions for user, group or service account to which this role is assigned.

## Example Usage

```hcl
resource "f5_distributed_cloud_k8s_cluster_role" "example" {
  name        = "example-k8s_cluster_role"
  namespace   = "system"
  description = "Example K8SClusterRole resource"
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

K8SClusterRole can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_k8s_cluster_role.example namespace/name
```
