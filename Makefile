.DEFAULT_GOAL := lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	go clean -testcache

.PHONY: ingest
ingest:
	aws s3 sync ./data s3://$(shell terraform -chdir=terraform output -raw bedrock_data_source_s3_bucket_name)
	go tool ingest --knowledge-base-id "$(shell terraform -chdir=terraform output -raw bedrock_knowledge_base_id)" --data-source-id "$(shell terraform -chdir=terraform output -raw bedrock_data_source_id)"

.PHONY: review
review:
	xargs -- go tool review --agent-id "$(shell terraform -chdir=terraform output -raw bedrock_agent_id)" < internal/checklist/awswellarchitectedframework/urls.txt

.PHONY: generate
generate:
	go generate ./...
