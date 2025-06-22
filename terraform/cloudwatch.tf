resource "aws_cloudwatch_log_group" "checklist" {
  name = "/aws/lambda/checklist"
}

resource "aws_cloudwatch_log_group" "reviewresult" {
  name = "/aws/lambda/reviewresult"
}
