# Example: User Identification
resource "f5xc_user_identification" "example" {
  name        = "my-user-identification"
  namespace   = "my-namespace"
  description = "Example user identification policy for tracking users"
}
