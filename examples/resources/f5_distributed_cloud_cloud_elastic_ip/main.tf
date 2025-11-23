# Example configuration for f5_distributed_cloud_cloud_elastic_ip

resource "f5_distributed_cloud_cloud_elastic_ip" "example" {
  name        = "example-cloud_elastic_ip"
  namespace   = "system"
  description = "Example CloudElasticIP resource managed by Terraform"

  # Add additional configuration as needed
}
