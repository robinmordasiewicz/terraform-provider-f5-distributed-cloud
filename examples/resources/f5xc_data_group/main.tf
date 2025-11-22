# Example configuration for f5xc_data_group

resource "f5xc_data_group" "example" {
  name        = "example-data_group"
  namespace   = "system"
  description = "Example DataGroup resource managed by Terraform"

  # Add additional configuration as needed
}
