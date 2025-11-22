resource "f5xc_alert_receiver" "example" {
  name        = "my-alert-receiver"
  namespace   = "system"
  description = "Alert receiver for monitoring notifications"
  email       = "alerts@example.com"
}

resource "f5xc_alert_receiver" "slack_example" {
  name        = "slack-alerts"
  namespace   = "system"
  description = "Slack alert receiver"
  slack_url   = "https://hooks.slack.com/services/xxx/yyy/zzz"
}

resource "f5xc_alert_receiver" "pagerduty_example" {
  name         = "pagerduty-alerts"
  namespace    = "system"
  description  = "PagerDuty alert receiver"
  pagerduty_key = "your-pagerduty-integration-key"
}
