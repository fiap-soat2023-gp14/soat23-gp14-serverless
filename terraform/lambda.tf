resource "aws_lambda_function" "lambda-function-fiap-project" {
  function_name = "${var.domain_name}-lambda"
  role          = aws_iam_role.iam_for_lambda.arn
  image_uri = "495428969620.dkr.ecr.us-east-1.amazonaws.com/fiapauthlambda:latest"
  architectures = ["x86_64"]
  package_type = "Image"

  environment {
    variables = local.environment_variables
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "${var.domain_name}-lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

data "aws_iam_policy_document" "policy_for_lambda" {
  statement {
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "xray:PutTraceSegments",
      "xray:PutTelemetryRecords",
    ]

    resources = [aws_cloudwatch_log_group.lambda-function-fiap-project.arn]
  }
}

resource "aws_iam_role_policy" "policy_for_lambda" {
  name   = "${var.domain_name}-lambda"
  role   = aws_iam_role.iam_for_lambda.id
  policy = data.aws_iam_policy_document.policy_for_lambda.json
}

resource "aws_cloudwatch_log_group" "lambda-function-fiap-project" {
  name              = "/aws/lambda/${var.domain_name}-${var.sufix}"
  retention_in_days = var.log_retention_in_days
}