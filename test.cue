cluster: {
	name:               "aks-dev-marc-test-123"
	kubernetes_version: "1.27.3"
	owner: "MarcBuch"

	default_pool: {
		min_count: 2
		max_count: 6
		sku:       "Standard_D2_v2"
	}
}

network: {
    name: "vnet-dev-marc-test-123"
	cidr: "10.0.0.0/16"
}

features: {
	external_secrets: true
	grafana:          true
	snyk:             true
	argo:             false
	flux:             false
}
