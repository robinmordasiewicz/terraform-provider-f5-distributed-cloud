---
page_title: "f5_distributed_cloud_k8s_pod_security_admission Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Pod security admission allows users to enforce Pod Security Standards
---

# f5_distributed_cloud_k8s_pod_security_admission (Resource)

Pod security admission allows users to enforce Pod Security Standards

## Example Usage

```hcl
resource "f5_distributed_cloud_k8s_pod_security_admission" "example" {
  name        = "example-k8s_pod_security_admission"
  namespace   = "system"
  description = "Example K8SPodSecurityAdmission resource"
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

K8SPodSecurityAdmission can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_k8s_pod_security_admission.example namespace/name
```
