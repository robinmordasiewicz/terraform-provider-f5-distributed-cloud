---
page_title: "f5xc_report_config Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Report configuration contains the information like      List of namespaces for which the report should be generated.     Frequency of report generation.     Time at which the report should be gener...
---

# f5xc_report_config (Data Source)

Report configuration contains the information like      List of namespaces for which the report should be generated.     Frequency of report generation.     Time at which the report should be gener...

## Example Usage

```hcl
data "f5xc_report_config" "example" {
  name      = "example-report_config"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
