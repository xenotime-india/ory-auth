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

package cli

import (
	"fmt"
	"strings" //"encoding/json"

	"github.com/go-openapi/runtime"

	"github.com/ory/hydra/internal/httpclient/client/admin"
	"github.com/ory/x/pointerx"

	"github.com/spf13/cobra"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx" //"context"
)

type IntrospectionHandler struct{}

func newIntrospectionHandler() *IntrospectionHandler {
	return &IntrospectionHandler{}
}

func (h *IntrospectionHandler) Introspect(cmd *cobra.Command, args []string) {
	cmdx.ExactArgs(cmd, args, 1)
	c := configureClient(cmd)

	var ht runtime.ClientAuthInfoWriter
	if clientID, clientSecret := flagx.MustGetString(cmd, "client-id"), flagx.MustGetString(cmd, "client-secret"); clientID != "" || clientSecret != "" {
		ht = httptransport.BasicAuth(clientID, clientSecret)
	} else {
		fmt.Println("No OAuth 2.0 Client ID an secret set, skipping authorization header. This might fail if the introspection endpoint is protected.")
	}

	result, err := c.Admin.IntrospectOAuth2Token(admin.NewIntrospectOAuth2TokenParams().
		WithToken(args[0]).
		WithScope(pointerx.String(strings.Join(flagx.MustGetStringSlice(cmd, "scope"), " "))),
		ht,
	)
	cmdx.Must(err, "The request failed with the following error message:\n%s", formatSwaggerError(err))
	fmt.Println(formatResponse(result.Payload))
}
