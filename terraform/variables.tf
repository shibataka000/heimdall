variable "bedrock_agent_foundation_model_id" {
  default     = "anthropic.claude-3-5-sonnet-20241022-v2:0"
  description = "The foundation model ID for the agent."
  type        = string
}

variable "bedrock_knowledge_base_embedding_model_id" {
  default     = "amazon.titan-embed-text-v2:0"
  description = "The embedding model ID for the knowledge base."
  type        = string
}

variable "bedrock_knowledge_base_dimension" {
  default     = 1024
  description = "The number of dimensions in the vector field in the vector store for the knowledge base. See https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup.html for more details."
  type        = number
}
