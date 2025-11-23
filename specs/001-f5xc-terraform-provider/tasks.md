# Tasks: F5 Distributed Cloud Terraform Provider

**Input**: Design documents from `/specs/001-f5xc-terraform-provider/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md, contracts/

**Tests**: Tests are generated as acceptance tests per Terraform provider convention (test-first per constitution).

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

Based on plan.md structure:
- Provider code: `internal/provider/`
- Client code: `internal/client/`
- Resources: `internal/resources/{resource_name}/`
- Data sources: `internal/datasources/{resource_name}/`
- Common utilities: `internal/common/`
- Tests: `tests/acceptance/`, co-located `*_test.go` files
- Examples: `examples/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and Go module setup

- [ ] T001 Create project directory structure per plan.md in repository root
- [ ] T002 Initialize Go module with `go mod init` and create go.mod
- [ ] T003 [P] Add terraform-plugin-framework dependencies to go.mod
- [ ] T004 [P] Add terraform-plugin-testing dependencies to go.mod
- [ ] T005 [P] Add retryablehttp and other HTTP client dependencies to go.mod
- [ ] T006 [P] Create GNUmakefile with build, test, and install targets
- [ ] T007 [P] Create .goreleaser.yml for cross-platform release builds
- [ ] T008 [P] Configure golangci-lint with .golangci.yml
- [ ] T009 Create main.go provider entrypoint at repository root

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core infrastructure that MUST be complete before ANY user story can be implemented

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [ ] T010 Define error types and handling in internal/client/errors.go
- [ ] T011 Implement HTTP client wrapper in internal/client/client.go
- [ ] T012 [P] Implement certificate authentication in internal/client/auth.go
- [ ] T013 [P] Implement token authentication in internal/client/auth.go
- [ ] T014 Implement retry logic with exponential backoff in internal/client/client.go
- [ ] T015 Create shared metadata schema types in internal/common/metadata.go
- [ ] T016 [P] Create custom validators in internal/common/validators.go
- [ ] T017 [P] Create plan modifiers in internal/common/planmodifiers.go
- [ ] T018 Implement base provider scaffold in internal/provider/provider.go

**Checkpoint**: Foundation ready - user story implementation can now begin

---

## Phase 3: User Story 1 - Provider Authentication Configuration (Priority: P1) 🎯 MVP

**Goal**: Infrastructure engineers can configure the provider with F5 XC credentials (certificate or token) to authenticate API requests.

**Independent Test**: Initialize provider with credentials and verify successful connection to F5 XC API endpoint.

### Acceptance Tests for User Story 1

- [ ] T019 [P] [US1] Write acceptance test for certificate auth in internal/provider/provider_test.go
- [ ] T020 [P] [US1] Write acceptance test for token auth in internal/provider/provider_test.go
- [ ] T021 [P] [US1] Write acceptance test for invalid credentials error handling in internal/provider/provider_test.go
- [ ] T022 [P] [US1] Write acceptance test for missing credentials error in internal/provider/provider_test.go

### Implementation for User Story 1

- [ ] T023 [US1] Define provider schema with api_url, cert_file, key_file, api_token in internal/provider/provider.go
- [ ] T024 [US1] Implement provider Configure method with auth selection logic in internal/provider/provider.go
- [ ] T025 [US1] Add environment variable support (F5XC_API_URL, F5XC_API_TOKEN, etc.) in internal/provider/provider.go
- [ ] T026 [US1] Implement provider validation (require api_url, require either cert or token) in internal/provider/provider.go
- [ ] T027 [US1] Add clear error messages for authentication failures in internal/provider/provider.go
- [ ] T028 [P] [US1] Create provider example in examples/provider/provider.tf

**Checkpoint**: Provider authentication fully functional and independently testable

---

## Phase 4: User Story 2 - Namespace Resource Management (Priority: P1)

**Goal**: Infrastructure engineers can create, read, update, and delete namespaces as the fundamental organizational unit.

**Independent Test**: Create namespace, verify in F5 XC console, modify description, delete it.

### Acceptance Tests for User Story 2

