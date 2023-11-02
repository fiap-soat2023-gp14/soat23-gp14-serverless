variable "function_name" {
  description = "Function name"
  default     = "auth-lambda"
}

variable "stage_name" {
  description = "Api version number"
  default     = "v1"
}

variable "domain_name" {
  default = "gp-14"
}

variable "sufix" {
  default = "fiap-project"
}

variable "environment_variables" {
  description = "Map with environment variables for the function"

  default = {
    myenvvar = "test"
  }
}

variable "default_region" {
    default = "us-east-1"
}

variable "log_retention_in_days" {
  default = 1
}