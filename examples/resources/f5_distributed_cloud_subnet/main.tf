# Example configuration for f5_distributed_cloud_subnet

resource "f5_distributed_cloud_subnet" "example" {
  name        = "example-subnet"
  namespace   = "system"
  description = "Example Subnet resource managed by Terraform"

  # Add additional configuration as needed
}
