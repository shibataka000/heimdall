resource "aws_cloudwatch_log_group" "get_requirement" {
  name = "/aws/lambda/get-requirement"
}

resource "aws_cloudwatch_log_group" "save_review_result" {
  name = "/aws/lambda/save-review-result"
}
