# Example: Site
resource "f5distributedcloud_site" "example" {
  name        = "my-site"
  namespace   = "system"
  description = "Example site representing an edge location"
}
