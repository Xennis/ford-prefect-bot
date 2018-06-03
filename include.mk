GCP_PROJECT = hacker-playground
GCP_ZONE = europe-west3-b

KUBERNETES_CLUSTER = ford-perfect-cluster
KUBERNTES_CONTEXT = gke_$(GCP_PROJECT)_$(GCP_ZONE)_$(KUBERNETES_CLUSTER)

DOCKER_IMAGE = eu.gcr.io/$(GCP_PROJECT)/$(SERVICE_NAME)

dockerBuild:
	@echo "Build Docker image"
	docker build -t $(DOCKER_IMAGE) .

dockerPush: dockerBuild
	@echo "Build and push Docker images to GCP project"
	docker push $(DOCKER_IMAGE)

kubernetesApply: dockerPush
	@echo "Push Docker image and apply Kubernetes resources"
	kubectl --context $(KUBERNTES_CONTEXT) apply -f .