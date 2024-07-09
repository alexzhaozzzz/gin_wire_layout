go_version = $(shell go version)
commit_id = $(shell git rev-parse HEAD)
branch_name = $(shell git name-rev --name-only HEAD)
build_time = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
work_dir = bin
app_version = 1.0.0
version_package = gin_wire_layout
app_name = gin_wire_layout

.PHONY: target
# target
target:
	@mkdir -p ${work_dir}

.PHONY: build
# build
build: target
	@go build -ldflags \
	"-X ${version_package}.CommitId=${commit_id} \
	-X ${version_package}.BranchName=${branch_name} \
	-X ${version_package}.BuildTime=${build_time} \
	-X ${version_package}.AppVersion=${app_version}" -v \
	-buildvcs=false \
	-o ${work_dir}/${app_name} ./cmd/${app_name}/.

.PHONY: package
# package
.ONESHELL:
package: build
	@# 使用tar命令对${word_dir下面的文件打包}
	cp -r configs  ${work_dir}/
	cp ./start.sh ${work_dir}

clean:
	@rm -rf ${work_dir}


# show HELP
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
