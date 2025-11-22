# Example: Segment Connection
resource "f5xc_segment_connection" "example" {
  name        = "my-segment-connection"
  namespace   = "my-namespace"
  description = "Example segment connection linking two segments"
}
