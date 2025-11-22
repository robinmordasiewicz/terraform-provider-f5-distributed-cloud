resource "f5xc_subnet" "example" {
  name        = "my-subnet"
  namespace   = "system"
  description = "Example subnet"
  cidr        = "10.0.1.0/24"
}
