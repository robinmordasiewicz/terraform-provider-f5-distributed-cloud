---
page_title: "f5xc_app_firewall Resource - F5 Distributed Cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud Application Firewall (WAF) policy.
---

# f5xc_app_firewall (Resource)

Manages an F5 Distributed Cloud Application Firewall (WAF) policy.

Application Firewall policies define web application security settings including detection modes, bot protection, and blocking responses.

## Example Usage

### Basic App Firewall

```terraform
resource "f5xc_app_firewall" "example" {
  name        = "app-waf"
  namespace   = "my-namespace"
  description = "Application firewall policy"
  mode        = "monitoring"

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "medium"
  }
}
```

### Blocking Mode with Bot Protection

```terraform
resource "f5xc_app_firewall" "blocking" {
  name        = "blocking-waf"
  namespace   = "my-namespace"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    xhr_check                      = true
    cookie_protection              = true
    violation_rating               = "high"
  }

  bot_protection {
    good_bot_action       = "allow"
    malicious_bot_action  = "block"
    suspicious_bot_action = "challenge"
  }
}
```

### Custom Blocking Page

```terraform
resource "f5xc_app_firewall" "custom" {
  name                      = "custom-waf"
  namespace                 = "my-namespace"
  mode                      = "blocking"
  use_default_blocking_page = false

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "medium"
  }

  custom_blocking_page {
    blocking_page_body = "<html><body><h1>Blocked</h1></body></html>"
    response_code      = "403"
  }
}
```

## Argument Reference

- `name` - (Required) Name of the app firewall. Changing this forces a new resource.
- `namespace` - (Required) Namespace where the firewall is created. Changing this forces a new resource.
- `description` - (Optional) Description of the firewall.
- `mode` - (Optional) Firewall mode: `monitoring` or `blocking`. Defaults to `monitoring`.
- `allowed_response_codes` - (Optional) List of allowed HTTP response codes.
- `default_anonymization` - (Optional) Enable default anonymization. Defaults to true.
- `use_default_blocking_page` - (Optional) Use default blocking page. Defaults to true.
- `labels` - (Optional) Labels for the firewall.
- `detection_settings` - (Optional) Detection settings block.
- `bot_protection` - (Optional) Bot protection settings block.
- `custom_blocking_page` - (Optional) Custom blocking page configuration.

### detection_settings

- `signature_based_bot_protection` - (Optional) Enable signature-based bot protection.
- `xhr_check` - (Optional) Enable XHR request checking.
- `cookie_protection` - (Optional) Enable cookie protection.
- `enable_suppression` - (Optional) Enable violation suppression.
- `violation_rating` - (Optional) Rating threshold: `low`, `medium`, `high`.

### bot_protection

- `good_bot_action` - (Optional) Action for good bots: `allow`, `block`, `challenge`.
- `malicious_bot_action` - (Optional) Action for malicious bots.
- `suspicious_bot_action` - (Optional) Action for suspicious bots.

### custom_blocking_page

- `blocking_page_body` - (Required) HTML body content.
- `response_code` - (Optional) HTTP response code. Defaults to `403`.

## Attribute Reference

- `id` - The unique identifier of the app firewall.

## Import

App Firewalls can be imported using `namespace/name`:

```shell
terraform import f5xc_app_firewall.example my-namespace/app-waf
```
