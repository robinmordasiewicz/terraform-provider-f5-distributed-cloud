---
page_title: "f5_distributed_cloud_k8s_pod_security_admission Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Pod security admission allows users to enforce Pod Security Standards
---

# f5_distributed_cloud_k8s_pod_security_admission (Data Source)

Pod security admission allows users to enforce Pod Security Standards

## Example Usage

```hcl
data "f5_distributed_cloud_k8s_pod_security_admission" "example" {
  name      = "example-k8s_pod_security_admission"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
