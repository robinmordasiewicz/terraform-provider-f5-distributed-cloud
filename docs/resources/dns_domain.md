---
page_title: "f5_distributed_cloud_dns_domain Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Domain object is used for delegating DNS sub domain to volterra. It can also be used to just let volterra know about a verified sub domain that can be used for different types of load balancers...
---

# f5_distributed_cloud_dns_domain (Resource)

DNS Domain object is used for delegating DNS sub domain to volterra. It can also be used to just let volterra know about a verified sub domain that can be used for different types of load balancers...

## Example Usage

```hcl
resource "f5_distributed_cloud_dns_domain" "example" {
  name        = "example-dns_domain"
  namespace   = "system"
  description = "Example DNSDomain resource"
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

DNSDomain can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_dns_domain.example namespace/name
```
