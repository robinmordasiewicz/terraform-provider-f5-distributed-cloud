# Example: Service Policy Rule
resource "f5xc_service_policy_rule" "example" {
  name        = "my-service-policy-rule"
  namespace   = "my-namespace"
  description = "Example service policy rule"
}
