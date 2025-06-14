locals {
  aoss_collection_name     = "bedrock-knowledge-base-${random_string.aoss_collection_name_suffix.result}"
  aoss_vector_field_name   = "bedrock-knowledge-base-default-vector"
  aoss_metadata_field_name = "AMAZON_BEDROCK_METADATA"
  aoss_text_field_name     = "AMAZON_BEDROCK_TEXT_CHUNK"

  lambda_function_get_requirement_image_uri    = "${aws_ecr_repository.get_requirement.repository_url}:latest"
  lambda_function_save_review_result_image_uri = "${aws_ecr_repository.save_review_result.repository_url}:latest"
}
