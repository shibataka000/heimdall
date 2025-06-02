resource "aws_s3_bucket" "bedrock_data_source" {
  bucket_prefix = "documents-"
  force_destroy = true
}
