# Example configuration for f5xc_receiver

resource "f5xc_receiver" "example" {
  name        = "example-receiver"
  namespace   = "system"
  description = "Example Receiver resource managed by Terraform"

  # Add additional configuration as needed
}
