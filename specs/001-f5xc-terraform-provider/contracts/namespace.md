# Resource Contract: f5xc_namespace

## Resource Schema

```hcl
resource "f5xc_namespace" "example" {
  name        = string           # Required: DNS-1035 format
  description = string           # Optional: Max 1200 bytes
  labels      = map(string)      # Optional
  annotations = map(string)      # Optional
  disable     = bool             # Optional: default false
}
```

## CRUD Operations

### Create
```
API: POST /api/web/namespaces

Request:
{
  "metadata": {
    "name": "<name>",
    "description": "<description>",
    "labels": { ... },
    "annotations": { ... },
    "disable": false
  },
  "spec": {}
}

Preconditions:
  - name is unique within tenant
  - name follows DNS-1035 format (lowercase, alphanumeric, hyphens)

Postconditions:
  - Namespace exists in F5 XC
  - State contains computed uid, creation_timestamp

Errors:
  - 409 Conflict: Namespace already exists
  - 400 Bad Request: Invalid name format
```

### Read
```
API: GET /api/web/namespaces/<name>

Postconditions:
  - State synchronized with F5 XC
  - All attributes populated from API response

Errors:
  - 404 Not Found: Mark resource for recreation
```

### Update
```
API: PUT /api/web/namespaces/<name>

Request: Same as Create

Preconditions:
  - Namespace exists
  - name cannot be changed (ForceNew)

Postconditions:
  - Updated attributes reflected in F5 XC
  - State synchronized

Errors:
  - 404 Not Found: Resource deleted externally
  - 409 Conflict: Concurrent modification
```

### Delete
```
API: DELETE /api/web/namespaces/<name>

Preconditions:
  - Namespace has no child resources (or fail_if_referred=false)

Postconditions:
  - Namespace removed from F5 XC
  - State removed

Errors:
  - 404 Not Found: Already deleted (no-op)
  - 409 Conflict: Namespace not empty
```

### Import
```
Import ID: <name>

terraform import f5xc_namespace.example production

Postconditions:
  - State populated from existing namespace
  - All attributes synchronized
```

## Validation Rules

| Attribute | Rule |
|-----------|------|
| name | DNS-1035: `^[a-z]([-a-z0-9]*[a-z0-9])?$`, max 63 chars |
| description | Max 1200 bytes |
| labels.key | Min 1, max 64 chars |
| labels.value | Min 1, max 1024 chars |
| annotations.key | Min 1, max 64 chars |
| annotations.value | Min 1, max 1024 chars |
