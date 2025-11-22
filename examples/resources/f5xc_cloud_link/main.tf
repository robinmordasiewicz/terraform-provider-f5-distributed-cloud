# Example configuration for f5xc_cloud_link

resource "f5xc_cloud_link" "example" {
  name        = "example-cloud_link"
  namespace   = "system"
  description = "Example CloudLink resource managed by Terraform"

  # Add additional configuration as needed
}
