---
page_title: "f5_distributed_cloud_user_token Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Use this API to get one time user token to connect to Web App Scanning Service
---

# f5_distributed_cloud_user_token (Data Source)

Use this API to get one time user token to connect to Web App Scanning Service

## Example Usage

```hcl
data "f5_distributed_cloud_user_token" "example" {
  name      = "example-user_token"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
