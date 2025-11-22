---
page_title: "f5xc_workload Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Workload is used to configure and deploy a workload in Virtual Kubernetes. A workload can be either a service or stateful service or a batch job. Services and jobs can be deployed on regional edges...
---

# f5xc_workload (Data Source)

Workload is used to configure and deploy a workload in Virtual Kubernetes. A workload can be either a service or stateful service or a batch job. Services and jobs can be deployed on regional edges...

## Example Usage

```hcl
data "f5xc_workload" "example" {
  name      = "example-workload"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
