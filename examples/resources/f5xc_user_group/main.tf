# Example: User Group
resource "f5xc_user_group" "example" {
  name        = "my-user-group"
  namespace   = "system"
  description = "Example user group for access management"
}
