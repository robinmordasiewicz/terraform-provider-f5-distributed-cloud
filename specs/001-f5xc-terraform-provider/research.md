# Technical Research: F5 Distributed Cloud Terraform Provider

**Branch**: `001-f5xc-terraform-provider` | **Date**: 2025-11-21 | **Spec**: [spec.md](./spec.md)

## Executive Summary

This document captures technical research for building an open source Terraform provider for F5 Distributed Cloud (F5 XC) services. The provider will leverage 269 OpenAPI specifications in `docs/specifications/api/` and follow HashiCorp's Terraform Plugin Framework patterns.

## 1. Framework Selection: Terraform Plugin Framework

**Recommendation**: Use `terraform-plugin-framework` (not legacy SDKv2)

### Rationale
- HashiCorp stopped most feature development on SDKv2 for Terraform 1.x+
- Framework provides superior data handling with separate config/plan/state access
- Better type safety with explicit null/unknown value handling
- Improved validation capabilities through multiple integration points
- Support for resource private state management (e.g., API ETag values)
- Modern request-response patterns familiar to server-based applications

### Key Framework Advantages
| Feature | SDKv2 | Framework |
|---------|-------|-----------|
| Data Access | Single `schema.ResourceData` | Separate config/plan/state |
| Null Handling | Zero-value ambiguity | Explicit `IsNull()`/`IsUnknown()` |
| Type Assertions | Required with `interface{}` | Type-safe data models |
| Validation | Limited field-level | Multi-point extensible validation |
| Plan Modification | Automatic (often unexpected) | Explicit plan modifiers |

## 2. API Architecture Analysis

### Base URL Pattern
```
https://<tenant>.console.ves.volterra.io/api
```

### Service Prefixes
- `/api/config/` - Configuration CRUD operations
- `/api/web/` - Query and reporting operations
- `/api/cdn/` - CDN-specific operations
- `/api/ml/` - Machine learning/analytics endpoints

### Authentication Methods
1. **Certificate Authentication (mTLS)**
   - P12 certificate extracted to cert/key PEM files
   - Configured via `cert_file` and `key_file` provider attributes

2. **API Token Authentication**
   - Token in Authorization header: `APIToken <token>`
   - Configured via `api_token` provider attribute

### Standard CRUD Pattern (from OpenAPI analysis)
```
POST   /api/config/namespaces/{namespace}/{resource_type}s          # Create
GET    /api/config/namespaces/{namespace}/{resource_type}s          # List
GET    /api/config/namespaces/{namespace}/{resource_type}s/{name}   # Read
PUT    /api/config/namespaces/{namespace}/{resource_type}s/{name}   # Replace
DELETE /api/config/namespaces/{namespace}/{resource_type}s/{name}   # Delete
```

### Response Codes
| Code | Meaning | Action |
|------|---------|--------|
| 200 | Success | Process response |
| 401 | Not authorized | Check credentials |
| 403 | No permission | Check RBAC |
| 404 | Not found | Handle missing resource |
| 409 | Conflict | Handle concurrent modification |
| 429 | Rate limited | Implement backoff retry |
| 500 | Server error | Retry with backoff |
| 503 | Service unavailable | Retry with backoff |
| 504 | Timeout | Retry with backoff |

## 3. OpenAPI Specification Analysis

### Specification Count
- **Total**: 269 OpenAPI 3.0 specification files
- **Location**: `docs/specifications/api/`

### Key Resource Specifications Identified
| Resource | OpenAPI File | Priority |
|----------|-------------|----------|
| HTTP Load Balancer | `...http_loadbalancer.ves-swagger.json` | P1 |
| TCP Load Balancer | `...tcp_loadbalancer.ves-swagger.json` | P1 |
| Origin Pool | `...origin_pool.ves-swagger.json` | P2 |
| App Firewall | `...app_firewall.ves-swagger.json` | P2 |
| Namespace | `...namespace.ves-swagger.json` | P1 |
| Cloud Credentials | `...cloud_credentials.ves-swagger.json` | P2 |
| AWS VPC Site | `...aws_vpc_site.ves-swagger.json` | P2 |
| Azure VNet Site | `...azure_vnet_site.ves-swagger.json` | P2 |
| GCP VPC Site | `...gcp_vpc_site.ves-swagger.json` | P2 |
| DNS Load Balancer | `...dns_load_balancer.ves-swagger.json` | P1 |
| CDN Load Balancer | `...cdn_loadbalancer.ves-swagger.json` | P1 |
| Certificate | `...certificate.ves-swagger.json` | P2 |

### Schema Patterns Observed
- All resources use `metadata` object with `name`, `namespace`, `labels`, `annotations`
- Spec objects contain resource-specific configuration
- Status objects contain operational state
- OneOf fields indicated by `x-ves-oneof-field-*` attributes
- Validation rules in `x-ves-validation-rules` attributes

### Common Schema Types
```yaml
ObjectCreateMetaType:
  - name (required, DNS-1035 format)
  - namespace (DNS_LABEL format)
  - labels (map[string]string)
  - annotations (map[string]string)
  - description (max 1200 bytes)
  - disable (boolean)

SecretType:
  - blindfold_secret_info (F5XC managed)
  - clear_secret_info (user provided)
```

## 4. Proprietary Provider Reference Analysis

