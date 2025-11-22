resource "f5xc_fast_acl_rule" "example" {
  name        = "my-fast-acl-rule"
  namespace   = "system"
  description = "Individual access control entry"
}
