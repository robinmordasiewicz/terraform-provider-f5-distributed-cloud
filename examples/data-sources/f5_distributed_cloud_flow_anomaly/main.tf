# Example configuration for f5_distributed_cloud_flow_anomaly data source

data "f5_distributed_cloud_flow_anomaly" "example" {
  name      = "existing-flow_anomaly"
  namespace = "system"
}

output "flow_anomaly_id" {
  value = data.f5_distributed_cloud_flow_anomaly.example.id
}
