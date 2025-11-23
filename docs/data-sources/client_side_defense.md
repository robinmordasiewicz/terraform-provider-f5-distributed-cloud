---
page_title: "f5_distributed_cloud_client_side_defense Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Custom handler in Client-Side Defense microservice will forward request(s) to backend API(s)
---

# f5_distributed_cloud_client_side_defense (Data Source)

Custom handler in Client-Side Defense microservice will forward request(s) to backend API(s)

## Example Usage

```hcl
data "f5_distributed_cloud_client_side_defense" "example" {
  name      = "example-client_side_defense"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
