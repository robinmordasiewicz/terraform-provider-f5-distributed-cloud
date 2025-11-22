---
page_title: "f5xc_report Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Report configuration contains the information like      Time at which the report was last sent to object store.     Report ID.
---

# f5xc_report (Data Source)

Report configuration contains the information like      Time at which the report was last sent to object store.     Report ID.

## Example Usage

```hcl
data "f5xc_report" "example" {
  name      = "example-report"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
