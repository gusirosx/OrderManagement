# ================ Global Params ================ #
PROTO_FLAG = -I proto/ proto/ordermgn.proto # Proto source
#GO Flags
GO_OPT = --go_opt=paths=source_relative
GRPC_OPT = --go-grpc_opt=paths=source_relative

# ========= Proto Destination Services ========== #
SRV_PROTO = server/ecommerce
CLI_PROTO = client/ecommerce

define GenerateProtoFiles
	@protoc --go_out=$(1) $(GO_OPT) --go-grpc_out=$(1) $(GRPC_OPT) $(PROTO_FLAG)
	@echo '---> Generating $(2) Protofiles'
endef

generate: 
	$(call GenerateProtoFiles,$(SRV_PROTO),Service)
	$(call GenerateProtoFiles,$(CLI_PROTO),Client)

run_server:
	@echo "---- Running Server ----"
	@go run server/*

run_client:
	@echo "---- Running Client ----"
	@go run client/*
