# Example configuration for f5xc_nfv_service

resource "f5xc_nfv_service" "example" {
  name        = "example-nfv_service"
  namespace   = "system"
  description = "Example NfvService resource managed by Terraform"

  # Add additional configuration as needed
}
