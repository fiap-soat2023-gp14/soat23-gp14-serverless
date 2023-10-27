resource "aws_api_gateway_rest_api" "apigw-lambda-fiap-project" {
  name = "${var.domain_name}-${var.sufix}"
}

resource "aws_api_gateway_resource" "proxy" {
  rest_api_id = aws_api_gateway_rest_api.apigw-lambda-fiap-project.id
  parent_id   = aws_api_gateway_rest_api.apigw-lambda-fiap-project.root_resource_id
  path_part   = "{proxy+}"
}

resource "aws_api_gateway_method" "proxy" {
  rest_api_id   = aws_api_gateway_rest_api.apigw-lambda-fiap-project.id
  resource_id   = aws_api_gateway_resource.proxy.id
  http_method   = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "apigw-lambda-fiap-project" {
  rest_api_id = aws_api_gateway_rest_api.apigw-lambda-fiap-project.id
  resource_id = aws_api_gateway_method.proxy.resource_id
  http_method = aws_api_gateway_method.proxy.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.lambda-function-fiap-project.invoke_arn
}

resource "aws_api_gateway_deployment" "apigw-lambda-fiap-project" {
  depends_on = [aws_api_gateway_integration.apigw-lambda-fiap-project]

  rest_api_id = aws_api_gateway_rest_api.apigw-lambda-fiap-project.id
  stage_name  = "prod"
}

resource "aws_lambda_permission" "lambda_permission" {
  action        = "lambda:InvokeFunction"
  function_name = "${var.domain_name}-lambda"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.apigw-lambda-fiap-project.execution_arn}/*/*/*"

  depends_on = [aws_lambda_function.lambda-function-fiap-project]
}