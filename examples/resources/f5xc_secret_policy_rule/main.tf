# Example: Secret Policy Rule
resource "f5xc_secret_policy_rule" "example" {
  name        = "my-secret-policy-rule"
  namespace   = "my-namespace"
  description = "Example secret policy rule"
}
