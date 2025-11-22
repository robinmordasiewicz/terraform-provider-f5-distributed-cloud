# Example configuration for f5xc_k8s_pod_security_policy data source

data "f5xc_k8s_pod_security_policy" "example" {
  name      = "existing-k8s_pod_security_policy"
  namespace = "system"
}

output "k8s_pod_security_policy_id" {
  value = data.f5xc_k8s_pod_security_policy.example.id
}
