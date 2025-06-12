resource "time_sleep" "wait_agent_resource_role_creation" {
  create_duration = "10s"

  triggers = {
    policy_document = aws_iam_policy.bedrock_agent.policy
  }

  depends_on = [aws_iam_role_policy_attachment.bedrock_agent]
}
