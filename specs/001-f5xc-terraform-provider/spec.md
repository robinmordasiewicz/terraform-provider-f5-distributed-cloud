# Feature Specification: F5 Distributed Cloud Terraform Provider (Open Source)

**Feature Branch**: `001-f5xc-terraform-provider`
**Created**: 2025-11-21
**Status**: Draft
**Input**: User description: "Create an open source terraform resource provider for the F5 distributed cloud as an alternative to the proprietary volterraedge/volterra provider"

## Overview

This project creates an **open source Terraform provider** for F5 Distributed Cloud (F5 XC) services. The provider will be functionally equivalent to the proprietary `volterraedge/volterra` provider (https://registry.terraform.io/providers/volterraedge/volterra/latest) but built from the public OpenAPI specifications available in `docs/specifications/api/`.

**Key Differentiator**: This is a community-driven, open source implementation providing transparency, community contributions, and flexibility that the proprietary provider does not offer.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provider Authentication Configuration (Priority: P1)

Infrastructure engineers need to configure the Terraform provider with valid F5 Distributed Cloud credentials to authenticate API requests.

**Why this priority**: Without authentication, no other provider functionality can work. This is the foundational capability that enables all subsequent resource management.

**Independent Test**: Can be fully tested by initializing the provider with credentials and verifying successful connection to the F5 XC API endpoint.

**Acceptance Scenarios**:

1. **Given** valid API certificate and key files, **When** provider is initialized with certificate paths, **Then** provider authenticates successfully and is ready to manage resources
2. **Given** valid API token, **When** provider is initialized with token in Authorization header format, **Then** provider authenticates successfully using token-based auth
3. **Given** invalid or expired credentials, **When** provider attempts to initialize, **Then** provider returns clear authentication error with guidance on credential generation
4. **Given** missing required credential configuration, **When** user runs terraform init, **Then** provider displays helpful error message indicating which credentials are needed

---

### User Story 2 - Namespace Resource Management (Priority: P1)

Infrastructure engineers need to create, read, update, and delete namespaces as the fundamental organizational unit for all F5 XC resources.

**Why this priority**: Namespaces are required containers for most other resources. Managing namespaces is essential before creating load balancers, origin pools, or security policies.

**Independent Test**: Can be fully tested by creating a namespace resource, verifying it exists in F5 XC console, modifying its description, and deleting it.

**Acceptance Scenarios**:

1. **Given** a configured provider, **When** user defines a namespace resource with a unique name, **Then** namespace is created in the specified F5 XC tenant
2. **Given** an existing namespace managed by Terraform, **When** user modifies namespace attributes (labels, description), **Then** changes are applied to the F5 XC namespace
3. **Given** an existing namespace managed by Terraform, **When** user removes the resource from configuration and runs apply, **Then** namespace is deleted from F5 XC
4. **Given** a namespace exists in F5 XC but not in Terraform state, **When** user imports the namespace, **Then** Terraform state reflects the current namespace configuration

---

### User Story 3 - HTTP Load Balancer Resource Management (Priority: P1)

Infrastructure engineers need to create and manage HTTP load balancers to distribute traffic to application backends.

**Why this priority**: HTTP load balancers are the primary traffic management resource that most users need to configure for web applications.

**Independent Test**: Can be fully tested by creating an HTTP load balancer with a domain, verifying it appears in F5 XC console, and confirming it can receive traffic.

**Acceptance Scenarios**:

1. **Given** a configured provider and existing namespace, **When** user defines an HTTP load balancer with domain and origin pool, **Then** load balancer is created and advertised on F5 XC global network
2. **Given** an existing HTTP load balancer, **When** user adds WAF policy or rate limiting configuration, **Then** security settings are applied to the load balancer
3. **Given** an existing HTTP load balancer, **When** user modifies routing rules or origin pools, **Then** traffic routing changes are applied without service interruption
4. **Given** an HTTP load balancer with dependencies, **When** user attempts to delete it, **Then** proper dependency validation occurs before deletion

---

### User Story 4 - Origin Pool Resource Management (Priority: P2)

Infrastructure engineers need to define origin pools that specify backend servers for load balancers.

**Why this priority**: Origin pools are required by load balancers to route traffic to actual application servers. They are a core dependency but not standalone-useful.

**Independent Test**: Can be fully tested by creating an origin pool with backend endpoints and verifying the configuration in F5 XC console.

**Acceptance Scenarios**:

1. **Given** a configured provider and namespace, **When** user defines an origin pool with backend endpoints, **Then** origin pool is created with specified servers
2. **Given** an origin pool with health checks configured, **When** user applies configuration, **Then** health monitoring is enabled for backend endpoints
3. **Given** an origin pool referenced by load balancers, **When** user modifies endpoint list, **Then** changes propagate to dependent load balancers

---

### User Story 5 - Cloud Site Resource Management (Priority: P2)

Infrastructure engineers need to deploy and manage F5 XC sites in cloud environments (AWS, Azure, GCP).

**Why this priority**: Cloud sites enable F5 XC to integrate with existing cloud infrastructure, but many users may use F5's global network without deploying their own sites.

**Independent Test**: Can be fully tested by creating an AWS VPC site configuration and verifying site registration in F5 XC console.

**Acceptance Scenarios**:

1. **Given** valid cloud credentials configured in F5 XC, **When** user defines an AWS VPC site resource, **Then** site is provisioned in the specified AWS region
2. **Given** an existing cloud site, **When** user modifies networking configuration, **Then** site is updated with new network settings
3. **Given** an existing cloud site with running workloads, **When** user requests deletion, **Then** proper warning and confirmation workflow executes

---

### User Story 6 - Application Firewall Policy Management (Priority: P2)

Security engineers need to define and apply web application firewall (WAF) policies to protect applications.

**Why this priority**: WAF policies are essential for security but require load balancers to be functional first.

**Independent Test**: Can be fully tested by creating a WAF policy with specific rules and attaching it to an HTTP load balancer.

**Acceptance Scenarios**:

1. **Given** a configured provider, **When** user defines an app_firewall resource with detection rules, **Then** WAF policy is created with specified protections
2. **Given** an existing WAF policy, **When** user adds exclusion rules for specific endpoints, **Then** exceptions are applied to the policy
3. **Given** a WAF policy attached to load balancers, **When** policy is updated, **Then** changes apply to all attached load balancers

---

### User Story 7 - Migration from Proprietary Provider (Priority: P2)

Organizations using the proprietary `volterraedge/volterra` provider need to migrate to this open source alternative with minimal disruption.

**Why this priority**: Many potential users already have infrastructure managed by the proprietary provider. Easy migration accelerates adoption.

**Independent Test**: Can be fully tested by taking an existing Terraform configuration using `volterra_*` resources and updating it to use the new provider with state import.

**Acceptance Scenarios**:

1. **Given** existing Terraform configuration using `volterra_*` resources, **When** user updates provider source and resource names, **Then** configuration is compatible with the new provider
2. **Given** existing resources managed by proprietary provider, **When** user imports resources into new provider state, **Then** all resource attributes are correctly captured
3. **Given** mixed environment with both providers, **When** user operates the new provider alongside existing tooling, **Then** no conflicts occur

---

### User Story 8 - Data Source Queries (Priority: P3)

Infrastructure engineers need to query existing F5 XC resources for use in Terraform configurations.

**Why this priority**: Data sources enable integration with existing infrastructure but are not required for greenfield deployments.

**Independent Test**: Can be fully tested by querying an existing namespace and using its attributes in other resource definitions.

**Acceptance Scenarios**:

1. **Given** existing resources in F5 XC, **When** user defines a data source query, **Then** resource attributes are available for reference in other resources
2. **Given** a query for a non-existent resource, **When** Terraform plan runs, **Then** clear error message indicates resource not found

---

### Edge Cases

- What happens when API rate limits are exceeded during large deployments?
- How does the provider handle partial resource creation failures during apply?
- What happens when concurrent Terraform runs attempt to modify the same resource?
- How does the provider handle F5 XC API version changes or deprecations?
- What happens when a resource is modified outside of Terraform (drift detection)?
- How does migration handle resources with configuration drift between providers?

## Requirements *(mandatory)*

### Functional Requirements

**Core Provider**
- **FR-001**: Provider MUST support authentication via API certificate (P12 format with extracted cert/key)
- **FR-002**: Provider MUST support authentication via API token in Authorization header
- **FR-003**: Provider MUST allow configuration of tenant-specific API endpoint URL
- **FR-004**: Provider MUST be published as open source with permissive licensing

**Resource Management**
- **FR-005**: Provider MUST implement full CRUD operations for namespace resources
- **FR-006**: Provider MUST implement full CRUD operations for HTTP load balancer resources
- **FR-007**: Provider MUST implement full CRUD operations for origin pool resources
- **FR-008**: Provider MUST implement full CRUD operations for application firewall resources
- **FR-009**: Provider MUST implement full CRUD operations for TCP load balancer resources
- **FR-010**: Provider MUST implement full CRUD operations for DNS load balancer resources
- **FR-011**: Provider MUST implement full CRUD operations for CDN load balancer resources
- **FR-012**: Provider MUST support cloud site resources for AWS (VPC and TGW), Azure (VNet), and GCP deployments
- **FR-013**: Provider MUST implement cloud credentials management resources
- **FR-014**: Provider MUST implement certificate management resources
- **FR-015**: Provider MUST implement route resources for traffic routing control
- **FR-016**: Provider MUST implement discovery resources for service discovery

**Data Sources**
- **FR-017**: Provider MUST implement data sources for querying existing F5 XC resources
- **FR-018**: Provider MUST support data sources for all resource types that support read operations

**State Management**
- **FR-019**: Provider MUST support resource import for state synchronization
- **FR-020**: Provider MUST maintain state consistency with F5 XC API during plan/apply operations
- **FR-021**: Provider MUST support resource references between related objects

**Quality & Reliability**
- **FR-022**: Provider MUST validate resource configurations before API submission
- **FR-023**: Provider MUST handle API errors and return actionable error messages
- **FR-024**: Provider MUST handle API pagination for list operations
- **FR-025**: Provider MUST implement retry logic with exponential backoff for transient API failures

### Key Entities

**Core Resources (matching proprietary provider)**
- **Provider Configuration**: Tenant URL, authentication credentials (certificate path, key path, or API token), timeout settings
- **Namespace**: Logical workspace container - name, labels, annotations, description
- **HTTP Load Balancer** (`http_loadbalancer`): Traffic distribution - domains, routes, origin pools, WAF policies, TLS settings
- **TCP Load Balancer** (`tcp_loadbalancer`): Layer 4 TCP traffic distribution
- **DNS Load Balancer** (`dns_load_balancer`): DNS-based traffic distribution
- **CDN Load Balancer** (`cdn_loadbalancer`): Content delivery network configuration
- **Origin Pool** (`origin_pool`): Backend server definitions - endpoints, health checks, load balancing algorithms
- **App Firewall** (`app_firewall`): WAF policy configuration - detection mode, rule sets, exclusions
- **Cloud Credentials** (`cloud_credentials`): Cloud provider authentication for site provisioning
- **Certificate** (`certificate`): TLS certificate management
- **Route** (`route`): Traffic routing rules
- **Discovery** (`discovery`): Service and network discovery configuration

**Cloud Site Resources**
- **AWS VPC Site** (`aws_vpc_site`): AWS VPC-based site deployment
- **AWS TGW Site** (`aws_tgw_site`): AWS Transit Gateway site deployment
- **Azure VNet Site** (`azure_vnet_site`): Azure Virtual Network site deployment
- **GCP VPC Site** (`gcp_vpc_site`): Google Cloud VPC site deployment

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Users can complete provider authentication setup in under 5 minutes with documentation
- **SC-002**: Users can create a functional HTTP load balancer with origin pool in under 10 minutes
- **SC-003**: Terraform plan accurately reflects differences between configuration and F5 XC state for 95% of scenarios
- **SC-004**: Resource creation, update, and deletion operations complete within expected timeframes (simple resources under 30 seconds, cloud sites under 10 minutes)
- **SC-005**: Provider handles 50+ resources in a single Terraform configuration without performance degradation
- **SC-006**: 90% of common F5 XC configuration tasks can be accomplished through the provider without console access
- **SC-007**: Error messages enable users to resolve issues without external support for 80% of error scenarios
- **SC-008**: Provider supports import for all managed resource types, enabling adoption of existing infrastructure
- **SC-009**: Migration from proprietary provider can be completed in under 1 hour for configurations with fewer than 50 resources
- **SC-010**: Provider achieves feature parity with 80% of proprietary provider resources within initial release

## Assumptions

- Users have valid F5 Distributed Cloud tenant with API access enabled
- Users have generated API credentials (certificate or token) from the F5 XC console
- The 269 OpenAPI specification files in `docs/specifications/api/` accurately represent the current F5 XC API
- F5 XC API follows RESTful patterns with consistent CRUD operations across resource types
- API base URL follows pattern: `https://<tenant>.console.ves.volterra.io/api`
- Primary service prefixes are `/api/config/` for configuration and `/api/web/` for queries
- Resources are organized by namespace (system, shared, default, or user-defined)
- Resource naming and schema should align closely with the proprietary `volterraedge/volterra` provider for easy migration
- The proprietary provider documentation at https://registry.terraform.io/providers/volterraedge/volterra/latest/docs serves as functional reference
- Open source licensing allows community contributions and modifications

## Out of Scope

- Custom provider plugins or extensions beyond standard Terraform provider patterns
- Migration tools that automatically convert proprietary provider configurations (manual process supported)
- Real-time monitoring or alerting integration
- Cost management or billing API integration
- Multi-tenant management across multiple F5 XC organizations
- Proprietary features not exposed through the public F5 XC API
