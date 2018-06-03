SERVICE_NAME = ford-perfect

include ../include.mk

KUBERNETES_NGINX_SECRET = nginx-ssl
CRT_FILE = /tmp/nginx.crt

kubernetesCreateSecrets:
	@echo "Generate self-signed key and certificate"
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
		-keyout /tmp/nginx.key \
		-out $(CRT_FILE)
	@echo "Create a Kubernetes secret with SSL key and certificate"
	kubectl --context $(KUBERNTES_CONTEXT) create secret generic $(KUBERNETES_NGINX_SECRET) \
		--from-file=/tmp/nginx.key \
		--from-file=$(CRT_FILE)

kubernetesDeleteSecrets:
	kubectl --context $(KUBERNTES_CONTEXT) delete secret $(KUBERNETES_NGINX_SECRET)

telegramSetWebhook:
	export SERVICE_IP=$(kubectl get services ford-perfect-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')
	curl -F "url=https://$(SERVICE_IP)/updatesHook?key=$(GCP_ENDPOINTS_KEY)" -F "certificate=@$(CRT_FILE)" https://api.telegram.org/bot$(TELEGRAM_API_TOKEN)/setWebhook

telegramDeleteWebhook:
	curl https://api.telegram.org/bot$(TELEGRAM_API_TOKEN)/deleteWebhook
