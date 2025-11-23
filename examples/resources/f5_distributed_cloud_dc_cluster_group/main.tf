# Example configuration for f5_distributed_cloud_dc_cluster_group

resource "f5_distributed_cloud_dc_cluster_group" "example" {
  name        = "example-dc_cluster_group"
  namespace   = "system"
  description = "Example DcClusterGroup resource managed by Terraform"

  # Add additional configuration as needed
}
