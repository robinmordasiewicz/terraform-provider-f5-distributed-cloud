---
page_title: "f5_distributed_cloud_k8s_cluster_role_binding Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cluster role binding allows administrator to assign cluster wide cluster role to a users, groups or service accounts
---

# f5_distributed_cloud_k8s_cluster_role_binding (Resource)

Cluster role binding allows administrator to assign cluster wide cluster role to a users, groups or service accounts

## Example Usage

```hcl
resource "f5_distributed_cloud_k8s_cluster_role_binding" "example" {
  name        = "example-k8s_cluster_role_binding"
  namespace   = "system"
  description = "Example K8SClusterRoleBinding resource"
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

K8SClusterRoleBinding can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_k8s_cluster_role_binding.example namespace/name
```
