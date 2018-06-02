include ../include.mk


createCluster:
	@echo "Create a Kubernetes cluster"
	gcloud container --project $(GCP_PROJECT) clusters create $(KUBERNETES_CLUSTER) \
		--zone $(GCP_ZONE) \
		--machine-type "g1-small" \
		--disk-size "25" \
		--num-nodes "2" \
		--preemptible 

getCredentials:
	@echo "Get credentials of the cluster"
	gcloud container clusters get-credentials $(KUBERNETES_CLUSTER)