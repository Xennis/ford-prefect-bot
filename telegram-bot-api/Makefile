compileProtocolBuffers:
	protoc \
		--include_imports \
		--include_source_info \
		--proto_path=${GOOGLEAPIS_DIR} \
		--proto_path=. \
		--descriptor_set_out=api_descriptor.pb \
		--go_out=plugins=grpc:. \
		telegram_bot_api.proto
