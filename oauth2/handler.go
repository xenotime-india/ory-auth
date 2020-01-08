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
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package oauth2

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/token/jwt"
	"github.com/ory/hydra/client"
	"github.com/ory/hydra/consent"
	"github.com/ory/hydra/driver/configuration"
	"github.com/ory/hydra/x"
	"github.com/ory/x/urlx"
)

const (
	DefaultLoginPath      = "/oauth2/fallbacks/login"
	DefaultConsentPath    = "/oauth2/fallbacks/consent"
	DefaultPostLogoutPath = "/oauth2/fallbacks/logout/callback"
	DefaultLogoutPath     = "/oauth2/fallbacks/logout"
	DefaultErrorPath      = "/oauth2/fallbacks/error"
	TokenPath             = "/oauth2/token"
	AuthPath              = "/oauth2/auth"
	LogoutPath            = "/oauth2/sessions/logout"

	UserinfoPath  = "/userinfo"
	WellKnownPath = "/.well-known/openid-configuration"
	JWKPath       = "/.well-known/jwks.json"

	// IntrospectPath points to the OAuth2 introspection endpoint.
	IntrospectPath = "/oauth2/introspect"
	RevocationPath = "/oauth2/revoke"
	FlushPath      = "/oauth2/flush"
)

type Handler struct {
	r InternalRegistry
	c Configuration
}

func NewHandler(r InternalRegistry, c Configuration) *Handler {
	return &Handler{r: r, c: c}
}

func (h *Handler) SetRoutes(admin *x.RouterAdmin, public *x.RouterPublic, corsMiddleware func(http.Handler) http.Handler) {
	public.Handler("OPTIONS", TokenPath, corsMiddleware(http.HandlerFunc(h.handleOptions)))
	public.Handler("POST", TokenPath, corsMiddleware(http.HandlerFunc(h.TokenHandler)))

	public.GET(AuthPath, h.AuthHandler)
	public.POST(AuthPath, h.AuthHandler)
	public.GET(LogoutPath, h.LogoutHandler)

	public.GET(DefaultLoginPath, h.fallbackHandler("", "", http.StatusInternalServerError, configuration.ViperKeyLoginURL))
	public.GET(DefaultConsentPath, h.fallbackHandler("", "", http.StatusInternalServerError, configuration.ViperKeyConsentURL))
	public.GET(DefaultLogoutPath, h.fallbackHandler("", "", http.StatusInternalServerError, configuration.ViperKeyLogoutURL))
	public.GET(DefaultPostLogoutPath, h.fallbackHandler(
		"You logged out successfully!",
		"The Default Post Logout URL is not set which is why you are seeing this fallback page. Your log out request however succeeded.",
		http.StatusOK,
		configuration.ViperKeyLogoutRedirectURL,
	))
	public.GET(DefaultErrorPath, h.DefaultErrorHandler)

	public.Handler("OPTIONS", RevocationPath, corsMiddleware(http.HandlerFunc(h.handleOptions)))
	public.Handler("POST", RevocationPath, corsMiddleware(http.HandlerFunc(h.RevocationHandler)))
	public.Handler("OPTIONS", WellKnownPath, corsMiddleware(http.HandlerFunc(h.handleOptions)))
	public.Handler("GET", WellKnownPath, corsMiddleware(http.HandlerFunc(h.WellKnownHandler)))
	public.Handler("OPTIONS", UserinfoPath, corsMiddleware(http.HandlerFunc(h.handleOptions)))
	public.Handler("GET", UserinfoPath, corsMiddleware(http.HandlerFunc(h.UserinfoHandler)))
	public.Handler("POST", UserinfoPath, corsMiddleware(http.HandlerFunc(h.UserinfoHandler)))

	admin.POST(IntrospectPath, h.IntrospectHandler)
	admin.POST(FlushPath, h.FlushHandler)
}

