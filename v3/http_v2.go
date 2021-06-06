package v3

import (
	"log"
)

// HTTPMockProviderV2 is the entrypoint for V3 http consumer tests
type HTTPMockProviderV2 struct {
	*httpMockProvider
}

// TODO: this gets cumbersome per test.
// Make it parallelisable

// NewV2Pact configures a new V2 HTTP Mock Provider for consumer tests
func NewV2Pact(config MockHTTPProviderConfig) (*HTTPMockProviderV2, error) {
	provider := &HTTPMockProviderV2{
		httpMockProvider: &httpMockProvider{
			config:               config,
			specificationVersion: V2,
		},
	}
	err := provider.validateConfig()

	if err != nil {
		return nil, err
	}

	return provider, err
}

// AddInteraction creates a new Pact interaction, initialising all
// required things. Will automatically start a Mock Service if none running.
func (p *HTTPMockProviderV2) AddInteraction() *InteractionV2 {
	log.Println("[DEBUG] pact add v2 interaction")
	interaction := p.httpMockProvider.mockserver.NewInteraction("")

	i := &InteractionV2{
		Interaction: Interaction{
			specificationVersion: V2,
			interaction:          interaction,
		},
	}

	p.httpMockProvider.v2Interactions = append(p.httpMockProvider.v2Interactions, i)

	return i
}