- [ ] T029 [P] [US2] Write acceptance test for namespace create in internal/resources/namespace/resource_test.go
- [ ] T030 [P] [US2] Write acceptance test for namespace update in internal/resources/namespace/resource_test.go
- [ ] T031 [P] [US2] Write acceptance test for namespace delete in internal/resources/namespace/resource_test.go
- [ ] T032 [P] [US2] Write acceptance test for namespace import in internal/resources/namespace/resource_test.go

### Implementation for User Story 2

- [ ] T033 [P] [US2] Define namespace data model in internal/resources/namespace/model.go
- [ ] T034 [P] [US2] Define namespace schema in internal/resources/namespace/schema.go
- [ ] T035 [US2] Implement namespace Create method in internal/resources/namespace/resource.go
- [ ] T036 [US2] Implement namespace Read method in internal/resources/namespace/resource.go
- [ ] T037 [US2] Implement namespace Update method in internal/resources/namespace/resource.go
- [ ] T038 [US2] Implement namespace Delete method in internal/resources/namespace/resource.go
- [ ] T039 [US2] Implement namespace ImportState method in internal/resources/namespace/resource.go
- [ ] T040 [US2] Register namespace resource with provider in internal/provider/provider.go
- [ ] T041 [P] [US2] Create namespace example in examples/resources/f5_distributed_cloud_namespace/resource.tf

**Checkpoint**: Namespace resource fully functional and independently testable

---

## Phase 5: User Story 3 - HTTP Load Balancer Resource Management (Priority: P1)

**Goal**: Infrastructure engineers can create and manage HTTP load balancers to distribute traffic to application backends.

**Independent Test**: Create HTTP load balancer with domain, verify in F5 XC console, confirm traffic routing.

**Dependencies**: Requires namespace (US2), benefits from origin_pool (US4) but can use external pools

### Acceptance Tests for User Story 3

- [ ] T042 [P] [US3] Write acceptance test for http_loadbalancer create in internal/resources/http_loadbalancer/resource_test.go
- [ ] T043 [P] [US3] Write acceptance test for http_loadbalancer update in internal/resources/http_loadbalancer/resource_test.go
- [ ] T044 [P] [US3] Write acceptance test for http_loadbalancer delete in internal/resources/http_loadbalancer/resource_test.go
- [ ] T045 [P] [US3] Write acceptance test for http_loadbalancer import in internal/resources/http_loadbalancer/resource_test.go

### Implementation for User Story 3

- [ ] T046 [P] [US3] Define http_loadbalancer data model in internal/resources/http_loadbalancer/model.go
- [ ] T047 [P] [US3] Define http_loadbalancer schema in internal/resources/http_loadbalancer/schema.go
- [ ] T048 [US3] Implement http_loadbalancer Create method in internal/resources/http_loadbalancer/resource.go
- [ ] T049 [US3] Implement http_loadbalancer Read method in internal/resources/http_loadbalancer/resource.go
- [ ] T050 [US3] Implement http_loadbalancer Update method in internal/resources/http_loadbalancer/resource.go
- [ ] T051 [US3] Implement http_loadbalancer Delete method in internal/resources/http_loadbalancer/resource.go
- [ ] T052 [US3] Implement http_loadbalancer ImportState method in internal/resources/http_loadbalancer/resource.go
- [ ] T053 [US3] Register http_loadbalancer resource with provider in internal/provider/provider.go
- [ ] T054 [P] [US3] Create http_loadbalancer example in examples/resources/f5_distributed_cloud_http_loadbalancer/resource.tf

**Checkpoint**: HTTP Load Balancer resource fully functional and independently testable

---

## Phase 6: User Story 4 - Origin Pool Resource Management (Priority: P2)

**Goal**: Infrastructure engineers can define origin pools that specify backend servers for load balancers.

**Independent Test**: Create origin pool with backend endpoints and verify configuration in F5 XC console.

### Acceptance Tests for User Story 4

- [ ] T055 [P] [US4] Write acceptance test for origin_pool create in internal/resources/origin_pool/resource_test.go
- [ ] T056 [P] [US4] Write acceptance test for origin_pool update in internal/resources/origin_pool/resource_test.go
- [ ] T057 [P] [US4] Write acceptance test for origin_pool delete in internal/resources/origin_pool/resource_test.go
- [ ] T058 [P] [US4] Write acceptance test for origin_pool import in internal/resources/origin_pool/resource_test.go

### Implementation for User Story 4

