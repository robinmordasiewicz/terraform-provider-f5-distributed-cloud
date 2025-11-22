resource "f5xc_virtual_host" "example" {
  name        = "my-vhost"
  namespace   = "system"
  description = "Example virtual host"
  domains     = ["app.example.com", "www.example.com"]
}
