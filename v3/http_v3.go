package v3

import (
	"log"
)

// HTTPMockProviderV3 is the entrypoint for V3 http consumer tests
type HTTPMockProviderV3 struct {
	*httpMockProvider
}

// NewV3Pact configures a new V3 HTTP Mock Provider for consumer tests
func NewV3Pact(config MockHTTPProviderConfig) (*HTTPMockProviderV3, error) {
	provider := &HTTPMockProviderV3{
		httpMockProvider: &httpMockProvider{
			config:               config,
			specificationVersion: V3,
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
func (p *HTTPMockProviderV3) AddInteraction() *InteractionV3 {
	log.Println("[DEBUG] pact add v3 interaction")
	interaction := p.httpMockProvider.mockserver.NewInteraction("")

	i := &InteractionV3{
		Interaction: Interaction{
			specificationVersion: V3,
			interaction:          interaction,
		},
	}

	p.httpMockProvider.v3Interactions = append(p.httpMockProvider.v3Interactions, i)
	return i
}