- [ ] T059 [P] [US4] Define origin_pool data model in internal/resources/origin_pool/model.go
- [ ] T060 [P] [US4] Define origin_pool schema in internal/resources/origin_pool/schema.go
- [ ] T061 [US4] Implement origin_pool Create method in internal/resources/origin_pool/resource.go
- [ ] T062 [US4] Implement origin_pool Read method in internal/resources/origin_pool/resource.go
- [ ] T063 [US4] Implement origin_pool Update method in internal/resources/origin_pool/resource.go
- [ ] T064 [US4] Implement origin_pool Delete method in internal/resources/origin_pool/resource.go
- [ ] T065 [US4] Implement origin_pool ImportState method in internal/resources/origin_pool/resource.go
- [ ] T066 [US4] Register origin_pool resource with provider in internal/provider/provider.go
- [ ] T067 [P] [US4] Create origin_pool example in examples/resources/f5_distributed_cloud_origin_pool/resource.tf

**Checkpoint**: Origin Pool resource fully functional and independently testable

---

## Phase 7: User Story 5 - Cloud Site Resource Management (Priority: P2)

**Goal**: Infrastructure engineers can deploy and manage F5 XC sites in cloud environments (AWS, Azure, GCP).

**Independent Test**: Create AWS VPC site configuration and verify site registration in F5 XC console.

### Acceptance Tests for User Story 5

- [ ] T068 [P] [US5] Write acceptance test for cloud_credentials create in internal/resources/cloud_credentials/resource_test.go
- [ ] T069 [P] [US5] Write acceptance test for aws_vpc_site create in internal/resources/aws_vpc_site/resource_test.go
- [ ] T070 [P] [US5] Write acceptance test for azure_vnet_site create in internal/resources/azure_vnet_site/resource_test.go
- [ ] T071 [P] [US5] Write acceptance test for gcp_vpc_site create in internal/resources/gcp_vpc_site/resource_test.go

### Implementation for User Story 5

- [ ] T072 [P] [US5] Define cloud_credentials data model in internal/resources/cloud_credentials/model.go
- [ ] T073 [P] [US5] Define cloud_credentials schema in internal/resources/cloud_credentials/schema.go
- [ ] T074 [US5] Implement cloud_credentials CRUD methods in internal/resources/cloud_credentials/resource.go
- [ ] T075 [P] [US5] Define aws_vpc_site data model in internal/resources/aws_vpc_site/model.go
- [ ] T076 [P] [US5] Define aws_vpc_site schema in internal/resources/aws_vpc_site/schema.go
- [ ] T077 [US5] Implement aws_vpc_site CRUD methods in internal/resources/aws_vpc_site/resource.go
- [ ] T078 [P] [US5] Define azure_vnet_site data model in internal/resources/azure_vnet_site/model.go
- [ ] T079 [P] [US5] Define azure_vnet_site schema in internal/resources/azure_vnet_site/schema.go
- [ ] T080 [US5] Implement azure_vnet_site CRUD methods in internal/resources/azure_vnet_site/resource.go
- [ ] T081 [P] [US5] Define gcp_vpc_site data model in internal/resources/gcp_vpc_site/model.go
- [ ] T082 [P] [US5] Define gcp_vpc_site schema in internal/resources/gcp_vpc_site/schema.go
- [ ] T083 [US5] Implement gcp_vpc_site CRUD methods in internal/resources/gcp_vpc_site/resource.go
- [ ] T084 [US5] Register all cloud site resources with provider in internal/provider/provider.go
- [ ] T085 [P] [US5] Create cloud_credentials example in examples/resources/f5_distributed_cloud_cloud_credentials/resource.tf
- [ ] T086 [P] [US5] Create aws_vpc_site example in examples/resources/f5_distributed_cloud_aws_vpc_site/resource.tf

**Checkpoint**: Cloud Site resources fully functional and independently testable

---

## Phase 8: User Story 6 - Application Firewall Policy Management (Priority: P2)

**Goal**: Security engineers can define and apply WAF policies to protect applications.

**Independent Test**: Create WAF policy with rules and attach to an HTTP load balancer.

### Acceptance Tests for User Story 6

- [ ] T087 [P] [US6] Write acceptance test for app_firewall create in internal/resources/app_firewall/resource_test.go
- [ ] T088 [P] [US6] Write acceptance test for app_firewall update in internal/resources/app_firewall/resource_test.go
- [ ] T089 [P] [US6] Write acceptance test for app_firewall with exclusions in internal/resources/app_firewall/resource_test.go

