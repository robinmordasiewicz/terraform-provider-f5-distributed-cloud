# Example configuration for f5xc_workload

resource "f5xc_workload" "example" {
  name        = "example-workload"
  namespace   = "system"
  description = "Example Workload resource managed by Terraform"

  # Add additional configuration as needed
}