// swagger:route GET /oauth2/sessions/logout public disconnectUser
//
// OpenID Connect Front-Backchannel enabled Logout
//
// This endpoint initiates and completes user logout at ORY Hydra and initiates OpenID Connect Front-/Back-channel logout:
//
// - https://openid.net/specs/openid-connect-frontchannel-1_0.html
// - https://openid.net/specs/openid-connect-backchannel-1_0.html
//
//     Schemes: http, https
//
//     Responses:
//       302: emptyResponse
func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handled, err := h.r.ConsentStrategy().HandleOpenIDConnectLogout(w, r)

	if errors.Cause(err) == consent.ErrAbortOAuth2Request {
		return
	} else if err != nil {
		x.LogError(err, h.r.Logger())
		h.forwardError(w, r, err)
		return
	}

	if len(handled.FrontChannelLogoutURLs) == 0 {
		http.Redirect(w, r, handled.RedirectTo, http.StatusFound)
		return
	}

	// TODO How are we supposed to test this? Maybe with cypress? #1368
	t, err := template.New("logout").Parse(`<html>
<head>
    <meta http-equiv="refresh" content="7; URL={{ .RedirectTo }}">
</head>
<style type="text/css">
    iframe { position: absolute; left: 0; top: 0; height: 0; width: 0; border: none; }
</style>
<script>
    var total = {{ len .FrontChannelLogoutURLs }};
    var redir = {{ .RedirectTo }};

	function redirect() {
		window.location.replace(redir);

		// In case replace failed try href
		setTimeout(function () {
			window.location.href = redir;
		}, 250); // Show message after http-equiv="refresh"
	}

    function done() {
        total--;
        if (total < 1) {
			setTimeout(redirect, 500);
        }
    }

	setTimeout(redirect, 7000); // redirect after 5 seconds if e.g. an iframe doesn't load

	// If the redirect takes unusually long, show a message
	setTimeout(function () {
		document.getElementById("redir").style.display = "block";
	}, 2000);
</script>
<body>
<noscript>
    <p>
        JavaScript is disabled - you should be redirected in 5 seconds but if not, click <a
            href="{{ .RedirectTo }}">here</a> to continue.
    </p>
</noscript>

<p id="redir" style="display: none">
    Redirection takes unusually long. If you are not being redirected within the next seconds, click <a href="{{ .RedirectTo }}">here</a> to continue.
</p>

{{ range .FrontChannelLogoutURLs }}<iframe src="{{ . }}" onload="done(this)"></iframe>
{{ end }}
</body>
</html>`)
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.forwardError(w, r, err)
		return
	}

	if err := t.Execute(w, handled); err != nil {
		x.LogError(err, h.r.Logger())
		h.forwardError(w, r, err)
		return
	}
}

// swagger:route GET /.well-known/openid-configuration public discoverOpenIDConfiguration
//
// OpenID Connect Discovery
//
// The well known endpoint an be used to retrieve information for OpenID Connect clients. We encourage you to not roll
// your own OpenID Connect client but to use an OpenID Connect client library instead. You can learn more on this
// flow at https://openid.net/specs/openid-connect-discovery-1_0.html .
//
// Popular libraries for OpenID Connect clients include oidc-client-js (JavaScript), go-oidc (Golang), and others.
// For a full list of clients go here: https://openid.net/developers/certified/
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: wellKnown
//       401: genericError
//       500: genericError
func (h *Handler) WellKnownHandler(w http.ResponseWriter, r *http.Request) {
	h.r.Writer().Write(w, r, &WellKnown{
		Issuer:                             strings.TrimRight(h.c.IssuerURL().String(), "/") + "/",
		AuthURL:                            urlx.AppendPaths(h.c.IssuerURL(), AuthPath).String(),
		TokenURL:                           urlx.AppendPaths(h.c.IssuerURL(), TokenPath).String(),
		JWKsURI:                            urlx.AppendPaths(h.c.IssuerURL(), JWKPath).String(),
		RevocationEndpoint:                 urlx.AppendPaths(h.c.IssuerURL(), RevocationPath).String(),
		RegistrationEndpoint:               h.c.OAuth2ClientRegistrationURL().String(),
		SubjectTypes:                       h.c.SubjectTypesSupported(),
		ResponseTypes:                      []string{"code", "code id_token", "id_token", "token id_token", "token", "token id_token code"},
		ClaimsSupported:                    h.c.OIDCDiscoverySupportedClaims(),
		ScopesSupported:                    h.c.OIDCDiscoverySupportedScope(),
		UserinfoEndpoint:                   h.c.OIDCDiscoveryUserinfoEndpoint(),
		TokenEndpointAuthMethodsSupported:  []string{"client_secret_post", "client_secret_basic", "private_key_jwt", "none"},
		IDTokenSigningAlgValuesSupported:   []string{"RS256"},
		GrantTypesSupported:                []string{"authorization_code", "implicit", "client_credentials", "refresh_token"},
		ResponseModesSupported:             []string{"query", "fragment"},
		UserinfoSigningAlgValuesSupported:  []string{"none", "RS256"},
		RequestParameterSupported:          true,
		RequestURIParameterSupported:       true,
		RequireRequestURIRegistration:      true,
		BackChannelLogoutSupported:         true,
		BackChannelLogoutSessionSupported:  true,
		FrontChannelLogoutSupported:        true,
		FrontChannelLogoutSessionSupported: true,
		EndSessionEndpoint:                 urlx.AppendPaths(h.c.IssuerURL(), LogoutPath).String(),
	})
}

