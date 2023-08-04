build:
	IMAGE_TAG=$(date "+%Y%m%d%H%M%S")
	./docker.sh -t $(IMAGE_TAG)


.PHONY: deploy

deploy:
	IMAGE_TAG=$(shell date "+%Y%m%d%H%M%S"); \
	./docker.sh -t $${IMAGE_TAG}; \
	./deploy.sh -t $${IMAGE_TAG}

run:
	./output/bootstrap.sh