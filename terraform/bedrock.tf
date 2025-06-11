resource "aws_bedrockagent_agent" "reviewer" {
  agent_name              = "reviewer"
  agent_resource_role_arn = aws_iam_role.bedrock_agent.arn
  foundation_model        = data.aws_bedrock_inference_profile.agent.inference_profile_arn
  instruction             = file("${path.module}/prompts/instruction.md")
  description             = "The agent reviewing the design documents stored in the knowledge base according to the requirements in the checklist."

  depends_on = [aws_iam_role_policy_attachment.bedrock_agent]
}

resource "aws_bedrockagent_agent_knowledge_base_association" "documents" {
  agent_id             = aws_bedrockagent_agent.reviewer.id
  knowledge_base_id    = aws_bedrockagent_knowledge_base.documents.id
  description          = "A knowledge base containing design documents."
  knowledge_base_state = "ENABLED"
}

resource "aws_bedrockagent_knowledge_base" "documents" {
  name        = "documents"
  role_arn    = aws_iam_role.bedrock_knowledge_base.arn
  description = "The knowledge base containing design documents for the reviewer agent."

  knowledge_base_configuration {
    vector_knowledge_base_configuration {
      embedding_model_arn = data.aws_bedrock_foundation_model.knowledge_base_embedding.model_arn
    }
    type = "VECTOR"
  }

  storage_configuration {
    type = "OPENSEARCH_SERVERLESS"
    opensearch_serverless_configuration {
      collection_arn    = aws_opensearchserverless_collection.bedrock_knowledge_base.arn
      vector_index_name = opensearch_index.default.name
      field_mapping {
        vector_field   = local.aoss_vector_field_name
        text_field     = local.aoss_text_field_name
        metadata_field = local.aoss_metadata_field_name
      }
    }
  }
}

resource "aws_bedrockagent_data_source" "documents" {
  name                 = aws_s3_bucket.bedrock_data_source.bucket
  knowledge_base_id    = aws_bedrockagent_knowledge_base.documents.id
  data_deletion_policy = "RETAIN"

  data_source_configuration {
    type = "S3"
    s3_configuration {
      bucket_arn = aws_s3_bucket.bedrock_data_source.arn
    }
  }
}

data "aws_bedrock_inference_profile" "agent" {
  inference_profile_id = var.bedrock_agent_inference_profile_id
}

data "aws_bedrock_foundation_model" "knowledge_base_embedding" {
  model_id = var.bedrock_knowledge_base_embedding_model_id
}