// swagger:route GET /userinfo public userinfo
//
// OpenID Connect Userinfo
//
// This endpoint returns the payload of the ID Token, including the idTokenExtra values, of
// the provided OAuth 2.0 Access Token.
//
// For more information please [refer to the spec](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo).
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//       oauth2:
//
//     Responses:
//       200: userinfoResponse
//       401: genericError
//       500: genericError
func (h *Handler) UserinfoHandler(w http.ResponseWriter, r *http.Request) {
	session := NewSession("")
	tokenType, ar, err := h.r.OAuth2Provider().IntrospectToken(r.Context(), fosite.AccessTokenFromRequest(r), fosite.AccessToken, session)
	if err != nil {
		h.r.Writer().WriteError(w, r, err)
		return
	}

	if tokenType != fosite.AccessToken {
		h.r.Writer().WriteErrorCode(w, r, http.StatusUnauthorized, errors.New("Only access tokens are allowed in the authorization header"))
		return
	}

	c, ok := ar.GetClient().(*client.Client)
	if !ok {
		h.r.Writer().WriteError(w, r, errors.WithStack(fosite.ErrServerError.WithHint("Unable to type assert to *client.Client")))
		return
	}

	if c.UserinfoSignedResponseAlg == "RS256" {
		interim := ar.GetSession().(*Session).IDTokenClaims().ToMap()

		delete(interim, "nonce")
		delete(interim, "at_hash")
		delete(interim, "c_hash")
		delete(interim, "auth_time")
		delete(interim, "iat")
		delete(interim, "rat")
		delete(interim, "exp")
		delete(interim, "jti")

		keyID, err := h.r.OpenIDJWTStrategy().GetPublicKeyID(r.Context())
		if err != nil {
			h.r.Writer().WriteError(w, r, err)
			return
		}

		token, _, err := h.r.OpenIDJWTStrategy().Generate(r.Context(), jwt2.MapClaims(interim), &jwt.Headers{
			Extra: map[string]interface{}{
				"kid": keyID,
			},
		})
		if err != nil {
			h.r.Writer().WriteError(w, r, err)
			return
		}

		w.Header().Set("Content-Type", "application/jwt")
		w.Write([]byte(token))
	} else if c.UserinfoSignedResponseAlg == "" || c.UserinfoSignedResponseAlg == "none" {
		interim := ar.GetSession().(*Session).IDTokenClaims().ToMap()
		delete(interim, "aud")
		delete(interim, "iss")
		delete(interim, "nonce")
		delete(interim, "at_hash")
		delete(interim, "c_hash")
		delete(interim, "auth_time")
		delete(interim, "iat")
		delete(interim, "rat")
		delete(interim, "exp")
		delete(interim, "jti")

		h.r.Writer().Write(w, r, interim)
	} else {
		h.r.Writer().WriteError(w, r, errors.WithStack(fosite.ErrServerError.WithHint(fmt.Sprintf("Unsupported userinfo signing algorithm \"%s\"", c.UserinfoSignedResponseAlg))))
		return
	}
}

