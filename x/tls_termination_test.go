package x_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/viper"

	"github.com/ory/hydra/driver/configuration"
	"github.com/ory/hydra/internal"
	. "github.com/ory/hydra/x"
)

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("should not have been called")
}

func noopHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func TestDoesRequestSatisfyTermination(t *testing.T) {
	c := internal.NewConfigurationWithDefaultsAndHTTPS()
	r := internal.NewRegistry(c)

	t.Run("case=tls-termination-disabled", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{Header: http.Header{}, URL: new(url.URL)}, panicHandler)
		assert.EqualValues(t, http.StatusBadGateway, res.Code)
	})

	t.Run("case=missing-x-forwarded-proto", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{Header: http.Header{}, URL: new(url.URL)}, panicHandler)
		assert.EqualValues(t, http.StatusBadGateway, res.Code)
	})

	t.Run("case=x-forwarded-proto-is-http", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{Header: http.Header{"X-Forwarded-Proto": []string{"http"}}, URL: new(url.URL)}, panicHandler)
		assert.EqualValues(t, http.StatusBadGateway, res.Code)
	})

	t.Run("case=missing-x-forwarded-for", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{Header: http.Header{"X-Forwarded-Proto": []string{"https"}}, URL: new(url.URL)}, panicHandler)
		assert.EqualValues(t, http.StatusBadGateway, res.Code)
	})

	t.Run("case=remote-not-in-cidr", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "227.0.0.1:123",
			Header:     http.Header{"X-Forwarded-Proto": []string{"https"}}, URL: new(url.URL)},
			panicHandler,
		)
		assert.EqualValues(t, http.StatusBadGateway, res.Code)
	})

	t.Run("case=remote-and-forwarded-not-in-cidr", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "227.0.0.1:123",
			Header: http.Header{
				"X-Forwarded-Proto": []string{"https"},
				"X-Forwarded-For":   []string{"227.0.0.1"},
			}, URL: new(url.URL)},
			panicHandler,
		)
		assert.EqualValues(t, http.StatusBadGateway, res.Code)
	})

	t.Run("case=remote-matches-cidr", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "127.0.0.1:123",
			Header: http.Header{
				"X-Forwarded-Proto": []string{"https"},
			}, URL: new(url.URL)},
			noopHandler,
		)
		assert.EqualValues(t, http.StatusNoContent, res.Code)
	})

	t.Run("case=passes-because-health-alive-endpoint", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "227.0.0.1:123",
			Header: http.Header{
				"X-Forwarded-Proto": []string{"https"},
			},
			URL: &url.URL{Path: "/health/alive"},
		},
			noopHandler,
		)
		assert.EqualValues(t, http.StatusNoContent, res.Code)
	})

	t.Run("case=passes-because-health-ready-endpoint", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "227.0.0.1:123",
			Header: http.Header{
				"X-Forwarded-Proto": []string{"https"},
			},
			URL: &url.URL{Path: "/health/alive"},
		},
			noopHandler,
		)
		assert.EqualValues(t, http.StatusNoContent, res.Code)
	})

	t.Run("case=forwarded-matches-cidr", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "227.0.0.2:123",
			Header: http.Header{
				"X-Forwarded-For":   []string{"227.0.0.1, 127.0.0.1, 227.0.0.2"},
				"X-Forwarded-Proto": []string{"https"},
			}, URL: new(url.URL)},
			noopHandler,
		)
		assert.EqualValues(t, http.StatusNoContent, res.Code)
	})

	t.Run("case=forwarded-matches-cidr-without-spaces", func(t *testing.T) {
		viper.Set(configuration.ViperKeyAllowTLSTerminationFrom, "126.0.0.1/24,127.0.0.1/24")

		res := httptest.NewRecorder()
		RejectInsecureRequests(r, c)(res, &http.Request{
			RemoteAddr: "227.0.0.2:123",
			Header: http.Header{
				"X-Forwarded-For":   []string{"227.0.0.1,127.0.0.1,227.0.0.2"},
				"X-Forwarded-Proto": []string{"https"},
			}, URL: new(url.URL)},
			noopHandler,
		)
		assert.EqualValues(t, http.StatusNoContent, res.Code)
	})
}
