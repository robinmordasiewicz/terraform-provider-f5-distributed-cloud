# Example configuration for f5_distributed_cloud_k8s_pod_security_admission data source

data "f5_distributed_cloud_k8s_pod_security_admission" "example" {
  name      = "existing-k8s_pod_security_admission"
  namespace = "system"
}

output "k8s_pod_security_admission_id" {
  value = data.f5_distributed_cloud_k8s_pod_security_admission.example.id
}