// swagger:route POST /oauth2/revoke public revokeOAuth2Token
//
// Revoke OAuth2 tokens
//
// Revoking a token (both access and refresh) means that the tokens will be invalid. A revoked access token can no
// longer be used to make access requests, and a revoked refresh token can no longer be used to refresh an access token.
// Revoking a refresh token also invalidates the access token that was created with it. A token may only be revoked by
// the client the token was generated for.
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Schemes: http, https
//
//     Security:
//       basic:
//       oauth2:
//
//     Responses:
//       200: emptyResponse
//       401: genericError
//       500: genericError
func (h *Handler) RevocationHandler(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()

	err := h.r.OAuth2Provider().NewRevocationRequest(ctx, r)
	if err != nil {
		x.LogError(err, h.r.Logger())
	}

	h.r.OAuth2Provider().WriteRevocationResponse(w, err)
}

// swagger:route POST /oauth2/introspect admin introspectOAuth2Token
//
// Introspect OAuth2 tokens
//
// The introspection endpoint allows to check if a token (both refresh and access) is active or not. An active token
// is neither expired nor revoked. If a token is active, additional information on the token will be included. You can
// set additional data for a token by setting `accessTokenExtra` during the consent flow.
//
// For more information [read this blog post](https://www.oauth.com/oauth2-servers/token-introspection-endpoint/).
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//       basic:
//       oauth2:
//
//     Responses:
//       200: oAuth2TokenIntrospection
//       401: genericError
//       500: genericError
func (h *Handler) IntrospectHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var session = NewSession("")
	var ctx = r.Context()

	if r.Method != "POST" {
		err := errors.WithStack(fosite.ErrInvalidRequest.WithHintf("HTTP method is \"%s\", expected \"POST\".", r.Method))
		x.LogError(err, h.r.Logger())
		h.r.OAuth2Provider().WriteIntrospectionError(w, err)
		return
	} else if err := r.ParseMultipartForm(1 << 20); err != nil && err != http.ErrNotMultipart {
		err := errors.WithStack(fosite.ErrInvalidRequest.WithHint("Unable to parse HTTP body, make sure to send a properly formatted form request body.").WithDebug(err.Error()))
		x.LogError(err, h.r.Logger())
		h.r.OAuth2Provider().WriteIntrospectionError(w, err)
		return
	} else if len(r.PostForm) == 0 {
		err := errors.WithStack(fosite.ErrInvalidRequest.WithHint("The POST body can not be empty."))
		x.LogError(err, h.r.Logger())
		h.r.OAuth2Provider().WriteIntrospectionError(w, err)
		return
	}

	token := r.PostForm.Get("token")
	tokenType := r.PostForm.Get("token_type_hint")
	scope := r.PostForm.Get("scope")

	tt, ar, err := h.r.OAuth2Provider().IntrospectToken(ctx, token, fosite.TokenType(tokenType), session, strings.Split(scope, " ")...)
	if err != nil {
		x.LogError(err, h.r.Logger())
		err := errors.WithStack(fosite.ErrInactiveToken.WithHint("An introspection strategy indicated that the token is inactive.").WithDebug(err.Error()))
		h.r.OAuth2Provider().WriteIntrospectionError(w, err)
		return
	}

	resp := &fosite.IntrospectionResponse{
		Active:          true,
		AccessRequester: ar,
		TokenType:       tt,
	}

	exp := resp.GetAccessRequester().GetSession().GetExpiresAt(tt)
	if exp.IsZero() {
		if tt == fosite.RefreshToken {
			exp = resp.GetAccessRequester().GetRequestedAt().Add(h.c.RefreshTokenLifespan())
		} else {
			exp = resp.GetAccessRequester().GetRequestedAt().Add(h.c.AccessTokenLifespan())
		}
	}

	session, ok := resp.GetAccessRequester().GetSession().(*Session)
	if !ok {
		err := errors.WithStack(fosite.ErrServerError.WithHint("Expected session to be of type *Session, but got another type.").WithDebug(fmt.Sprintf("Got type %s", reflect.TypeOf(resp.GetAccessRequester().GetSession()))))
		x.LogError(err, h.r.Logger())
		h.r.OAuth2Provider().WriteIntrospectionError(w, err)
		return
	}

	var obfuscated string
	if len(session.Claims.Subject) > 0 && session.Claims.Subject != session.Subject {
		obfuscated = session.Claims.Subject
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err = json.NewEncoder(w).Encode(&Introspection{
		Active:            resp.IsActive(),
		ClientID:          resp.GetAccessRequester().GetClient().GetID(),
		Scope:             strings.Join(resp.GetAccessRequester().GetGrantedScopes(), " "),
		ExpiresAt:         exp.Unix(),
		IssuedAt:          resp.GetAccessRequester().GetRequestedAt().Unix(),
		Subject:           session.GetSubject(),
		Username:          session.GetUsername(),
		Extra:             session.Extra,
		Audience:          resp.GetAccessRequester().GetGrantedAudience(),
		Issuer:            strings.TrimRight(h.c.IssuerURL().String(), "/") + "/",
		ObfuscatedSubject: obfuscated,
		TokenType:         string(resp.GetTokenType()),
	}); err != nil {
		x.LogError(errors.WithStack(err), h.r.Logger())
	}
}

