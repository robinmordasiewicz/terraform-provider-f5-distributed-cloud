# Example configuration for f5_distributed_cloud_ike_phase1_profile

resource "f5_distributed_cloud_ike_phase1_profile" "example" {
  name        = "example-ike_phase1_profile"
  namespace   = "system"
  description = "Example IkePhase1Profile resource managed by Terraform"

  # Add additional configuration as needed
}
