---
page_title: "f5_distributed_cloud_authentication Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Authentication Object contains authentication specific config . This includes authentication type config, config specific to each authentication type and cookie specific paramters
---

# f5_distributed_cloud_authentication (Data Source)

Authentication Object contains authentication specific config . This includes authentication type config, config specific to each authentication type and cookie specific paramters

## Example Usage

```hcl
data "f5_distributed_cloud_authentication" "example" {
  name      = "example-authentication"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
