# Example configuration for f5xc_subnet

resource "f5xc_subnet" "example" {
  name        = "example-subnet"
  namespace   = "system"
  description = "Example Subnet resource managed by Terraform"

  # Add additional configuration as needed
}
