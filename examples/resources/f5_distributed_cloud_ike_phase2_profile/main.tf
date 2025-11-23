# Example configuration for f5_distributed_cloud_ike_phase2_profile

resource "f5_distributed_cloud_ike_phase2_profile" "example" {
  name        = "example-ike_phase2_profile"
  namespace   = "system"
  description = "Example IkePhase2Profile resource managed by Terraform"

  # Add additional configuration as needed
}
