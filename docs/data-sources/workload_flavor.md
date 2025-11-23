---
page_title: "f5_distributed_cloud_workload_flavor Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Workload flavor is used to assign CPU, memory, and storage resources to workloads.
---

# f5_distributed_cloud_workload_flavor (Data Source)

Workload flavor is used to assign CPU, memory, and storage resources to workloads.

## Example Usage

```hcl
data "f5_distributed_cloud_workload_flavor" "example" {
  name      = "example-workload_flavor"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
