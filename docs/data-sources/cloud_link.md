---
page_title: "f5_distributed_cloud_cloud_link Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  CloudLink is used to establish private connectivity from customer network to Cloud Sites or private connectivity from F5 XC Regional Edge(RE) to customer Cloud Sites
---

# f5_distributed_cloud_cloud_link (Data Source)

CloudLink is used to establish private connectivity from customer network to Cloud Sites or private connectivity from F5 XC Regional Edge(RE) to customer Cloud Sites

## Example Usage

```hcl
data "f5_distributed_cloud_cloud_link" "example" {
  name      = "example-cloud_link"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
