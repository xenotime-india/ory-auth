/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright 	2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package client_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	jose "gopkg.in/square/go-jose.v2"

	"github.com/ory/viper"

	. "github.com/ory/hydra/client"
	"github.com/ory/hydra/driver/configuration"
	"github.com/ory/hydra/internal"
)

func TestValidate(t *testing.T) {
	c := internal.NewConfigurationWithDefaults()
	viper.Set(configuration.ViperKeySubjectTypesSupported, []string{"pairwise", "public"})
	viper.Set(configuration.ViperKeyDefaultClientScope, []string{"openid"})

	v := NewValidator(c)
	for k, tc := range []struct {
		in        *Client
		check     func(t *testing.T, c *Client)
		expectErr bool
		v         func(t *testing.T) *Validator
	}{
		{
			in: new(Client),
			check: func(t *testing.T, c *Client) {
				assert.NotEmpty(t, c.ClientID)
				assert.NotEmpty(t, c.GetID())
				assert.Equal(t, c.GetID(), c.ClientID)
			},
		},
		{
			in: &Client{ClientID: "foo"},
			check: func(t *testing.T, c *Client) {
				assert.Equal(t, c.GetID(), c.ClientID)
			},
		},
		{
			in: &Client{ClientID: "foo"},
			check: func(t *testing.T, c *Client) {
				assert.Equal(t, c.GetID(), c.ClientID)
			},
		},
		{
			in:        &Client{ClientID: "foo", UserinfoSignedResponseAlg: "foo"},
			expectErr: true,
		},
		{
			in:        &Client{ClientID: "foo", TokenEndpointAuthMethod: "private_key_jwt"},
			expectErr: true,
		},
		{
			in:        &Client{ClientID: "foo", JSONWebKeys: &jose.JSONWebKeySet{}, JSONWebKeysURI: "asdf", TokenEndpointAuthMethod: "private_key_jwt"},
			expectErr: true,
		},
		{
			in:        &Client{ClientID: "foo", PostLogoutRedirectURIs: []string{"https://bar/"}, RedirectURIs: []string{"https://foo/"}},
			expectErr: true,
		},
		{
			in:        &Client{ClientID: "foo", PostLogoutRedirectURIs: []string{"http://foo/"}, RedirectURIs: []string{"https://foo/"}},
			expectErr: true,
		},
		{
			in:        &Client{ClientID: "foo", PostLogoutRedirectURIs: []string{"https://foo:1234/"}, RedirectURIs: []string{"https://foo/"}},
			expectErr: true,
		},
		{
			in: &Client{ClientID: "foo", PostLogoutRedirectURIs: []string{"https://foo/"}, RedirectURIs: []string{"https://foo/"}},
			check: func(t *testing.T, c *Client) {
				assert.Equal(t, []string{"https://foo/"}, c.PostLogoutRedirectURIs)
			},
		},
		{
			in: &Client{ClientID: "foo"},
			check: func(t *testing.T, c *Client) {
				assert.Equal(t, "public", c.SubjectType)
			},
		},
		{
			v: func(t *testing.T) *Validator {
				viper.Set(configuration.ViperKeySubjectTypesSupported, []string{"pairwise"})
				return NewValidator(c)
			},
			in: &Client{ClientID: "foo"},
			check: func(t *testing.T, c *Client) {
				assert.Equal(t, "pairwise", c.SubjectType)
			},
		},
		{
			in: &Client{ClientID: "foo", SubjectType: "pairwise"},
			check: func(t *testing.T, c *Client) {
				assert.Equal(t, "pairwise", c.SubjectType)
			},
		},
		{
			in:        &Client{ClientID: "foo", SubjectType: "foo"},
			expectErr: true,
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			if tc.v == nil {
				tc.v = func(t *testing.T) *Validator {
					return v
				}
			}
			err := tc.v(t).Validate(tc.in)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				tc.check(t, tc.in)
			}
		})
	}
}

func TestValidateSectorIdentifierURL(t *testing.T) {
	var payload string

	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(payload))
	}
	ts := httptest.NewTLSServer(h)
	defer ts.Close()

	v := NewValidatorWithClient(nil, ts.Client())

	for k, tc := range []struct {
		p         string
		r         []string
		u         string
		expectErr bool
	}{
		{
			u:         "",
			expectErr: true,
		},
		{
			u:         "http://foo/bar",
			expectErr: true,
		},
		{
			u:         ts.URL,
			expectErr: true,
		},
		{
			p:         `["http://foo"]`,
			u:         ts.URL,
			expectErr: false,
			r:         []string{"http://foo"},
		},
		{
			p:         `["http://foo"]`,
			u:         ts.URL,
			expectErr: true,
			r:         []string{"http://foo", "http://not-foo"},
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			payload = tc.p
			err := v.ValidateSectorIdentifierURL(tc.u, tc.r)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
