# Example: Registration
resource "f5xc_registration" "example" {
  name        = "my-registration"
  namespace   = "my-namespace"
  description = "Example site registration"
}
