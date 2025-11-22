---
page_title: "f5xc_cloud_credentials Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Credentials object is used to give user cloud credentials to public cloud like AWS, Azure and GCP. These credentials can then be used by different features in volterra to perform cloud API(s)...
---

# f5xc_cloud_credentials (Data Source)

Cloud Credentials object is used to give user cloud credentials to public cloud like AWS, Azure and GCP. These credentials can then be used by different features in volterra to perform cloud API(s)...

## Example Usage

```hcl
data "f5xc_cloud_credentials" "example" {
  name      = "example-cloud_credentials"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
