---
page_title: "f5_distributed_cloud_token Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration. Invalid tokens are refused and site with invalid to...
---

# f5_distributed_cloud_token (Data Source)

token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration. Invalid tokens are refused and site with invalid to...

## Example Usage

```hcl
data "f5_distributed_cloud_token" "example" {
  name      = "example-token"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
