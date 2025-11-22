---
page_title: "f5xc_k8s_cluster Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  K8s cluster represents the real physical K8s cluster on the site. It can be used to configure various aspect of of the K8s cluster e.g. Pod security, k8s_cluster roles, k8s_cluster role binding, Us...
---

# f5xc_k8s_cluster (Resource)

K8s cluster represents the real physical K8s cluster on the site. It can be used to configure various aspect of of the K8s cluster e.g. Pod security, k8s_cluster roles, k8s_cluster role binding, Us...

## Example Usage

```hcl
resource "f5xc_k8s_cluster" "example" {
  name        = "example-k8s_cluster"
  namespace   = "system"
  description = "Example K8SCluster resource"
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

K8SCluster can be imported using the namespace and name:

```shell
terraform import f5xc_k8s_cluster.example namespace/name
```
