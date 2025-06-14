resource "aws_lambda_function" "get_requirement" {
  function_name = "get-requirement"
  role          = aws_iam_role.lambda_execution_role.arn

  filename         = data.archive_file.get_requirement.output_path
  source_code_hash = data.archive_file.get_requirement.output_base64sha256
  runtime          = "provided.al2023"
  handler          = "bootstrap"

  logging_config {
    log_format = "Text"
    log_group  = aws_cloudwatch_log_group.get_requirement.name
  }
}

resource "aws_lambda_function" "save_review_result" {
  function_name = "save-review-result"
  role          = aws_iam_role.lambda_execution_role.arn

  filename         = data.archive_file.save_review_result.output_path
  source_code_hash = data.archive_file.save_review_result.output_base64sha256
  runtime          = "provided.al2023"
  handler          = "bootstrap"

  logging_config {
    log_format = "Text"
    log_group  = aws_cloudwatch_log_group.save_review_result.name
  }
}

data "archive_file" "get_requirement" {
  type        = "zip"
  source_file = "${path.module}/../dist/getrequirement"
  output_path = "get_requirement.zip"
}

data "archive_file" "save_review_result" {
  type        = "zip"
  source_file = "${path.module}/../dist/getrequirement"
  output_path = "save_review_result.zip"
}
