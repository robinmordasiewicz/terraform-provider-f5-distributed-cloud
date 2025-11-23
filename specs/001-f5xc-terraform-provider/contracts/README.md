# API Contracts

This directory contains API contracts for the F5 Distributed Cloud Terraform Provider.

## Contract Files

| Contract | Resource Type | Description |
|----------|---------------|-------------|
| [provider.md](./provider.md) | Provider | Provider configuration and authentication |
| [namespace.md](./namespace.md) | `f5_distributed_cloud_namespace` | Namespace resource operations |
| [http_loadbalancer.md](./http_loadbalancer.md) | `f5_distributed_cloud_http_loadbalancer` | HTTP/HTTPS load balancer |
| [origin_pool.md](./origin_pool.md) | `f5_distributed_cloud_origin_pool` | Origin pool backend definitions |

## Contract Structure

Each contract document follows this structure:

1. **Resource Schema**: HCL schema definition with attribute types
2. **CRUD Operations**: API endpoints, pre/postconditions, error handling
3. **Validation Rules**: Attribute constraints and format requirements
4. **Import**: Import ID format and examples

## Common Patterns

### API URL Pattern
```
/api/config/namespaces/{namespace}/{resource_type}s/{name}
```

### Standard CRUD Mapping
| Operation | HTTP Method | Endpoint |
|-----------|-------------|----------|
| Create | POST | `.../{resource_type}s` |
| Read | GET | `.../{resource_type}s/{name}` |
| Update | PUT | `.../{resource_type}s/{name}` |
| Delete | DELETE | `.../{resource_type}s/{name}` |
| List | GET | `.../{resource_type}s` |

### Import ID Format
```
{namespace}/{name}
```

For system-level resources (e.g., cloud_credentials):
```
system/{name}
```
