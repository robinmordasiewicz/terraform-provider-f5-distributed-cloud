# Example: Cluster
resource "f5xc_cluster" "example" {
  name        = "my-cluster"
  namespace   = "my-namespace"
  description = "Example cluster for container orchestration"
}
