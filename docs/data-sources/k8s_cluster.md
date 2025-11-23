---
page_title: "f5_distributed_cloud_k8s_cluster Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  K8s cluster represents the real physical K8s cluster on the site. It can be used to configure various aspect of of the K8s cluster e.g. Pod security, k8s_cluster roles, k8s_cluster role binding, Us...
---

# f5_distributed_cloud_k8s_cluster (Data Source)

K8s cluster represents the real physical K8s cluster on the site. It can be used to configure various aspect of of the K8s cluster e.g. Pod security, k8s_cluster roles, k8s_cluster role binding, Us...

## Example Usage

```hcl
data "f5_distributed_cloud_k8s_cluster" "example" {
  name      = "example-k8s_cluster"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
