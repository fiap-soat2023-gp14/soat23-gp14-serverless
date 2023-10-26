resource "aws_cognito_user_pool" "up-fiap-project" {
  name             = "auth-up-fiap-project"
  alias_attributes = ["preferred_username"]

  # This setting is what actually makes the confirmation code to be sent
  auto_verified_attributes = ["email"]

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }

  # Set to TRUE for project purposes, in real life should be FALSE (default)
  admin_create_user_config {
    allow_admin_create_user_only = "true"
  }

  schema {
    name                     = "cpf"
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = false
    required                 = false
    string_attribute_constraints {
      min_length = 11
      max_length = 11
    }
  }
}

resource "aws_cognito_user_pool_client" "up-client-fiap-project" {
  name = "auth-up-client-fiap-project"

  user_pool_id = aws_cognito_user_pool.up-fiap-project.id

  explicit_auth_flows = ["ADMIN_NO_SRP_AUTH", "USER_PASSWORD_AUTH"]
}

resource "aws_cognito_user_pool_domain" "up-domain-fiap-project" {
  domain       = var.domain_name
  user_pool_id = aws_cognito_user_pool.up-fiap-project.id
}
