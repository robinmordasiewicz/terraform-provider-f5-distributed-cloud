---
page_title: "f5_distributed_cloud_report_config Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Report configuration contains the information like      List of namespaces for which the report should be generated.     Frequency of report generation.     Time at which the report should be gener...
---

# f5_distributed_cloud_report_config (Resource)

Report configuration contains the information like      List of namespaces for which the report should be generated.     Frequency of report generation.     Time at which the report should be gener...

## Example Usage

```hcl
resource "f5_distributed_cloud_report_config" "example" {
  name        = "example-report_config"
  namespace   = "system"
  description = "Example ReportConfig resource"
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

ReportConfig can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_report_config.example namespace/name
```
