resource "aws_bedrockagent_agent" "reviewer" {
  agent_name              = "reviewer"
  agent_resource_role_arn = aws_iam_role.bedrock_agent.arn
  foundation_model        = data.aws_bedrock_inference_profile.agent.inference_profile_arn
  instruction             = file("${path.module}/files/instruction.md")
  description             = "The agent reviewing the design documents stored in the knowledge base according to the requirements in the checklist."

  depends_on = [time_sleep.wait_agent_resource_role_creation]
}

resource "aws_bedrockagent_agent_knowledge_base_association" "documents" {
  agent_id             = aws_bedrockagent_agent.reviewer.id
  knowledge_base_id    = aws_bedrockagent_knowledge_base.documents.id
  description          = "A knowledge base containing design documents."
  knowledge_base_state = "ENABLED"

  depends_on = [
    aws_bedrockagent_agent_action_group.checklist,
    aws_bedrockagent_agent_action_group.reviewresult,
  ]
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

resource "aws_bedrockagent_agent_action_group" "checklist" {
  action_group_name          = aws_lambda_function.checklist.function_name
  agent_id                   = aws_bedrockagent_agent.reviewer.agent_id
  agent_version              = "DRAFT"
  description                = "チェックリストの要件を提供します。"
  skip_resource_in_use_check = true

  action_group_executor {
    lambda = aws_lambda_function.checklist.arn
  }
  api_schema {
    payload = file("${path.module}/files/checklist/api_schema.yaml")
  }
}

resource "aws_bedrockagent_agent_action_group" "reviewresult" {
  action_group_name          = aws_lambda_function.reviewresult.function_name
  agent_id                   = aws_bedrockagent_agent.reviewer.agent_id
  agent_version              = "DRAFT"
  description                = "レビュー結果を保存します。"
  skip_resource_in_use_check = true

  action_group_executor {
    lambda = aws_lambda_function.reviewresult.arn
  }
  api_schema {
    payload = file("${path.module}/files/reviewresult/api_schema.yaml")
  }

  depends_on = [aws_bedrockagent_agent_action_group.checklist]
}

data "aws_bedrock_inference_profile" "agent" {
  inference_profile_id = var.bedrock_agent_inference_profile_id
}

data "aws_bedrock_foundation_model" "knowledge_base_embedding" {
  model_id = var.bedrock_knowledge_base_embedding_model_id
}
