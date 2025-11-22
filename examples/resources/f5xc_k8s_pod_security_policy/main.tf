# Example configuration for f5xc_k8s_pod_security_policy

resource "f5xc_k8s_pod_security_policy" "example" {
  name        = "example-k8s_pod_security_policy"
  namespace   = "system"
  description = "Example K8SPodSecurityPolicy resource managed by Terraform"

  # Add additional configuration as needed
}
