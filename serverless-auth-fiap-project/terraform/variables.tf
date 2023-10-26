variable "function_name" {
  description = "Function name"
  default     = "authLambdaFiapProject"
}

variable "stage_name" {
  description = "Api version number"
  default     = "v1"
}

variable "domain_name" {
  default = "gp14-fiap-project"
}

variable "environment_variables" {
  description = "Map with environment variables for the function"

  default = {
    myenvvar = "test"
  }
}