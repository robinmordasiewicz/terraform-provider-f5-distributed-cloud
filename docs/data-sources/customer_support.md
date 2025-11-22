---
page_title: "f5xc_customer_support Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Handles creation and listing of support issues (by tenant and user). Currently supported operations are: * create - to create a support request * list - to query for all issues created by a custome...
---

# f5xc_customer_support (Data Source)

Handles creation and listing of support issues (by tenant and user). Currently supported operations are: * create - to create a support request * list - to query for all issues created by a custome...

## Example Usage

```hcl
data "f5xc_customer_support" "example" {
  name      = "example-customer_support"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
