# Provider Contract: f5xc

## Provider Configuration Schema

```hcl
provider "f5xc" {
  api_url    = string           # Required: Tenant API URL
  api_token  = string           # Optional: API token auth (conflicts with cert auth)
  cert_file  = string           # Optional: Certificate file path
  key_file   = string           # Optional: Key file path
  timeout    = string           # Optional: Request timeout (default: "30s")
}
```

## Authentication Contract

### Certificate Authentication
```
Preconditions:
  - cert_file exists and is readable
  - key_file exists and is readable
  - Certificate is valid and not expired

Postconditions:
  - Provider authenticated to F5 XC API
  - All subsequent API calls include mTLS credentials
```

### Token Authentication
```
Preconditions:
  - api_token is non-empty
  - Token is valid and not expired

Postconditions:
  - Provider authenticated to F5 XC API
  - All subsequent API calls include Authorization header
```

## Error Contract

| Error Code | Provider Behavior |
|------------|------------------|
| 401 | Return authentication error, suggest credential check |
| 403 | Return permission error, include resource/action context |
| 404 | Return not found error for Read; no-op for Delete |
| 409 | Return conflict error, include current state hint |
| 429 | Retry with exponential backoff (max 5 attempts) |
| 500 | Retry with exponential backoff (max 3 attempts) |
| 503 | Retry with exponential backoff (max 3 attempts) |
| 504 | Retry with exponential backoff (max 3 attempts) |

## HTTP Client Contract

```
Base URL Pattern: https://<tenant>.console.ves.volterra.io/api

Default Headers:
  - Content-Type: application/json
  - Accept: application/json
  - Authorization: APIToken <token>  (if token auth)

Request Timeout: configurable, default 30s
Retry Policy: exponential backoff, 1s/2s/4s/8s/16s
Max Retries: 5 for rate limits, 3 for server errors
```
