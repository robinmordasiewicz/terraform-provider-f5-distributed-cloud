---
page_title: "f5xc_workload Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Workload is used to configure and deploy a workload in Virtual Kubernetes. A workload can be either a service or stateful service or a batch job. Services and jobs can be deployed on regional edges...
---

# f5xc_workload (Resource)

Workload is used to configure and deploy a workload in Virtual Kubernetes. A workload can be either a service or stateful service or a batch job. Services and jobs can be deployed on regional edges...

## Example Usage

```hcl
resource "f5xc_workload" "example" {
  name        = "example-workload"
  namespace   = "system"
  description = "Example Workload resource"
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

Workload can be imported using the namespace and name:

```shell
terraform import f5xc_workload.example namespace/name
```
