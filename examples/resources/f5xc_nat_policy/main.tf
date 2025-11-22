resource "f5xc_nat_policy" "example" {
  name        = "my-nat-policy"
  namespace   = "system"
  description = "Network address translation policy"
}