### Implementation for User Story 6

- [ ] T090 [P] [US6] Define app_firewall data model in internal/resources/app_firewall/model.go
- [ ] T091 [P] [US6] Define app_firewall schema in internal/resources/app_firewall/schema.go
- [ ] T092 [US6] Implement app_firewall Create method in internal/resources/app_firewall/resource.go
- [ ] T093 [US6] Implement app_firewall Read method in internal/resources/app_firewall/resource.go
- [ ] T094 [US6] Implement app_firewall Update method in internal/resources/app_firewall/resource.go
- [ ] T095 [US6] Implement app_firewall Delete method in internal/resources/app_firewall/resource.go
- [ ] T096 [US6] Implement app_firewall ImportState method in internal/resources/app_firewall/resource.go
- [ ] T097 [US6] Register app_firewall resource with provider in internal/provider/provider.go
- [ ] T098 [P] [US6] Create app_firewall example in examples/resources/f5_distributed_cloud_app_firewall/resource.tf

**Checkpoint**: App Firewall resource fully functional and independently testable

---

## Phase 9: User Story 7 - Migration from Proprietary Provider (Priority: P2)

**Goal**: Organizations can migrate from proprietary volterra provider with minimal disruption.

**Independent Test**: Take existing volterra_* configuration and update to use new provider with state import.

### Implementation for User Story 7

- [ ] T099 [US7] Document migration process in docs/guides/migration.md
- [ ] T100 [US7] Ensure all resource schemas are compatible with volterra provider in respective resource files
- [ ] T101 [US7] Validate import functionality works for all implemented resources
- [ ] T102 [P] [US7] Create migration example in examples/guides/migration/

**Checkpoint**: Migration documentation and validation complete

---

## Phase 10: User Story 8 - Data Source Queries (Priority: P3)

**Goal**: Infrastructure engineers can query existing F5 XC resources for use in Terraform configurations.

**Independent Test**: Query existing namespace and use attributes in other resource definitions.

### Acceptance Tests for User Story 8

- [ ] T103 [P] [US8] Write acceptance test for namespace data source in internal/datasources/namespace/datasource_test.go
- [ ] T104 [P] [US8] Write acceptance test for http_loadbalancer data source in internal/datasources/http_loadbalancer/datasource_test.go

### Implementation for User Story 8

- [ ] T105 [P] [US8] Implement namespace data source in internal/datasources/namespace/datasource.go
- [ ] T106 [P] [US8] Implement http_loadbalancer data source in internal/datasources/http_loadbalancer/datasource.go
- [ ] T107 [P] [US8] Implement origin_pool data source in internal/datasources/origin_pool/datasource.go
- [ ] T108 [P] [US8] Implement app_firewall data source in internal/datasources/app_firewall/datasource.go
- [ ] T109 [US8] Register all data sources with provider in internal/provider/provider.go
- [ ] T110 [P] [US8] Create data source examples in examples/data-sources/

**Checkpoint**: Data sources fully functional and independently testable

---

## Phase 11: Polish & Cross-Cutting Concerns

**Purpose**: Documentation, quality improvements, and release preparation

