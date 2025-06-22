# Heimdall

## What is Heimdall?

Heimdall reviews design documents and ensures that they satisfy the requirements specified in the checklist.

## Supported checklists

- [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html)

## Requirements

- [AWS CLI](https://aws.amazon.com/cli/)
- [Go](https://go.dev/)
- [Terraform](https://developer.hashicorp.com/terraform)
- [jq](https://jqlang.org/)
- [Access to Amazon Bedrock foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-access-modify.html)

## Setup

1. Construct the knowledge base and the agent.

```bash
make terraform-apply
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

> [!NOTE]
>
> By default, the `make review` command reviews design documents against **ONLY** the [OPS06](https://docs.aws.amazon.com/wellarchitected/latest/framework/ops-06.html) requirements in the [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html).
> To review against **ALL** requirements in the [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html), remove the `grep "OPS06"` command from the [Makefile](./Makefile).

4. Show the review results.

```bash
make tail-logs-reviewresult
```

## Cleanup

```bash
make terraform-destroy
```
