resource "f5xc_fast_acl" "example" {
  name        = "my-fast-acl"
  namespace   = "system"
  description = "High-performance access control list"
}
