# Example: Securemesh Site
resource "f5xc_securemesh_site" "example" {
  name        = "my-securemesh-site"
  namespace   = "system"
  description = "Example securemesh site for secure connectivity"
}
