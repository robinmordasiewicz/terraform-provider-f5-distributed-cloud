# Example: Fleet
resource "f5xc_fleet" "example" {
  name        = "my-fleet"
  namespace   = "system"
  description = "Example fleet for edge site management"
}
