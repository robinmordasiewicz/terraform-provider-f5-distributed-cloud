---
page_title: "f5_distributed_cloud_k8s_cluster_role Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  K8s Cluster Role is K8s ClusterRole object, which represents set of permissions for user, group or service account to which this role is assigned.
---

# f5_distributed_cloud_k8s_cluster_role (Data Source)

K8s Cluster Role is K8s ClusterRole object, which represents set of permissions for user, group or service account to which this role is assigned.

## Example Usage

```hcl
data "f5_distributed_cloud_k8s_cluster_role" "example" {
  name      = "example-k8s_cluster_role"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