- [ ] T111 [P] Create provider documentation template in templates/index.md.tmpl
- [ ] T112 [P] Create resource documentation templates in templates/resources/
- [ ] T113 Generate documentation using terraform-plugin-docs
- [ ] T114 [P] Add unit tests for common utilities in internal/common/*_test.go
- [ ] T115 [P] Add unit tests for client error handling in internal/client/*_test.go
- [ ] T116 Run golangci-lint and fix all issues
- [ ] T117 Run full acceptance test suite
- [ ] T118 Validate all examples with terraform validate
- [ ] T119 Run quickstart.md validation scenarios
- [ ] T120 [P] Create CHANGELOG.md with initial release notes
- [ ] T121 [P] Create CONTRIBUTING.md for open source contributions
- [ ] T122 Update README.md with installation and usage instructions

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup completion - BLOCKS all user stories
- **User Stories (Phase 3-10)**: All depend on Foundational phase completion
  - US1 (Auth) should complete first as other tests need working auth
  - US2 (Namespace) before US3-US6 (resources require namespaces)
  - US3-US6 can proceed in parallel after US2
  - US7 (Migration) after primary resources implemented
  - US8 (Data Sources) after resources are implemented
- **Polish (Phase 11)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Foundational - No dependencies on other stories
- **User Story 2 (P1)**: Depends on US1 (needs working auth for tests)
- **User Story 3 (P1)**: Depends on US2 (needs namespace), can use external origin pools
- **User Story 4 (P2)**: Depends on US2 (needs namespace)
- **User Story 5 (P2)**: Depends on US2 (needs namespace for some resources)
- **User Story 6 (P2)**: Depends on US2, benefits from US3 for integration testing
- **User Story 7 (P2)**: Depends on US2-US6 being implemented
- **User Story 8 (P3)**: Depends on corresponding resources being implemented

### Within Each User Story

- Acceptance tests MUST be written and FAIL before implementation (per constitution)
- Models and schemas before CRUD implementation
- CRUD implementation before ImportState
- Resource registration after implementation complete
- Examples after resource is functional

### Parallel Opportunities

**Phase 1 (Setup)**: T003, T004, T005, T006, T007, T008 can run in parallel
**Phase 2 (Foundation)**: T012/T013 (auth methods), T016/T017 (validators/modifiers) can run in parallel
**Each User Story**: Model and schema tasks marked [P] can run in parallel
**Data Sources (US8)**: All data source implementations can run in parallel

---

## Parallel Example: User Story 2 (Namespace)

```bash
# Launch all tests together (they will fail initially):
Task: T029 "Write acceptance test for namespace create"
Task: T030 "Write acceptance test for namespace update"
Task: T031 "Write acceptance test for namespace delete"
Task: T032 "Write acceptance test for namespace import"

# Launch model and schema in parallel:
Task: T033 "Define namespace data model in internal/resources/namespace/model.go"
Task: T034 "Define namespace schema in internal/resources/namespace/schema.go"
```

---

## Implementation Strategy

### MVP First (User Stories 1-2 Only)

1. Complete Phase 1: Setup
2. Complete Phase 2: Foundational (CRITICAL - blocks all stories)
3. Complete Phase 3: User Story 1 (Provider Authentication)
4. Complete Phase 4: User Story 2 (Namespace Resource)
5. **STOP and VALIDATE**: Test provider auth + namespace independently
6. Deploy/demo if ready - this is a functional MVP

### Incremental Delivery

1. Complete Setup + Foundational → Foundation ready
2. Add User Story 1 (Auth) → Test independently → MVP foundation
3. Add User Story 2 (Namespace) → Test independently → Minimal functional provider
4. Add User Story 3 (HTTP LB) → Test independently → Core use case
5. Add User Story 4-6 → Each tested independently → Full P1/P2 coverage
6. Add User Story 7-8 → Migration and data sources → Complete provider

### Recommended Priority Order

1. **P1 Stories** (US1, US2, US3): Core provider functionality
2. **P2 Stories** (US4, US5, US6, US7): Extended functionality
3. **P3 Stories** (US8): Enhancement features
4. **Polish**: Documentation and release preparation

---

## Summary

| Phase | User Story | Tasks | Parallel Tasks |
|-------|-----------|-------|----------------|
| 1 | Setup | 9 | 6 |
| 2 | Foundational | 9 | 4 |
| 3 | US1 - Auth (P1) | 10 | 5 |
| 4 | US2 - Namespace (P1) | 13 | 6 |
| 5 | US3 - HTTP LB (P1) | 13 | 6 |
| 6 | US4 - Origin Pool (P2) | 13 | 6 |
| 7 | US5 - Cloud Sites (P2) | 19 | 12 |
| 8 | US6 - App Firewall (P2) | 12 | 5 |
| 9 | US7 - Migration (P2) | 4 | 1 |
| 10 | US8 - Data Sources (P3) | 8 | 6 |
| 11 | Polish | 12 | 6 |
| **Total** | | **122** | **63** |

---

## Notes

- [P] tasks = different files, no dependencies on incomplete tasks
- [Story] label maps task to specific user story for traceability
- Each user story should be independently completable and testable
- Tests must fail before implementing (per constitution principle II)
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Follow terraform-plugin-framework patterns per constitution principle III
