# Const
#===============================================================
ZIP_DIR  := deployments/zip
LOG_DIR  := log

# MAIN_PATHを指定する
.PHONY: build_and_create_lambda
build_and_create_lambda: build packaging create_lambda

# MAIN_PATHを指定する
.PHONY: build_and_deploy_lambda
build_and_deploy_lambda: build packaging deploy_lambda

# BIN_PATHを指定する
# Lambdaにコードをデプロイするためにzip圧縮する
.PHONY: packaging
packaging:
	$(eval ZIP_PATH := $(subst $(BIN_DIR)/,$(ZIP_DIR)/,$(BIN_PATH)).zip)
	@mkdir -p $(dir $(ZIP_PATH));\
	rm -f $(ZIP_PATH);\
	echo Packaging to $(ZIP_PATH);\
	zip -j $(ZIP_PATH) $(BIN_PATH)

# ZIP_PATHを指定する
# localstackにLambdaを生成してコードをデプロイする
.PHONY: make_lambda
make_lambda:
	$(eval FUNC_NAME := $(basename $(notdir $(ZIP_PATH))))
	$(eval region := $(if $(REGION),$(REGION),ap-northeast-1))
	@echo create $(ZIP_NAME) Lambda;\
	aws --endpoint-url=http://localhost:4574 --profile localstack lambda create-function --profile=localstack --region=$(region) --function-name=$(FUNC_NAME) --runtime=go1.x --role=role --handler=$(FUNC_NAME) --zip-file fileb://$(ZIP_PATH)

# ZIP_PATHを指定する
# localstackのLambdaにコードをデプロイする
.PHONY: deploy_lambda
deploy_lambda:
	$(eval FUNC_NAME := $(basename $(notdir $(ZIP_PATH))))
	$(eval region := $(if $(REGION),$(REGION),ap-northeast-1))
	@echo deploy $(ZIP_PATH) to $(ZIP_NAME) Lambda;\
	aws --endpoint-url=http://localhost:4574 --profile localstack lambda update-function-code --profile=localstack --region=$(region) --function-name=$(FUNC_NAME) --zip-file fileb://$(ZIP_PATH)

# ZIP_PATHを指定する
# localstackのLambdaを呼び出す
.PHONY: invoke_lambda
invoke_lambda:
	$(eval FUNC_NAME := $(basename $(notdir $(ZIP_PATH))))
	$(eval LOG_PATH := $(basename $(subst $(ZIP_DIR)/,$(LOG_DIR)/,$(ZIP_PATH))).log)
	$(eval payload := $(if $(PAYLOAD),--payload '$(PAYLOAD)',))
	$(eval region := $(if $(REGION),$(REGION),ap-northeast-1))
	@echo invoke $(FUNC_NAME) Lambda;\
	aws lambda --endpoint-url=http://localhost:4574 invoke $(payload) --profile=localstack --region=$(region) --function-name=$(FUNC_NAME) $(LOG_PATH)
