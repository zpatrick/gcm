package gcm

func NewTestMux(values map[string]interface{}) *Mux {
	providers := make(map[string]Provider, len(values))
	for k, v := range values {
		vcopy := v
		provider := ProviderFunc(func() (interface{}, error) {
			return vcopy, nil
		})

		providers[k] = provider
	}

	return &Mux{
		Providers: providers,
	}
}
