---
page_title: "f5_distributed_cloud_fast_acl Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Fast ACL provides destination and specifies rules to protect the site from denial of service attacks.   - It is specified in terms of five tuple of the packet {destination ip, destination port, sou...
---

# f5_distributed_cloud_fast_acl (Resource)

Fast ACL provides destination and specifies rules to protect the site from denial of service attacks.   - It is specified in terms of five tuple of the packet {destination ip, destination port, sou...

## Example Usage

```hcl
resource "f5_distributed_cloud_fast_acl" "example" {
  name        = "example-fast_acl"
  namespace   = "system"
  description = "Example FastACL resource"
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

FastACL can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_fast_acl.example namespace/name
```
