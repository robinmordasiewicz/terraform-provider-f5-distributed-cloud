# Example: Segment
resource "f5xc_segment" "example" {
  name        = "my-segment"
  namespace   = "my-namespace"
  description = "Example network segment for isolation"
}
