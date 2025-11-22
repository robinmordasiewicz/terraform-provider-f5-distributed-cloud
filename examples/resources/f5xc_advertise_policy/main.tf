resource "f5xc_advertise_policy" "example" {
  name        = "my-advertise-policy"
  namespace   = "system"
  description = "Advertise policy for route advertisement"
}
