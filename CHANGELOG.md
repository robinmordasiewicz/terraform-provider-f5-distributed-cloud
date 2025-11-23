# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial provider implementation
- Authentication support for API tokens and P12 certificates
- HTTP client with retry logic and error handling

### Resources
- `f5_distributed_cloud_namespace` - Manage F5 XC namespaces
- `f5_distributed_cloud_origin_pool` - Manage origin server pools
- `f5_distributed_cloud_http_loadbalancer` - Manage HTTP load balancers
- `f5_distributed_cloud_cloud_site` - Manage cloud sites (AWS VPC, Azure VNET, GCP VPC)
- `f5_distributed_cloud_app_firewall` - Manage application firewall (WAF) policies

### Data Sources
- `f5_distributed_cloud_namespace` - Read existing namespaces
- `f5_distributed_cloud_origin_pool` - Read existing origin pools

### Documentation
- Getting started guide
- Authentication guide
- Migration guide from official provider
- Upgrade guide
- Resource and data source documentation

## [0.1.0] - TBD

### Added
- Initial release of the F5 Distributed Cloud Terraform Provider
- Full CRUD support for core resources
- Comprehensive documentation and examples
- Acceptance tests for all resources

### Security
- Sensitive fields (api_token, p12_password) are properly marked
- P12 certificate password handling with secure memory practices

### Dependencies
- Terraform Plugin Framework v1.x
- Go 1.21+
- Terraform 1.0+
