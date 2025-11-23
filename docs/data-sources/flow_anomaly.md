---
page_title: "f5_distributed_cloud_flow_anomaly Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Flow Anomaly.
---

# f5_distributed_cloud_flow_anomaly (Data Source)

Flow Anomaly.

## Example Usage

```hcl
data "f5_distributed_cloud_flow_anomaly" "example" {
  name      = "example-flow_anomaly"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
