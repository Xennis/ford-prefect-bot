GCP_PROJECT = hacker-playground
GCP_ZONE = europe-west3-b

KUBERNETES_CLUSTER = ford-perfect-cluster
KUBERNTES_CONTEXT = gke_$(GCP_PROJECT)_$(GCP_ZONE)_$(KUBERNETES_CLUSTER)

DOCKER_IMAGE = eu.gcr.io/$(GCP_PROJECT)/$(SERVICE_NAME)
# Version number of the service
SERVICE_VERSION = $(shell git rev-parse --short HEAD)
# Distribution directory
DIST_DIR = dist

goBuild:
	@echo "Build Go binary"
	# Envs:
	# * CGO_ENABLED=0 Disables CGO tools. For cross-compling it is disabled by
	#	default. CG0 enables the creation of Go packages that call C code.
	#
	# Build parameters:
	# * -a Force rebuild of packages that are already up-to-date.
	# * -o: Write executable
	CGO_ENABLED=0 GOOS=linux go build -a -o $(DIST_DIR)/service .

dockerBuild: goBuild
	@echo "Copy Dockerfile and required files into Docker context"
	cp ../infrastructure/dockerfile/* $(DIST_DIR)/
	@echo "Build Docker image"
	docker build \
		-t $(DOCKER_IMAGE):$(SERVICE_VERSION) \
		-t $(DOCKER_IMAGE):latest \
		$(DIST_DIR)/

dockerPush: dockerBuild
	@echo "Build and push Docker images to GCP project"
	docker push $(DOCKER_IMAGE)
	docker push $(DOCKER_IMAGE):$(SERVICE_VERSION)

kubernetesApply: dockerPush
	@echo "Push Docker image and apply Kubernetes resources"
	kubectl --context $(KUBERNTES_CONTEXT) apply -f .