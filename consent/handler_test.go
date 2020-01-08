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

package consent_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ory/hydra/x"

	"github.com/ory/hydra/internal"

	"github.com/stretchr/testify/require"

	"github.com/ory/hydra/client"
	. "github.com/ory/hydra/consent"
)

//func TestLogout(t *testing.T) {
//	conf := internal.NewConfigurationWithDefaults()
//	reg := internal.NewRegistry(conf)
//
//	r := x.NewRouterPublic()
//	h := NewHandler(reg, conf)
//
//	sid := uuid.New()
//
//	r.Handle("GET", "/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//		cookie, _ := reg.CookieStore().Get(r, CookieAuthenticationName)
//		require.NoError(t, reg.ConsentManager().CreateAuthenticationSession(context.TODO(), &AuthenticationSession{
//			ID:              sid,
//			Subject:         "foo",
//			AuthenticatedAt: time.Now(),
//		}))
//
//		cookie.Values[CookieAuthenticationSIDName] = sid
//		cookie.Options.MaxAge = 60
//
//		require.NoError(t, cookie.Save(r, w))
//	})
//
//	r.Handle("GET", "/logout", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	})
//
//	h.SetRoutes(r.RouterAdmin(), r)
//	ts := httptest.NewServer(r)
//	defer ts.Close()
//
//	viper.Set(configuration.ViperKeyLogoutRedirectURL, ts.URL+"/logout")
//
//	u, err := url.Parse(ts.URL)
//	require.NoError(t, err)
//
//	cj, err := cookiejar.New(new(cookiejar.Options))
//	require.NoError(t, err)
//
//	c := &http.Client{Jar: cj}
//	resp, err := c.Get(ts.URL + "/login")
//	require.NoError(t, err)
//	require.EqualValues(t, http.StatusOK, resp.StatusCode)
//	require.Len(t, cj.Cookies(u), 1)
//
//	resp, err = c.Get(ts.URL + "/oauth2/auth/sessions/login/revoke")
//	require.NoError(t, err)
//	require.EqualValues(t, http.StatusOK, resp.StatusCode)
//	assert.Len(t, cj.Cookies(u), 0)
//	assert.EqualValues(t, ts.URL+"/logout", resp.Request.URL.String())
//}

func TestGetLogoutRequest(t *testing.T) {
	for k, tc := range []struct {
		exists bool
		used   bool
		status int
	}{
		{false, false, http.StatusNotFound},
		{true, false, http.StatusOK},
		{true, true, http.StatusConflict},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			key := fmt.Sprint(k)
			challenge := "challenge" + key

			conf := internal.NewConfigurationWithDefaults()
			reg := internal.NewRegistry(conf)

			if tc.exists {
				require.NoError(t, reg.ConsentManager().CreateLogoutRequest(context.TODO(), &LogoutRequest{
					Client:    &client.Client{ClientID: "client" + key},
					Challenge: challenge,
					WasUsed:   tc.used,
				}))
			}

			h := NewHandler(reg, conf)
			r := x.NewRouterAdmin()
			h.SetRoutes(r)
			ts := httptest.NewServer(r)
			defer ts.Close()

			c := &http.Client{}
			resp, err := c.Get(ts.URL + LogoutPath + "?challenge=" + challenge)
			require.NoError(t, err)
			require.EqualValues(t, tc.status, resp.StatusCode)
		})
	}
}

func TestGetLoginRequest(t *testing.T) {
	for k, tc := range []struct {
		exists  bool
		handled bool
		status  int
	}{
		{false, false, http.StatusNotFound},
		{true, false, http.StatusOK},
		{true, true, http.StatusConflict},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			key := fmt.Sprint(k)
			challenge := "challenge" + key

			conf := internal.NewConfigurationWithDefaults()
			reg := internal.NewRegistry(conf)

			if tc.exists {
				require.NoError(t, reg.ConsentManager().CreateLoginRequest(context.TODO(), &LoginRequest{
					Client:     &client.Client{ClientID: "client" + key},
					Challenge:  challenge,
					WasHandled: tc.handled,
				}))
			}

			h := NewHandler(reg, conf)
			r := x.NewRouterAdmin()
			h.SetRoutes(r)
			ts := httptest.NewServer(r)
			defer ts.Close()

			c := &http.Client{}
			resp, err := c.Get(ts.URL + LoginPath + "?challenge=" + challenge)
			require.NoError(t, err)
			require.EqualValues(t, tc.status, resp.StatusCode)
		})
	}
}

func TestGetConsentRequest(t *testing.T) {
	for k, tc := range []struct {
		exists  bool
		handled bool
		status  int
	}{
		{false, false, http.StatusNotFound},
		{true, false, http.StatusOK},
		{true, true, http.StatusConflict},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			key := fmt.Sprint(k)
			challenge := "challenge" + key

			conf := internal.NewConfigurationWithDefaults()
			reg := internal.NewRegistry(conf)

			if tc.exists {
				require.NoError(t, reg.ConsentManager().CreateConsentRequest(context.TODO(), &ConsentRequest{
					Client:     &client.Client{ClientID: "client" + key},
					Challenge:  challenge,
					WasHandled: tc.handled,
				}))
			}

			h := NewHandler(reg, conf)

			r := x.NewRouterAdmin()
			h.SetRoutes(r)
			ts := httptest.NewServer(r)
			defer ts.Close()

			c := &http.Client{}
			resp, err := c.Get(ts.URL + ConsentPath + "?challenge=" + challenge)
			require.NoError(t, err)
			require.EqualValues(t, tc.status, resp.StatusCode)
		})
	}
}
