---
page_title: "f5xc_workload_flavor Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Workload flavor is used to assign CPU, memory, and storage resources to workloads.
---

# f5xc_workload_flavor (Resource)

Workload flavor is used to assign CPU, memory, and storage resources to workloads.

## Example Usage

```hcl
resource "f5xc_workload_flavor" "example" {
  name        = "example-workload_flavor"
  namespace   = "system"
  description = "Example WorkloadFlavor resource"
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

WorkloadFlavor can be imported using the namespace and name:

```shell
terraform import f5xc_workload_flavor.example namespace/name
```
