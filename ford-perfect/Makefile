SERVICE_NAME = ford-perfect

include ../include.mk

kubernetesCreateSecrets:
	@echo "Generate self-signed key and certificate"
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
		-keyout /tmp/nginx.key \
		-out /tmp/nginx.crt
	@echo "Create a Kubernetes secret with SSL key and certificate"
	kubectl --context $(KUBERNTES_CONTEXT) create secret generic nginx-ssl \
		--from-file=/tmp/nginx.crt \
		--from-file=/tmp/nginx.key
