include ../../include.mk

create: createCluster getCredentials

createCluster:
	@echo "Create a Kubernetes cluster"
	gcloud container --project $(GCP_PROJECT) clusters create $(KUBERNETES_CLUSTER) \
		--zone $(GCP_ZONE) \
		--machine-type "g1-small" \
		--disk-size "25" \
		--num-nodes "2" \
		--preemptible 

deleteCluster:
	@echo "Delete the Kubernetes cluster"
	gcloud container --project $(GCP_PROJECT) clusters delete fort-perfect-cluster

getCredentials:
	@echo "Get credentials of the cluster"
	gcloud container --project $(GCP_PROJECT)  clusters get-credentials $(KUBERNETES_CLUSTER)

deleteCredentials:
	@echo "Delete kubectl context"
	kubectl config delete-context $(KUBERNTES_CONTEXT)
