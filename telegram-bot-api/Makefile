include ../include.mk

GCP_ENDPOINT_NAME = telegrambotapi.endpoints.hacker-playground.cloud.goog
DESCRIPTOR_FILE = api_descriptor.pb

compileProtocolBuffers:
	@echo "Compile protocol buffers and generate Go API"
	protoc \
		--include_imports \
		--include_source_info \
		--proto_path=$(GOOGLEAPIS_DIR) \
		--proto_path=. \
		--descriptor_set_out=$(DESCRIPTOR_FILE) \
		--go_out=plugins=grpc:. \
		telegram_bot_api.proto

gcpDeployEndpoint: compileProtocolBuffers
	@echo "Deploy the proto descriptor and the configuration to GCP"
	gcloud endpoints --project $(GCP_PROJECT) services deploy $(DESCRIPTOR_FILE) api_config.yaml

gcpDeleteEntpoint:
	gcloud endpoints --project $(GCP_PROJECT) services delete $(GCP_ENDPOINT_NAME)
