---
page_title: "f5xc_oidc_provider Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  F5XC Identity supports identity brokering and third-party identity can be added to provide Single Sign-On (SSO) login for user to access tenant via VoltConsole.  Using Volterra's OIDC Provider conf...
---

# f5xc_oidc_provider (Data Source)

F5XC Identity supports identity brokering and third-party identity can be added to provide Single Sign-On (SSO) login for user to access tenant via VoltConsole.  Using Volterra's OIDC Provider conf...

## Example Usage

```hcl
data "f5xc_oidc_provider" "example" {
  name      = "example-oidc_provider"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
