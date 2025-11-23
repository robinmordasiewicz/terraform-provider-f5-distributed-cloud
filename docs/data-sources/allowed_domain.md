---
page_title: "f5_distributed_cloud_allowed_domain Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Allowed Domain Object defines which domains will be allowed by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for enterpr...
---

# f5_distributed_cloud_allowed_domain (Data Source)

Allowed Domain Object defines which domains will be allowed by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for enterpr...

## Example Usage

```hcl
data "f5_distributed_cloud_allowed_domain" "example" {
  name      = "example-allowed_domain"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