// swagger:route POST /oauth2/flush admin flushInactiveOAuth2Tokens
//
// Flush Expired OAuth2 Access Tokens
//
// This endpoint flushes expired OAuth2 access tokens from the database. You can set a time after which no tokens will be
// not be touched, in case you want to keep recent tokens for auditing. Refresh tokens can not be flushed as they are deleted
// automatically when performing the refresh flow.
//
//     Consumes:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       204: emptyResponse
//       401: genericError
//       500: genericError
func (h *Handler) FlushHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var fr FlushInactiveOAuth2TokensRequest
	if err := json.NewDecoder(r.Body).Decode(&fr); err != nil {
		h.r.Writer().WriteError(w, r, err)
		return
	}

	if fr.NotAfter.IsZero() {
		fr.NotAfter = time.Now()
	}

	if err := h.r.OAuth2Storage().FlushInactiveAccessTokens(r.Context(), fr.NotAfter); err != nil {
		h.r.Writer().WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /oauth2/token public oauth2Token
//
// The OAuth 2.0 token endpoint
//
// The client makes a request to the token endpoint by sending the
// following parameters using the "application/x-www-form-urlencoded" HTTP
// request entity-body.
//
// > Do not implement a client for this endpoint yourself. Use a library. There are many libraries
// > available for any programming language. You can find a list of libraries here: https://oauth.net/code/
// >
// > Do note that Hydra SDK does not implement this endpoint properly. Use one of the libraries listed above!
//
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Security:
//       basic:
//       oauth2:
//
//     Responses:
//       200: oauth2TokenResponse
//       401: genericError
//       500: genericError
func (h *Handler) TokenHandler(w http.ResponseWriter, r *http.Request) {
	var session = NewSession("")
	var ctx = r.Context()

	accessRequest, err := h.r.OAuth2Provider().NewAccessRequest(ctx, r, session)
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.r.OAuth2Provider().WriteAccessError(w, accessRequest, err)
		return
	}

	if accessRequest.GetGrantTypes().Exact("client_credentials") {
		var accessTokenKeyID string
		if h.c.AccessTokenStrategy() == "jwt" {
			accessTokenKeyID, err = h.r.AccessTokenJWTStrategy().GetPublicKeyID(r.Context())
			if err != nil {
				x.LogError(err, h.r.Logger())
				h.r.OAuth2Provider().WriteAccessError(w, accessRequest, err)
				return
			}
		}

		session.Subject = accessRequest.GetClient().GetID()
		session.ClientID = accessRequest.GetClient().GetID()
		session.KID = accessTokenKeyID
		session.DefaultSession.Claims.Issuer = strings.TrimRight(h.c.IssuerURL().String(), "/") + "/"
		session.DefaultSession.Claims.IssuedAt = time.Now().UTC()

		for _, scope := range accessRequest.GetRequestedScopes() {
			if h.r.ScopeStrategy()(accessRequest.GetClient().GetScopes(), scope) {
				accessRequest.GrantScope(scope)
			}
		}

		for _, audience := range accessRequest.GetRequestedAudience() {
			if h.r.AudienceStrategy()(accessRequest.GetClient().GetAudience(), []string{audience}) == nil {
				accessRequest.GrantAudience(audience)
			}
		}
	}

	accessResponse, err := h.r.OAuth2Provider().NewAccessResponse(ctx, accessRequest)
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.r.OAuth2Provider().WriteAccessError(w, accessRequest, err)
		return
	}

	h.r.OAuth2Provider().WriteAccessResponse(w, accessRequest, accessResponse)
}

