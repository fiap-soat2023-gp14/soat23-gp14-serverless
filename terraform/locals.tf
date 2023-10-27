locals {
  computed_environment_variables = {
    "AWS_CLIENT_ID" = aws_cognito_user_pool_client.client.id
    "AWS_USER_POOL_REGION" = var.default_region
    "AWS_USER_POOL_ID" = aws_cognito_user_pool.user-pool.id
  }
  environment_variables = merge(local.computed_environment_variables, var.environment_variables)
}