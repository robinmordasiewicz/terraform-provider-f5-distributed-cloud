---
page_title: "f5_distributed_cloud_mitigated_domain Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Mitigated Domain Object defines which domains will be mitigated by Client-Side Defense. Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...
---

# f5_distributed_cloud_mitigated_domain (Data Source)

Mitigated Domain Object defines which domains will be mitigated by Client-Side Defense. Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...

## Example Usage

```hcl
data "f5_distributed_cloud_mitigated_domain" "example" {
  name      = "example-mitigated_domain"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
