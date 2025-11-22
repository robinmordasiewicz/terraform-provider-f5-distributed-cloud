---
page_title: "f5xc_cloud_link Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  CloudLink is used to establish private connectivity from customer network to Cloud Sites or private connectivity from F5 XC Regional Edge(RE) to customer Cloud Sites
---

# f5xc_cloud_link (Resource)

CloudLink is used to establish private connectivity from customer network to Cloud Sites or private connectivity from F5 XC Regional Edge(RE) to customer Cloud Sites

## Example Usage

```hcl
resource "f5xc_cloud_link" "example" {
  name        = "example-cloud_link"
  namespace   = "system"
  description = "Example CloudLink resource"
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

CloudLink can be imported using the namespace and name:

```shell
terraform import f5xc_cloud_link.example namespace/name
```
