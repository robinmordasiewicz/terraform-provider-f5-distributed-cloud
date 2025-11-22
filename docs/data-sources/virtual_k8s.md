---
page_title: "f5xc_virtual_k8s Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual K8s object exposes a Kubernetes API endpoint in the namespace that operates on all the physical Kubernetes clusters in each of the sites that are selected by the virtual-site referred to in...
---

# f5xc_virtual_k8s (Data Source)

Virtual K8s object exposes a Kubernetes API endpoint in the namespace that operates on all the physical Kubernetes clusters in each of the sites that are selected by the virtual-site referred to in...

## Example Usage

```hcl
data "f5xc_virtual_k8s" "example" {
  name      = "example-virtual_k8s"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
