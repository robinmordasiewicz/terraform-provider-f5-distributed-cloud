---
page_title: "f5_distributed_cloud_virtual_k8s Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual K8s object exposes a Kubernetes API endpoint in the namespace that operates on all the physical Kubernetes clusters in each of the sites that are selected by the virtual-site referred to in...
---

# f5_distributed_cloud_virtual_k8s (Resource)

Virtual K8s object exposes a Kubernetes API endpoint in the namespace that operates on all the physical Kubernetes clusters in each of the sites that are selected by the virtual-site referred to in...

## Example Usage

```hcl
resource "f5_distributed_cloud_virtual_k8s" "example" {
  name        = "example-virtual_k8s"
  namespace   = "system"
  description = "Example VirtualK8S resource"
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

VirtualK8S can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_virtual_k8s.example namespace/name
```
