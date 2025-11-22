# Example configuration for f5xc_cminstance

resource "f5xc_cminstance" "example" {
  name        = "example-cminstance"
  namespace   = "system"
  description = "Example Cminstance resource managed by Terraform"

  # Add additional configuration as needed
}
