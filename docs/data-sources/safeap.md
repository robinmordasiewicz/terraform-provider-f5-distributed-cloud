---
page_title: "f5_distributed_cloud_safeap Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Use this API to interact with SAFE Account Protection endpoints. All calls which not include user interaction allow a blob as the payload and return a CubeJS Structure as the return parameter.
---

# f5_distributed_cloud_safeap (Data Source)

Use this API to interact with SAFE Account Protection endpoints. All calls which not include user interaction allow a blob as the payload and return a CubeJS Structure as the return parameter.

## Example Usage

```hcl
data "f5_distributed_cloud_safeap" "example" {
  name      = "example-safeap"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
