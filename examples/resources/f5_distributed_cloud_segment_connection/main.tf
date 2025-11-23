# Example: Segment Connection
resource "f5_distributed_cloud_segment_connection" "example" {
  name        = "my-segment-connection"
  namespace   = "my-namespace"
  description = "Example segment connection linking two segments"
}
