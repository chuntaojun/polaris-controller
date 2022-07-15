.PHONY: build

REGISTRY = ""
REPO = ccr.ccs.tencentyun.com/chuntaojun/polaris-controller
SIDECAR_INIT_REPO = ccr.ccs.tencentyun.com/chuntaojun/polaris-sidecar-init
IMAGE_TAG = v1.2.4

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o ./bin/polaris-controller ./cmd/polaris-controller/main.go

build-image:
	docker build . -f ./docker/Dockerfile -t $(REGISTRY)$(REPO):$(IMAGE_TAG)

build-sidecar-init:
	docker build ./sidecar/polaris-sidecar-init -f ./sidecar/polaris-sidecar-init/Dockerfile -t $(REGISTRY)$(SIDECAR_INIT_REPO):$(IMAGE_TAG)

push-image-withlogin: build build-image build-sidecar-init login push-image

push-image: build build-image build-sidecar-init
	docker push $(REGISTRY)$(REPO):$(IMAGE_TAG)
	docker tag $(REGISTRY)$(REPO):$(IMAGE_TAG) $(REGISTRY)$(REPO):latest
	docker push $(REGISTRY)$(REPO):latest

	docker push $(REGISTRY)$(SIDECAR_INIT_REPO):$(IMAGE_TAG)
	docker tag $(REGISTRY)$(SIDECAR_INIT_REPO):$(IMAGE_TAG) $(REGISTRY)$(SIDECAR_INIT_REPO):latest
	docker push $(REGISTRY)$(SIDECAR_INIT_REPO):latest

login:
	@docker login --username=$(DOCKER_USER) --password=$(DOCKER_PASS) $(REGISTRY)