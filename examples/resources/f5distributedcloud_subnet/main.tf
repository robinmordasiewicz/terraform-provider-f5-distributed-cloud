# Example configuration for f5distributedcloud_subnet

resource "f5distributedcloud_subnet" "example" {
  name        = "example-subnet"
  namespace   = "system"
  description = "Example Subnet resource managed by Terraform"

  # Add additional configuration as needed
}
