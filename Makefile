.PHONY: start
start:
	ctlptl apply -f ./ctlptl-cluster.yaml
	kubectl apply -f https://github.com/jetstack/cert-manager/releases/latest/download/cert-manager.yaml
	kubectl -n cert-manager wait --for=condition=available --timeout=180s --all deployments

.PHONY: install-mcing
install-mcing:
	kustomize build ./MCing | kubectl apply --server-side -f -

.PHONY: fetch-schema
fetch-schema:
	kubectl get --raw /openapi/v3/apis/mcing.kmdkuk.com/v1alpha1 > ./mcing-frontend/src/apis/openapi-v3-mcing.json

.PHONY: stop
stop:
	ctlptl delete -f ./ctlptl-cluster.yaml
