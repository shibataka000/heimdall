.DEFAULT_GOAL := help

.PHONY: help
help:
	@cat README.md

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build: terraform/files/checklist/bootstrap terraform/files/reviewresult/bootstrap

.PHONY: clean
clean:
	$(RM) terraform/files/*/bootstrap terraform/files/*/*.zip

.PHONY: ingest
ingest:
	aws s3 sync ./data s3://$(shell terraform -chdir=terraform output -raw bedrock_data_source_s3_bucket_name)
	go tool ingest --knowledge-base-id "$(shell terraform -chdir=terraform output -raw bedrock_knowledge_base_id)" --data-source-id "$(shell terraform -chdir=terraform output -raw bedrock_data_source_id)"

.PHONY: review
review:
	jq '.[].id' internal/checklist/awswellarchitectedframework/requirements.json | grep "OPS06" | xargs -I {RequirementID} go tool review --agent-id "$(shell terraform -chdir=terraform output -raw bedrock_agent_id)" --checklist-id "aws-well-architected-framework" --requirement-id "{RequirementID}"

.PHONY: terraform-init
terraform-init: terraform/.terraform/terraform.tfstate

.PHONY: terraform-apply
terraform-apply: terraform-init build
	terraform -chdir=terraform apply

.PHONY: terraform-destroy
terraform-destroy: terraform-init build
	terraform -chdir=terraform destroy

terraform/files/%/bootstrap: $(shell find . -name *.go)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ -tags lambda.norpc cmd/lambda/$*/main.go

terraform/.terraform/terraform.tfstate:
	terraform -chdir=terraform init

tail-logs-checklist:
	aws logs tail $(shell terraform -chdir=terraform output -raw checklist_log_group_name) --follow --since 1d

tail-logs-reviewresult:
	aws logs tail $(shell terraform -chdir=terraform output -raw reviewresult_log_group_name) --follow  --since 1d
