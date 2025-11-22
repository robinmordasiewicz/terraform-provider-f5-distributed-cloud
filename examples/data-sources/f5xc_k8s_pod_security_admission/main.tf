# Example configuration for f5xc_k8s_pod_security_admission data source

data "f5xc_k8s_pod_security_admission" "example" {
  name      = "existing-k8s_pod_security_admission"
  namespace = "system"
}

output "k8s_pod_security_admission_id" {
  value = data.f5xc_k8s_pod_security_admission.example.id
}
