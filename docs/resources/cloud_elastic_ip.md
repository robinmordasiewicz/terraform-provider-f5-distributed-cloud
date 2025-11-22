---
page_title: "f5xc_cloud_elastic_ip Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Elastic IP object represents a cloud elastic IP address that are created for a cloud site
---

# f5xc_cloud_elastic_ip (Resource)

Cloud Elastic IP object represents a cloud elastic IP address that are created for a cloud site

## Example Usage

```hcl
resource "f5xc_cloud_elastic_ip" "example" {
  name        = "example-cloud_elastic_ip"
  namespace   = "system"
  description = "Example CloudElasticIP resource"
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

CloudElasticIP can be imported using the namespace and name:

```shell
terraform import f5xc_cloud_elastic_ip.example namespace/name
```
