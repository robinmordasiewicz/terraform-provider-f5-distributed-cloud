<!--
Sync Impact Report
==================
Version change: 1.0.0 (new constitution)
Modified principles: N/A (initial creation)
Added sections: Core Principles (5), Development Workflow, Quality Gates, Governance
Removed sections: None
Templates requiring updates:
  - .specify/templates/plan-template.md: ✅ Compatible (Constitution Check section exists)
  - .specify/templates/spec-template.md: ✅ Compatible (Requirements section aligns)
  - .specify/templates/tasks-template.md: ✅ Compatible (Test-first approach supported)
Follow-up TODOs: None
-->

# Terraform Provider for F5 Distributed Cloud Constitution

## Core Principles

### I. OpenAPI Specification Fidelity

All Terraform resources and data sources MUST be derived from and remain synchronized with
the official F5 Distributed Cloud OpenAPI specifications located in `docs/specifications/api/`.

- Resource schemas MUST mirror the OpenAPI schema definitions exactly
- Field names MUST match the API field names using appropriate Terraform naming conventions
- Required/optional field designation MUST reflect the OpenAPI specification
- Validation rules MUST be derived from OpenAPI constraints (enums, patterns, min/max)
- Any deviation from the OpenAPI spec MUST be documented with technical justification

**Rationale**: The provider serves as the Terraform interface to the F5 XC API. Drift between
the provider and API specification creates user confusion and integration failures.

### II. Test-First Development

All new resources, data sources, and features MUST follow test-driven development practices.

- Acceptance tests MUST be written before implementation begins
- Tests MUST fail initially (red phase) before implementation makes them pass (green phase)
- Unit tests MUST cover schema validation and state handling logic
- Integration tests MUST validate actual API interactions when feasible
- Test coverage MUST NOT decrease with new changes

**Rationale**: Terraform providers interact with external APIs where failures are costly.
Test-first ensures correctness and prevents regressions in production infrastructure.

### III. Idiomatic Terraform Patterns

All provider code MUST follow HashiCorp's Terraform Provider SDK conventions and
community-established patterns for Go-based providers.

- Use terraform-plugin-framework or terraform-plugin-sdk/v2 as appropriate
- Follow the standard CRUD operation patterns for resources
- Implement proper state management with accurate schema definitions
- Handle API errors gracefully with actionable user-facing messages
- Support import functionality for all resources where API permits

**Rationale**: Consistency with the Terraform ecosystem reduces user learning curve
and ensures compatibility with Terraform tooling and workflows.

### IV. Documentation Excellence

Every public resource, data source, and attribute MUST have comprehensive documentation.

- Resource documentation MUST include usage examples with realistic configurations
- All attributes MUST have descriptions explaining purpose and valid values
- Breaking changes MUST be documented in CHANGELOG with migration guidance
- Examples MUST be tested and validated against current provider version
- README MUST provide quickstart path for new users

**Rationale**: Terraform users depend on provider documentation for infrastructure-as-code
development. Poor documentation leads to misconfiguration and support burden.

### V. Semantic Versioning

All releases MUST follow semantic versioning with clear categorization of changes.

- MAJOR: Breaking changes to resource schemas, removed resources, or incompatible behaviors
- MINOR: New resources, data sources, or backward-compatible enhancements
- PATCH: Bug fixes, documentation updates, non-functional improvements
- Pre-release versions MUST be used for testing significant changes before stable release
- Deprecation MUST precede removal by at least one minor version

**Rationale**: Infrastructure teams depend on version stability. Clear versioning
enables safe upgrades and predictable dependency management.

## Development Workflow

### Code Review Requirements

- All changes MUST be submitted via pull request
- Pull requests MUST pass automated checks before merge
- Schema changes MUST include corresponding test updates
- API specification changes MUST trigger provider synchronization review

### Testing Gates

- Linting: `golangci-lint` MUST pass with zero errors
- Unit tests: All unit tests MUST pass
- Acceptance tests: Affected resource tests MUST pass
- Documentation: Examples MUST be syntactically valid

## Quality Gates

### Pre-Merge Checklist

- [ ] OpenAPI specification alignment verified
- [ ] Tests written and passing (red-green cycle completed)
- [ ] Documentation updated for user-facing changes
- [ ] CHANGELOG entry added for notable changes
- [ ] No decrease in test coverage

### Release Checklist

- [ ] All tests passing in CI
- [ ] Version number follows semantic versioning
- [ ] CHANGELOG reflects all changes since last release
- [ ] Documentation site reflects current version

## Governance

This constitution establishes the foundational principles for developing and maintaining
the Terraform Provider for F5 Distributed Cloud. All contributors MUST adhere to these
principles.

### Amendment Process

1. Propose amendment via pull request to this document
2. Document rationale for change and impact assessment
3. Obtain maintainer approval
4. Update dependent documentation and templates if affected
5. Increment constitution version appropriately

### Compliance Review

- Code reviews MUST verify adherence to constitution principles
- Automated checks SHOULD enforce measurable requirements where feasible
- Violations MUST be documented and remediated before merge

### Version Policy

- MAJOR: Principle removal, fundamental governance changes
- MINOR: New principles, significant guidance expansion
- PATCH: Clarifications, wording improvements, non-semantic updates

**Version**: 1.0.0 | **Ratified**: 2025-11-21 | **Last Amended**: 2025-11-21
