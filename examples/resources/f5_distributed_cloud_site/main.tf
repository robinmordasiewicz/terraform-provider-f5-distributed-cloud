# Example: Site
resource "f5_distributed_cloud_site" "example" {
  name        = "my-site"
  namespace   = "system"
  description = "Example site representing an edge location"
}
