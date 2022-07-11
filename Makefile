#
# Copyright 2022 DSR Corporation, Denver, Colorado.
# https://www.dsr-corporation.com
#
# This file is part of ssi-medical-prescriptions-demo.
#
# ssi-medical-prescriptions-demo is free software: you can redistribute it
# and/or modify it under the terms of the GNU Affero General Public License
# as published by the Free Software Foundation, either version 3 of the License,
# or (at your option) any later version.
#
# ssi-medical-prescriptions-demo is distributed in the hope that it will be
# useful, but WITHOUT ANY WARRANTY; without even the implied warranty
# of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
# See the GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License along
# with ssi-medical-prescriptions-demo. If not, see <https://www.gnu.org/licenses/>.
#

GO_CMD ?= go
DEMO_SERVER_PATH = cmd/demo-server
MOCK_SERVER_PATH=cmd/mock-server

# Namespace for the images
DOCKER_OUTPUT_NS   ?= ssi-medical-prescriptions-demo
DEMO_SERVER_IMAGE_NAME 	?= demo-server
MOCK_SERVER_IMAGE_NAME 	?= mock-server
DEMO_SERVER_IMAGE_TAG 	?= latest
MOCK_SERVER_IMAGE_TAG 	?= latest

# Tool commands (overridable)
DOCKER_CMD ?= docker
GO_CMD     ?= go
ALPINE_VER ?= 3.16
GO_TAGS    ?=
GO_VER ?= 1.18.3

.PHONY: demo-server
demo-server:
	@echo "Building demo-server"
	@mkdir -p ./build/bin
	@cd ${DEMO_SERVER_PATH} && go build -o ../../build/bin/demo-server main.go


.PHONY: mock-server
mock-server:
	@echo "Building mock-server"
	@mkdir -p ./build/bin
	@cd ${MOCK_SERVER_PATH} && go build -o ../../build/bin/mock-server main.go


.PHONY: demo-server-docker
demo-server-docker:
	@echo "Building demo-server docker image"
	@docker build -f ./images/demo-server/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(DEMO_SERVER_IMAGE_NAME):$(DEMO_SERVER_IMAGE_TAG) \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) \
	--build-arg GO_TAGS=$(GO_TAGS) \
	--build-arg GOPROXY=$(GOPROXY) .


.PHONY: mock-server-docker
mock-server-docker:
	@echo "Building mock-server docker image"
	@docker build -f ./images/mock-server/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(MOCK_SERVER_IMAGE_NAME):$(MOCK_SERVER_IMAGE_TAG) \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) \
	--build-arg GO_TAGS=$(GO_TAGS) \
	--build-arg GOPROXY=$(GOPROXY) .


.PHONY: run-demo-server
run-demo-server:
	@echo "Starting demo server containers ..."
	@docker-compose -f deployment/demo-server/docker-compose.yml up --force-recreate -d
	@docker-compose -f deployment/openapi/docker-compose.yml up --force-recreate -d


.PHONY: run-mock-server
run-mock-server:
	@echo "Starting mock server containers ..."
	@docker-compose -f deployment/mock-server/docker-compose.yml up --force-recreate -d
	@docker-compose -f deployment/openapi/docker-compose.yml up --force-recreate -d


.PHONY: stop-mock-server
stop-mock-server:
	@echo "Stopping mock server containers ..."
	@docker-compose -f deployment/mock-server/docker-compose.yml down
	@docker-compose -f deployment/openapi/docker-compose.yml down


.PHONY: stop-demo-server
stop-demo-server:
	@echo "Stopping demo server containers ..."
	@docker-compose -f deployment/demo-server/docker-compose.yml down
	@docker-compose -f deployment/openapi/docker-compose.yml down