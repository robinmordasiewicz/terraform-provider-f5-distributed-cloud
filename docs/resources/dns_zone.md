---
page_title: "f5_distributed_cloud_dns_zone Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Zone object is used for configuring Primary and Secondary DNS zones.  User configures zone 'example.com'  Status for this object will show following  * List of nameservers  User can configure D...
---

# f5_distributed_cloud_dns_zone (Resource)

DNS Zone object is used for configuring Primary and Secondary DNS zones.  User configures zone 'example.com'  Status for this object will show following  * List of nameservers  User can configure D...

## Example Usage

```hcl
resource "f5_distributed_cloud_dns_zone" "example" {
  name        = "example-dns_zone"
  namespace   = "system"
  description = "Example DNSZone resource"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace where the resource will be created.
- `description` - (Optional) Description of the resource.
- `labels` - (Optional) Labels for the resource.
- `annotations` - (Optional) Annotations for the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.

## Import

DNSZone can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_dns_zone.example namespace/name
```
