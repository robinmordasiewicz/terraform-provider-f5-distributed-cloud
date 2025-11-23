# Example configuration for f5_distributed_cloud_gcp_vpc_site

resource "f5_distributed_cloud_gcp_vpc_site" "example" {
  name        = "example-gcp_vpc_site"
  namespace   = "system"
  description = "Example GCPVPCSite resource managed by Terraform"

  # Add additional configuration as needed
}
