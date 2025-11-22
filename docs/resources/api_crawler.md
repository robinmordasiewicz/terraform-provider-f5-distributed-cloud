---
page_title: "f5xc_api_crawler Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  This is the api crawler type
---

# f5xc_api_crawler (Resource)

This is the api crawler type

## Example Usage

```hcl
resource "f5xc_api_crawler" "example" {
  name        = "example-api_crawler"
  namespace   = "system"
  description = "Example APICrawler resource"
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

APICrawler can be imported using the namespace and name:

```shell
terraform import f5xc_api_crawler.example namespace/name
```
