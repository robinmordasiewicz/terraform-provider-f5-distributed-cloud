resource "f5xc_policer" "example" {
  name        = "my-policer"
  namespace   = "system"
  description = "Traffic rate limiting policer"
}
