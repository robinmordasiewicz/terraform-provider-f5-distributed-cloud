# Example configuration for f5xc_flow_anomaly data source

data "f5xc_flow_anomaly" "example" {
  name      = "existing-flow_anomaly"
  namespace = "system"
}

output "flow_anomaly_id" {
  value = data.f5xc_flow_anomaly.example.id
}
