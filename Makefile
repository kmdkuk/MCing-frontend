.PHONY: start
start:
	ctlptl apply -f ./ctlptl-cluster.yaml

.PHONY: stop
stop:
	ctlptl delete -f ./ctlptl-cluster.yaml
