# Example configuration for f5_distributed_cloud_data_group

resource "f5_distributed_cloud_data_group" "example" {
  name        = "example-data_group"
  namespace   = "system"
  description = "Example DataGroup resource managed by Terraform"

  # Add additional configuration as needed
}
