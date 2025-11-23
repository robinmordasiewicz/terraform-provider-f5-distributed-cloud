# Example: Service Policy Set
resource "f5distributedcloud_service_policy_set" "example" {
  name        = "my-service-policy-set"
  namespace   = "my-namespace"
  description = "Example service policy set grouping policies"
}
