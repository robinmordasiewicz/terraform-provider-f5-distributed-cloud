# Example: Public IP
resource "f5distributedcloud_public_ip" "example" {
  name        = "my-public-ip"
  namespace   = "my-namespace"
  description = "Example public IP for external access"
}
