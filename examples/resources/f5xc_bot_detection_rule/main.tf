# Example configuration for f5xc_bot_detection_rule

resource "f5xc_bot_detection_rule" "example" {
  name        = "example-bot_detection_rule"
  namespace   = "system"
  description = "Example BotDetectionRule resource managed by Terraform"

  # Add additional configuration as needed
}
