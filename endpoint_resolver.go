package veneur

import (
	"math/rand"
	"net"
	"net/url"
)

// EndpointResolver is a probabilistic cache that resolves URLs to their IP
// addresses. It is thread-safe and holds a map of endpoints.
type EndpointResolver struct {
	CacheProbability float64
	lastResolvedURLs  sync.Map
}

// NewEndpointResolver returns a new EndpointResolver resolver with the
// probability determining how often it will cache.
func NewEndpointResolver(probability float64) *EndpointResolver {
	return &EndpointResolver{
		CacheProbability: probability,
    lastResolvedURLs: sync.Map{}
	}
}

// ResolveEndpoint attempts to resolve the url's host, and returns a new url whose
// host has been replaced by the first resolved address
// on failure, it returns the argument, and the resulting error
func (er *EndpointResolver) ResolveEndpoint(endpoint string) (string, error) {
	if er.lastResolvedURL != "" && rand.Float64() < er.CacheProbability {
		return er.lastResolvedURL, nil
	}

	origURL, err := url.Parse(endpoint)
	if err != nil {
		// caution: this error contains the endpoint itself, so if the endpoint
		// has secrets in it, you have to remove them
		return endpoint, err
	}

	origHost, origPort, err := net.SplitHostPort(origURL.Host)
	if err != nil {
		return endpoint, err
	}

	resolvedNames, err := net.LookupHost(origHost)
	if err != nil {
		return endpoint, err
	}
	if len(resolvedNames) == 0 {
		return endpoint, &net.DNSError{
			Err:  "no hosts found",
			Name: origHost,
		}
	}

	origURL.Host = net.JoinHostPort(resolvedNames[0], origPort)
	er.lastResolvedURL = origURL.String()
	return origURL.String(), nil
}
