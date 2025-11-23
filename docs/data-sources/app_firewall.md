---
page_title: "f5_distributed_cloud_app_firewall Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  WAF Configuration
---

# f5_distributed_cloud_app_firewall (Data Source)

WAF Configuration

## Example Usage

```hcl
data "f5_distributed_cloud_app_firewall" "example" {
  name      = "example-app_firewall"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
