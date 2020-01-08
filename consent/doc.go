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

package consent

// swagger:parameters getLoginRequest
type swaggerGetLoginRequestByChallenge struct {
	// in: query
	// required: true
	Challenge string `json:"login_challenge"`
}

// swagger:parameters getConsentRequest
type swaggerGetConsentRequestByChallenge struct {
	// in: query
	// required: true
	Challenge string `json:"consent_challenge"`
}

// swagger:parameters getLogoutRequest
type swaggerGetLogoutRequestByChallenge struct {
	// in: query
	// required: true
	Challenge string `json:"logout_challenge"`
}

// swagger:parameters revokeConsentSessions
type swaggerRevokeConsentSessions struct {
	// The subject (Subject) who's consent sessions should be deleted.
	//
	// in: query
	// required: true
	Subject string `json:"subject"`

	// If set, deletes only those consent sessions by the Subject that have been granted to the specified OAuth 2.0 Client ID
	//
	// in: query
	Client string `json:"client"`
}

// swagger:parameters listSubjectConsentSessions
type swaggerListSubjectConsentSessionsPayload struct {
	// in: query
	// required: true
	Subject string `json:"subject"`
}

// swagger:parameters revokeAuthenticationSession
type swaggerRevokeAuthenticationSessionPayload struct {
	// in: query
	// required: true
	Subject string `json:"subject"`
}

// swagger:parameters acceptLoginRequest
type swaggerAcceptLoginRequest struct {
	// in: query
	// required: true
	Challenge string `json:"login_challenge"`

	// in: body
	Body HandledLoginRequest
}

// swagger:parameters acceptConsentRequest
type swaggerAcceptConsentRequest struct {
	// in: query
	// required: true
	Challenge string `json:"consent_challenge"`

	// in: body
	Body HandledConsentRequest
}

// swagger:parameters acceptLogoutRequest
type swaggerAcceptLogoutRequest struct {
	// in: query
	// required: true
	Challenge string `json:"logout_challenge"`
}

// swagger:parameters rejectConsentRequest
type swaggerRejectConsentRequest struct {
	// in: query
	// required: true
	Challenge string `json:"consent_challenge"`

	// in: body
	Body RequestDeniedError
}

// swagger:parameters rejectLoginRequest
type swaggerRejectLoginRequest struct {
	// in: query
	// required: true
	Challenge string `json:"login_challenge"`

	// in: body
	Body RequestDeniedError
}

// swagger:parameters rejectLogoutRequest
type swaggerRejectLogoutRequest struct {
	// in: query
	// required: true
	Challenge string `json:"logout_challenge"`

	// in: body
	Body RequestDeniedError
}

// A list of used consent requests.
// swagger:response handledConsentRequestList
type swaggerListHandledConsentRequestsResult struct {
	// in: body
	// type: array
	Body []PreviousConsentSession
}
