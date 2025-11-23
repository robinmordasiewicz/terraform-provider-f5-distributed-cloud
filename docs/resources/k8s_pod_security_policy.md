---
page_title: "f5_distributed_cloud_k8s_pod_security_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Pod Security Policies enable fine-grained authorization of pod creation and updates.
---

# f5_distributed_cloud_k8s_pod_security_policy (Resource)

Pod Security Policies enable fine-grained authorization of pod creation and updates.

## Example Usage

```hcl
resource "f5_distributed_cloud_k8s_pod_security_policy" "example" {
  name        = "example-k8s_pod_security_policy"
  namespace   = "system"
  description = "Example K8SPodSecurityPolicy resource"
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

K8SPodSecurityPolicy can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_k8s_pod_security_policy.example namespace/name
```
