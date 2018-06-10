GCP_PROJECT = hacker-playground
GCP_ZONE = europe-west3-b

KUBERNETES_CLUSTER = ford-perfect-cluster
KUBERNTES_CONTEXT = gke_$(GCP_PROJECT)_$(GCP_ZONE)_$(KUBERNETES_CLUSTER)

# Version number of the service
SERVICE_VERSION = $(shell git rev-parse --short HEAD)

gcpContainerBuild:
	gcloud container --project $(GCP_PROJECT) builds submit \
		--config ../cloudbuild.yaml \
		--substitutions SHORT_SHA=$(SERVICE_VERSION),_COMPUTE_ZONE=$(GCP_ZONE),_CONTAINER_CLUSTER=$(KUBERNETES_CLUSTER) \
		../

gcpBuildStatus:
	gcloud container --project $(GCP_PROJECT) builds list --filter "tags=$(SERVICE_VERSION)"