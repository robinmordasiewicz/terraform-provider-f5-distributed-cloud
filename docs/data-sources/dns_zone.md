---
page_title: "f5_distributed_cloud_dns_zone Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Zone object is used for configuring Primary and Secondary DNS zones.  User configures zone 'example.com'  Status for this object will show following  * List of nameservers  User can configure D...
---

# f5_distributed_cloud_dns_zone (Data Source)

DNS Zone object is used for configuring Primary and Secondary DNS zones.  User configures zone 'example.com'  Status for this object will show following  * List of nameservers  User can configure D...

## Example Usage

```hcl
data "f5_distributed_cloud_dns_zone" "example" {
  name      = "example-dns_zone"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
