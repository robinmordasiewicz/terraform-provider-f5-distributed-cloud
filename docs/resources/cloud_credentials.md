---
page_title: "f5xc_cloud_credentials Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Credentials object is used to give user cloud credentials to public cloud like AWS, Azure and GCP. These credentials can then be used by different features in volterra to perform cloud API(s)...
---

# f5xc_cloud_credentials (Resource)

Cloud Credentials object is used to give user cloud credentials to public cloud like AWS, Azure and GCP. These credentials can then be used by different features in volterra to perform cloud API(s)...

## Example Usage

```hcl
resource "f5xc_cloud_credentials" "example" {
  name        = "example-cloud_credentials"
  namespace   = "system"
  description = "Example CloudCredentials resource"
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

CloudCredentials can be imported using the namespace and name:

```shell
terraform import f5xc_cloud_credentials.example namespace/name
```
