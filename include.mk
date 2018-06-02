GCP_PROJECT = hacker-playground
GCP_ZONE = europe-west3-b

KUBERNETES_CLUSTER = fort-perfect-cluster
KUBERNTES_CONTEXT = gke_$(GCP_PROJECT)_$(GCP_ZONE)_$(KUBERNETES_CLUSTER)

DOCKER_IMAGE = eu.gcr.io/$(GCP_PROJECT)/$(SERVICE_NAME)

dockerBuild:
	@echo "Build Docker image"
	docker build -t $(DOCKER_IMAGE) .

dockerPush: dockerBuild
	@echo "Build and push Docker images to GCP project"
	docker push $(DOCKER_IMAGE)

dockerRun: dockerBuild
	@echo "Build Docker image and run service inside the Docker container at http://localhost:8080"
	docker run -it --rm \
		--publish 8080:8080 \
		--name $(SERVICE_NAME) \
		$(DOCKER_IMAGE)

kubernetesApply: dockerPush
	@echo "Push Docker image and apply Kubernetes resources"
	kubectl --context $(KUBERNTES_CONTEXT) apply -f .