### Resource Naming Convention
Resources from `volterraedge/volterra` provider:
- `volterra_namespace`
- `volterra_http_loadbalancer`
- `volterra_tcp_loadbalancer`
- `volterra_origin_pool`
- `volterra_app_firewall`
- `volterra_aws_vpc_site`
- `volterra_azure_vnet_site`
- `volterra_gcp_vpc_site`
- `volterra_cloud_credentials`
- `volterra_healthcheck`
- `volterra_certificate`

### Migration Compatibility
- Resource naming should use `f5distributedcloud_` prefix (e.g., `f5distributedcloud_http_loadbalancer`)
- Schema structure should align closely with proprietary provider
- Import functionality enables state migration

## 5. Project Structure Recommendation

```
terraform-provider-f5distributedcloud/
├── main.go                          # Provider entrypoint
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── GNUmakefile                      # Build automation
│
├── internal/
│   ├── provider/
│   │   ├── provider.go              # Provider definition
│   │   └── provider_test.go         # Provider tests
│   │
│   ├── client/
│   │   ├── client.go                # HTTP client wrapper
│   │   ├── auth.go                  # Authentication handling
│   │   └── errors.go                # Error types
│   │
│   ├── resources/
│   │   ├── namespace/
│   │   │   ├── resource.go          # Resource implementation
│   │   │   ├── schema.go            # Schema definition
│   │   │   ├── model.go             # Data models
│   │   │   └── resource_test.go     # Tests
│   │   ├── http_loadbalancer/
│   │   ├── tcp_loadbalancer/
│   │   ├── origin_pool/
│   │   └── ...
│   │
│   ├── datasources/
│   │   ├── namespace/
│   │   └── ...
│   │
│   └── common/
│       ├── metadata.go              # Shared metadata handling
│       ├── validators.go            # Custom validators
│       └── planmodifiers.go         # Plan modifiers
│
├── templates/                       # Documentation templates
│   ├── index.md.tmpl
│   └── resources/
│
├── examples/
│   ├── provider/
│   └── resources/
│
├── docs/                            # Generated documentation
│
└── tests/
    ├── acceptance/                  # Acceptance tests
    └── unit/                        # Unit tests
```

## 6. Code Generation Strategy

### Hybrid Approach Recommended
1. **Generated Code** (`*_gen.go`):
   - Schema definitions from OpenAPI specs
   - Data model structs with `tfsdk` tags
   - Request/Response type mappings

2. **Manual Code**:
   - CRUD operation implementations
   - Custom validators and plan modifiers
   - Error handling and retry logic
   - Complex nested attribute handling

### OpenAPI to Terraform Mapping
| OpenAPI Type | Terraform Framework Type |
|--------------|-------------------------|
| `string` | `types.String` |
| `integer` | `types.Int64` |
| `boolean` | `types.Bool` |
| `number` | `types.Float64` |
| `array` | `types.List` / `types.Set` |
| `object` | `types.Object` / Nested Attributes |
| `oneOf` | Dynamic attribute with validators |

## 7. Testing Strategy

### Test Pyramid
1. **Unit Tests**: Schema validation, model conversion
2. **Integration Tests**: Client API interactions (mocked)
3. **Acceptance Tests**: Full Terraform lifecycle with real API

### Acceptance Test Pattern
```go
func TestAccResource_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Create and Read
            {
                Config: testAccResourceConfig_basic("test"),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("f5distributedcloud_namespace.test", "name", "test"),
                ),
            },
            // ImportState
            {
                ResourceName:      "f5distributedcloud_namespace.test",
                ImportState:       true,
                ImportStateVerify: true,
            },
            // Update
            // Delete (implicit)
        },
    })
}
```

## 8. Implementation Priorities

### Phase 1: Foundation (P1)
1. Provider configuration (auth, tenant URL)
2. Client HTTP wrapper with retry logic
3. Namespace resource (simplest CRUD)
4. HTTP Load Balancer resource (core use case)

### Phase 2: Core Resources (P1-P2)
5. Origin Pool resource
6. TCP Load Balancer resource
7. DNS Load Balancer resource
8. App Firewall resource

### Phase 3: Cloud Sites (P2)
9. Cloud Credentials resource
10. AWS VPC Site resource
11. Azure VNet Site resource
12. GCP VPC Site resource

### Phase 4: Advanced Features (P2-P3)
13. Data sources for all resources
14. Import support for all resources
15. CDN Load Balancer resource
16. Certificate management resources

## 9. Key Technical Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Framework | terraform-plugin-framework | Modern, recommended by HashiCorp |
| Language | Go 1.21+ | Terraform provider requirement |
| HTTP Client | net/http + retryablehttp | Standard + retry support |
| JSON Handling | encoding/json | Standard library, reliable |
| Testing | testify + terraform-plugin-testing | Standard Terraform provider testing |
| Documentation | terraform-plugin-docs | Auto-generated from schemas |

## 10. Risk Mitigations

| Risk | Mitigation |
|------|-----------|
| API version changes | Version-specific client paths, deprecation handling |
| Rate limiting | Exponential backoff, configurable retry settings |
| Large schemas | Code generation for complex nested structures |
| State drift | Comprehensive Read implementation, drift detection |
| Concurrent access | Optimistic locking with API ETag support where available |

## References

- [Terraform Plugin Framework Documentation](https://developer.hashicorp.com/terraform/plugin/framework)
- [OpenAPI Provider Spec Generator](https://developer.hashicorp.com/terraform/plugin/code-generation/openapi-generator)
- [F5 XC API Documentation](https://docs.cloud.f5.com/docs-v2/api)
- [Proprietary Volterra Provider](https://registry.terraform.io/providers/volterraedge/volterra/latest/docs)
