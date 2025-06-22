resource "aws_lambda_function" "checklist" {
  function_name = "checklist"
  role          = aws_iam_role.lambda_execution_role.arn

  filename         = data.archive_file.checklist.output_path
  source_code_hash = data.archive_file.checklist.output_base64sha256
  runtime          = "provided.al2023"
  handler          = "bootstrap"

  logging_config {
    log_format = "Text"
    log_group  = aws_cloudwatch_log_group.checklist.name
  }
}

// https://docs.aws.amazon.com/bedrock/latest/userguide/agents-permissions.html#agents-permissions-identity
resource "aws_lambda_permission" "checklist" {
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.checklist.function_name
  principal      = "bedrock.amazonaws.com"
  source_account = data.aws_caller_identity.current.account_id
  source_arn     = aws_bedrockagent_agent.reviewer.agent_arn
}

data "archive_file" "checklist" {
  type        = "zip"
  source_file = "${path.module}/files/checklist/bootstrap"
  output_path = "${path.module}/files/checklist/lambda.zip"
}

resource "aws_lambda_function" "reviewresult" {
  function_name = "reviewresult"
  role          = aws_iam_role.lambda_execution_role.arn

  filename         = data.archive_file.reviewresult.output_path
  source_code_hash = data.archive_file.reviewresult.output_base64sha256
  runtime          = "provided.al2023"
  handler          = "bootstrap"

  logging_config {
    log_format = "Text"
    log_group  = aws_cloudwatch_log_group.reviewresult.name
  }
}

// https://docs.aws.amazon.com/bedrock/latest/userguide/agents-permissions.html#agents-permissions-identity
resource "aws_lambda_permission" "reviewresult" {
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.reviewresult.function_name
  principal      = "bedrock.amazonaws.com"
  source_account = data.aws_caller_identity.current.account_id
  source_arn     = aws_bedrockagent_agent.reviewer.agent_arn
}

data "archive_file" "reviewresult" {
  type        = "zip"
  source_file = "${path.module}/files/reviewresult/bootstrap"
  output_path = "${path.module}/files/reviewresult/lambda.zip"
}
