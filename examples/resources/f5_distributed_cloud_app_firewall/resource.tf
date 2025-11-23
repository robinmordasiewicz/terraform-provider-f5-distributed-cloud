# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: Basic App Firewall in Monitoring Mode
resource "f5_distributed_cloud_app_firewall" "basic" {
  name        = "basic-app-firewall"
  namespace   = "my-namespace"
  description = "Basic Application Firewall in monitoring mode"
  mode        = "monitoring"

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "medium"
  }
}

# Example: App Firewall in Blocking Mode
resource "f5_distributed_cloud_app_firewall" "blocking" {
  name        = "blocking-app-firewall"
  namespace   = "my-namespace"
  description = "Application Firewall in blocking mode"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    xhr_check                      = true
    cookie_protection              = true
    enable_suppression             = false
    violation_rating               = "high"
  }

  default_anonymization      = true
  use_default_blocking_page = true
}

# Example: App Firewall with Bot Protection
resource "f5_distributed_cloud_app_firewall" "with_bot_protection" {
  name        = "bot-protected-firewall"
  namespace   = "my-namespace"
  description = "Application Firewall with comprehensive bot protection"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    xhr_check                      = true
    cookie_protection              = true
    violation_rating               = "medium"
  }

  bot_protection {
    good_bot_action       = "allow"
    malicious_bot_action  = "block"
    suspicious_bot_action = "challenge"
  }
}

# Example: App Firewall with Custom Blocking Page
resource "f5_distributed_cloud_app_firewall" "custom_blocking" {
  name        = "custom-blocking-firewall"
  namespace   = "my-namespace"
  description = "Application Firewall with custom blocking page"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "medium"
  }

  use_default_blocking_page = false

  custom_blocking_page {
    blocking_page_body = <<-EOT
      <!DOCTYPE html>
      <html>
      <head>
        <title>Access Denied</title>
        <style>
          body { font-family: Arial, sans-serif; text-align: center; padding: 50px; }
          h1 { color: #d9534f; }
        </style>
      </head>
      <body>
        <h1>Access Denied</h1>
        <p>Your request has been blocked by our security system.</p>
        <p>If you believe this is an error, please contact support.</p>
      </body>
      </html>
    EOT
    response_code      = "403"
  }
}

# Example: App Firewall with Labels
resource "f5_distributed_cloud_app_firewall" "labeled" {
  name        = "labeled-app-firewall"
  namespace   = "my-namespace"
  description = "Application Firewall with labels"
  mode        = "monitoring"

  detection_settings {
    signature_based_bot_protection = true
    violation_rating               = "low"
  }

  labels = {
    environment = "production"
    team        = "security"
    compliance  = "pci-dss"
  }
}

# Example: Comprehensive App Firewall Configuration
resource "f5_distributed_cloud_app_firewall" "comprehensive" {
  name        = "comprehensive-app-firewall"
  namespace   = "my-namespace"
  description = "Comprehensive Application Firewall configuration"
  mode        = "blocking"

  detection_settings {
    signature_based_bot_protection = true
    xhr_check                      = true
    cookie_protection              = true
    enable_suppression             = true
    violation_rating               = "high"
  }

  bot_protection {
    good_bot_action       = "allow"
    malicious_bot_action  = "block"
    suspicious_bot_action = "challenge"
  }

  allowed_response_codes    = ["200", "201", "204", "301", "302", "400", "401", "403", "404"]
  default_anonymization     = true
  use_default_blocking_page = true

  labels = {
    environment = "production"
    team        = "security"
  }
}
