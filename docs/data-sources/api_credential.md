---
page_title: "f5xc_api_credential Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  F5XC supports 2 variation of credentials -  1. My Credentials or API credentials 2. Service Credentials  Credentials created via My credential or API credential inherits same user that of the creat...
---

# f5xc_api_credential (Data Source)

F5XC supports 2 variation of credentials -  1. My Credentials or API credentials 2. Service Credentials  Credentials created via My credential or API credential inherits same user that of the creat...

## Example Usage

```hcl
data "f5xc_api_credential" "example" {
  name      = "example-api_credential"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
