output "bedrock_agent_id" {
  value       = aws_bedrockagent_agent.reviewer.id
  description = "The agent ID."
}

output "bedrock_knowledge_base_id" {
  value       = aws_bedrockagent_knowledge_base.documents.id
  description = "The knowledge base ID."
}

output "bedrock_data_source_id" {
  value       = aws_bedrockagent_data_source.documents.data_source_id
  description = "The data source ID."
}

output "bedrock_data_source_s3_bucket_name" {
  value       = aws_s3_bucket.bedrock_data_source.bucket
  description = "The S3 bucket name for the data source."
}

output "checklist_log_group_name" {
  value       = aws_cloudwatch_log_group.checklist.name
  description = "The log group name for the lambda function 'checklist'."
}

output "reviewresult_log_group_name" {
  value       = aws_cloudwatch_log_group.reviewresult.name
  description = "The log group name for the lambda function 'reviewresult'."
}
