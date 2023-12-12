IMAGE_NAME = tms
TAG = latest

.PHONY: usage

usage:
	@echo "Usage: make <command>"
	@echo "available commands: docker-build, docker-run, docker-deploy, clean"


docker-build:
	@echo "Building the docker image: $(IMAGE_NAME):$(TAG)"
	docker build -t $(IMAGE_NAME):$(TAG) .


docker-run:
	@echo "Running the docker image: $(IMAGE_NAME):$(TAG)"
	docker run \
		--rm \
		--name $(IMAGE_NAME)-dev \
		--detach \
		--env MDS_LISTEN=0.0.0.0:8088 \
		--env MDS_LOG_LEVEL=debug \
		--env MDS_CONFIG_DIR=./configs \
		--publish 8088:8088 $(IMAGE_NAME):$(TAG)


docker-deploy:
	@echo "Push to ghcr.io"
	docker push ghcr.io/goproslowyo/$(IMAGE_NAME):$(TAG)


clean:
	@echo "Cleaning up"
	docker stop $(IMAGE_NAME)-dev || true
	docker rm $(IMAGE_NAME)-dev || true
	docker rmi $(IMAGE_NAME):$(TAG) || true
