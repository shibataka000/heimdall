resource "aws_lambda_function" "get_requirement" {
  function_name = "get-requirement"
  role          = aws_iam_role.lambda_execution_role.arn

  package_type = "Image"
  image_uri    = local.lambda_function_get_requirement_image_uri

  logging_config {
    log_format = "Text"
    log_group  = aws_cloudwatch_log_group.get_requirement.name
  }
}

resource "aws_lambda_function" "save_review_result" {
  function_name = "save-review-result"
  role          = aws_iam_role.lambda_execution_role.arn

  package_type = "Image"
  image_uri    = local.lambda_function_save_review_result_image_uri

  logging_config {
    log_format = "Text"
    log_group  = aws_cloudwatch_log_group.save_review_result.name
  }
}
