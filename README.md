# Heimdall

## What is Heimdall?

Heimdall reviews design documents and ensures that they satisfy the requirements specified in the checklist.

## Supported checklists

- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)

## Requirements

- AWS CLI
- Go
- Terraform
- [Access to Amazon Bedrock foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-access-modify.html)

## Setup

1. Construct the knowledge base and the agent.

```bash
terraform -chdir=terraform init
terraform -chdir=terraform apply
```

## Usage

1. Store the design documents in the [data](./data/) directory.
2. Send the design documents to the knowledge base.

```bash
make ingest
```

3. Review the design documents.

```bash
make review
```

## Cleanup

```bash
terraform -chdir=terraform destroy
```
