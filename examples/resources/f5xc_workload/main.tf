# Example: Workload
resource "f5xc_workload" "example" {
  name        = "my-workload"
  namespace   = "my-namespace"
  description = "Example workload definition for application deployment"
}