// swagger:route GET /oauth2/auth public oauthAuth
//
// The OAuth 2.0 authorize endpoint
//
// This endpoint is not documented here because you should never use your own implementation to perform OAuth2 flows.
// OAuth2 is a very popular protocol and a library for your programming language will exists.
//
// To learn more about this flow please refer to the specification: https://tools.ietf.org/html/rfc6749
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Schemes: http, https
//
//     Responses:
//       302: emptyResponse
//       401: genericError
//       500: genericError
func (h *Handler) AuthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var ctx = r.Context()

	authorizeRequest, err := h.r.OAuth2Provider().NewAuthorizeRequest(ctx, r)
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	session, err := h.r.ConsentStrategy().HandleOAuth2AuthorizationRequest(w, r, authorizeRequest)
	if errors.Cause(err) == consent.ErrAbortOAuth2Request {
		// do nothing
		return
	} else if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	for _, scope := range session.GrantedScope {
		authorizeRequest.GrantScope(scope)
	}

	for _, audience := range session.GrantedAudience {
		authorizeRequest.GrantAudience(audience)
	}

	openIDKeyID, err := h.r.OpenIDJWTStrategy().GetPublicKeyID(r.Context())
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	var accessTokenKeyID string
	if h.c.AccessTokenStrategy() == "jwt" {
		accessTokenKeyID, err = h.r.AccessTokenJWTStrategy().GetPublicKeyID(r.Context())
		if err != nil {
			x.LogError(err, h.r.Logger())
			h.writeAuthorizeError(w, r, authorizeRequest, err)
			return
		}
	}

	authorizeRequest.SetID(session.Challenge)

	claims := &jwt.IDTokenClaims{
		Subject:                             session.ConsentRequest.SubjectIdentifier,
		Issuer:                              strings.TrimRight(h.c.IssuerURL().String(), "/") + "/",
		IssuedAt:                            time.Now().UTC(),
		AuthTime:                            session.AuthenticatedAt,
		RequestedAt:                         session.RequestedAt,
		Extra:                               session.Session.IDToken,
		AuthenticationContextClassReference: session.ConsentRequest.ACR,

		// We do not need to pass the audience because it's included directly by ORY Fosite
		// Audience:    []string{authorizeRequest.GetClient().GetID()},

		// This is set by the fosite strategy
		// ExpiresAt:   time.Now().Add(h.IDTokenLifespan).UTC(),
	}
	claims.Add("sid", session.ConsentRequest.LoginSessionID)

	// done
	response, err := h.r.OAuth2Provider().NewAuthorizeResponse(ctx, authorizeRequest, &Session{
		DefaultSession: &openid.DefaultSession{
			Claims: claims,
			Headers: &jwt.Headers{Extra: map[string]interface{}{
				// required for lookup on jwk endpoint
				"kid": openIDKeyID,
			}},
			Subject: session.ConsentRequest.Subject,
		},
		Extra:            session.Session.AccessToken,
		KID:              accessTokenKeyID,
		ClientID:         authorizeRequest.GetClient().GetID(),
		ConsentChallenge: session.Challenge,
	})
	if err != nil {
		x.LogError(err, h.r.Logger())
		h.writeAuthorizeError(w, r, authorizeRequest, err)
		return
	}

	h.r.OAuth2Provider().WriteAuthorizeResponse(w, authorizeRequest, response)
}

func (h *Handler) writeAuthorizeError(w http.ResponseWriter, r *http.Request, ar fosite.AuthorizeRequester, err error) {
	if !ar.IsRedirectURIValid() {
		h.forwardError(w, r, err)
		return
	}

	h.r.OAuth2Provider().WriteAuthorizeError(w, ar, err)
}

func (h *Handler) forwardError(w http.ResponseWriter, r *http.Request, err error) {
	rfErr := fosite.ErrorToRFC6749Error(err)
	query := url.Values{"error": {rfErr.Name}, "error_description": {rfErr.Description}, "error_hint": {rfErr.Hint}}

	if h.c.ShareOAuth2Debug() {
		query.Add("error_debug", rfErr.Debug)
	}

	http.Redirect(w, r, urlx.CopyWithQuery(h.c.ErrorURL(), query).String(), http.StatusFound)
}

// This function will not be called, OPTIONS request will be handled by cors
// this is just a placeholder.
func (h *Handler) handleOptions(w http.ResponseWriter, r *http.Request) {}
