IMAGE_REPO = ghcr.io/goproslowyo
IMAGE_NAME = tms
TAG ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "dev")

.PHONY: usage

usage:
	@echo "Usage: make <command>"
	@echo "available commands: docker-build, docker-run, docker-deploy, clean"


docker-build:
	@echo "Building the docker image: $(IMAGE_NAME):$(TAG)"
	docker build -t $(IMAGE_REPO)/$(IMAGE_NAME):$(TAG) .


docker-run:
	@echo "Running the docker image: $(IMAGE_NAME):$(TAG)"
	docker run \
		--rm \
		--name $(IMAGE_NAME) \
		--detach \
		--env MDS_LISTEN=0.0.0.0:8088 \
		--env MDS_LOG_LEVEL=debug \
		--env MDS_CONFIG_DIR=./configs \
		--publish 8088:8088 $(IMAGE_REPO)/$(IMAGE_NAME):$(TAG)


docker-deploy:
	@echo "Push to ghcr.io"
	docker push $(IMAGE_REPO)/$(IMAGE_NAME):$(TAG)
	docker push $(IMAGE_REPO)/$(IMAGE_NAME):latest


docker-clean:
	@echo "Cleaning up"
	docker stop $(IMAGE_NAME) || true
	docker rm $(IMAGE_NAME) || true
	docker rmi $(IMAGE_NAME):$(TAG) || true
	docker rmi cgr.dev/chainguard/go:latest-dev || true
	docker rmi cgr.dev/chainguard/static:latest || true
