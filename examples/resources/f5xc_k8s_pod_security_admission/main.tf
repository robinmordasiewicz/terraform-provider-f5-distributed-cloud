# Example configuration for f5xc_k8s_pod_security_admission

resource "f5xc_k8s_pod_security_admission" "example" {
  name        = "example-k8s_pod_security_admission"
  namespace   = "system"
  description = "Example K8SPodSecurityAdmission resource managed by Terraform"

  # Add additional configuration as needed
}
