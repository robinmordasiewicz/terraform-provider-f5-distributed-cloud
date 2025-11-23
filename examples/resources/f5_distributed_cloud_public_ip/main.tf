# Example: Public IP
resource "f5_distributed_cloud_public_ip" "example" {
  name        = "my-public-ip"
  namespace   = "my-namespace"
  description = "Example public IP for external access"
}
