# Implementation Plan: F5 Distributed Cloud Terraform Provider

**Branch**: `001-f5xc-terraform-provider` | **Date**: 2025-11-21 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/001-f5xc-terraform-provider/spec.md`

## Summary

Open source Terraform provider for F5 Distributed Cloud (F5 XC) services, built using terraform-plugin-framework from 269 OpenAPI specifications. Provides feature parity with proprietary `volterraedge/volterra` provider while enabling community contributions and transparency.

## Technical Context

**Language/Version**: Go 1.21+
**Primary Dependencies**: terraform-plugin-framework, terraform-plugin-testing, retryablehttp
**Storage**: N/A (stateless provider, uses Terraform state)
**Testing**: go test, terraform-plugin-testing for acceptance tests
**Target Platform**: Cross-platform (darwin, linux, windows) - amd64, arm64
**Project Type**: Single (Terraform provider)
**Performance Goals**: Resource operations <30s for simple resources, <10min for cloud sites
**Constraints**: Must support F5 XC API rate limits, mTLS and token authentication
**Scale/Scope**: 269 OpenAPI specs, ~25 initial resources, unlimited configuration size

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. OpenAPI Specification Fidelity | ✅ Pass | Schemas derived from `docs/specifications/api/` |
| II. Test-First Development | ✅ Pass | Acceptance tests required before implementation |
| III. Idiomatic Terraform Patterns | ✅ Pass | Using terraform-plugin-framework |
| IV. Documentation Excellence | ✅ Pass | terraform-plugin-docs, examples required |
| V. Semantic Versioning | ✅ Pass | Standard semver for releases |

## Project Structure

### Documentation (this feature)

```text
specs/001-f5xc-terraform-provider/
├── spec.md                      # Feature specification
├── plan.md                      # This file
├── research.md                  # Technical research findings
├── data-model.md                # Entity schemas and relationships
├── quickstart.md                # Developer quickstart guide
├── contracts/                   # API contracts
│   ├── README.md
│   ├── provider.md
│   ├── namespace.md
│   ├── http_loadbalancer.md
│   └── origin_pool.md
└── checklists/
    └── requirements.md          # Specification quality checklist
```

### Source Code (repository root)

```text
terraform-provider-f5-distributed-cloud/
├── main.go                          # Provider entrypoint
├── go.mod                           # Go module definition
├── go.sum                           # Dependency checksums
├── GNUmakefile                      # Build automation
├── .goreleaser.yml                  # Release configuration
│
├── internal/
│   ├── provider/
│   │   ├── provider.go              # Provider definition
│   │   └── provider_test.go         # Provider tests
│   │
│   ├── client/
│   │   ├── client.go                # HTTP client wrapper
│   │   ├── auth.go                  # Authentication handling
│   │   └── errors.go                # Error types and handling
│   │
│   ├── resources/
│   │   ├── namespace/               # f5xc_namespace
│   │   ├── http_loadbalancer/       # f5xc_http_loadbalancer
│   │   ├── tcp_loadbalancer/        # f5xc_tcp_loadbalancer
│   │   ├── origin_pool/             # f5xc_origin_pool
│   │   ├── app_firewall/            # f5xc_app_firewall
│   │   ├── cloud_credentials/       # f5xc_cloud_credentials
│   │   ├── aws_vpc_site/            # f5xc_aws_vpc_site
│   │   ├── azure_vnet_site/         # f5xc_azure_vnet_site
│   │   ├── gcp_vpc_site/            # f5xc_gcp_vpc_site
│   │   ├── dns_load_balancer/       # f5xc_dns_load_balancer
│   │   └── cdn_loadbalancer/        # f5xc_cdn_loadbalancer
│   │
│   ├── datasources/
│   │   ├── namespace/
│   │   └── ... (mirrors resources)
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
│   │   └── provider.tf
│   └── resources/
│       ├── f5xc_namespace/
│       ├── f5xc_http_loadbalancer/
│       └── ...
│
├── docs/                            # Generated documentation
│   └── specifications/api/          # OpenAPI specs (existing)
│
└── tests/
    ├── acceptance/                  # Acceptance tests
    └── unit/                        # Unit tests
```

**Structure Decision**: Single project structure following standard Terraform provider layout. Resources organized by resource type with co-located tests.

## Implementation Phases

### Phase 1: Foundation
1. Provider scaffolding with terraform-plugin-framework
2. HTTP client with mTLS and token authentication
3. Error handling and retry logic
4. `f5xc_namespace` resource (simplest CRUD)

### Phase 2: Core Load Balancers
5. `f5xc_origin_pool` resource
6. `f5xc_http_loadbalancer` resource
7. `f5xc_tcp_loadbalancer` resource
8. `f5xc_dns_load_balancer` resource

### Phase 3: Security & CDN
9. `f5xc_app_firewall` resource
10. `f5xc_cdn_loadbalancer` resource

### Phase 4: Cloud Sites
11. `f5xc_cloud_credentials` resource
12. `f5xc_aws_vpc_site` resource
13. `f5xc_azure_vnet_site` resource
14. `f5xc_gcp_vpc_site` resource

### Phase 5: Data Sources & Polish
15. Data sources for all resources
16. Import support verification
17. Documentation generation
18. Release preparation

## Complexity Tracking

> No constitution violations requiring justification.

| Decision | Rationale | Alternative Considered |
|----------|-----------|------------------------|
| terraform-plugin-framework | HashiCorp recommended, better type safety | SDKv2 - legacy, feature-frozen |
| Per-resource packages | Scalability for 25+ resources | Single package - not maintainable |
| Hybrid code generation | Complex nested schemas need manual handling | Full generation - insufficient control |

## Related Artifacts

- [Technical Research](./research.md) - Framework selection, API analysis
- [Data Model](./data-model.md) - Entity schemas, relationships
- [Quickstart Guide](./quickstart.md) - Developer onboarding
- [API Contracts](./contracts/) - Resource operation contracts
- [Feature Specification](./spec.md) - Requirements and user stories
