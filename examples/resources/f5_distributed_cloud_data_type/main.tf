# Example configuration for f5_distributed_cloud_data_type

resource "f5_distributed_cloud_data_type" "example" {
  name        = "example-data_type"
  namespace   = "system"
  description = "Example DataType resource managed by Terraform"

  # Add additional configuration as needed
}
