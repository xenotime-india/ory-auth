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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ory/x/httpx"

	"github.com/olekukonko/tablewriter"
	"github.com/sawadashota/encrypta"
	"github.com/spf13/cobra"

	httptransport "github.com/go-openapi/runtime/client"

	hydra "github.com/ory/hydra/internal/httpclient/client"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
)

func configureClient(cmd *cobra.Command) *hydra.OryHydra {
	return configureClientBase(cmd, true)
}

type transport struct {
	Transport http.RoundTripper
	cmd       *cobra.Command
}

func newTransport(cmd *cobra.Command) *transport {
	return &transport{
		cmd: cmd,
		Transport: httpx.NewResilientRoundTripper(
			&http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: flagx.MustGetBool(cmd, "skip-tls-verify")},
			},
			time.Second,
			flagx.MustGetDuration(cmd, "fail-after"),
		).WithShouldRetry(
			func(res *http.Response, err error) bool {
				if err != nil {
					fmt.Printf("Unable to connect: %s\n", err)
					return true
				} else if res.StatusCode == 0 || res.StatusCode >= 500 {
					fmt.Printf(`Unable to connect to "%s", unexpected HTTP error status code: %d\n`, res.Request.URL.String(), res.StatusCode)
					return true
				}
				return false
			},
		),
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if flagx.MustGetBool(t.cmd, "fake-tls-termination") {
		req.Header.Set("X-Forwarded-Proto", "https")
	}
	return t.Transport.RoundTrip(req)
}

func configureClientBase(cmd *cobra.Command, withAuth bool) *hydra.OryHydra {
	u := RemoteURI(cmd)
	ht := httptransport.New(
		u.Host,
		u.Path,
		[]string{u.Scheme},
	)

	ht.Transport = newTransport(cmd)

	if withAuth {
		if token := flagx.MustGetString(cmd, "access-token"); token != "" {
			ht.DefaultAuthentication = httptransport.BearerToken(token)
		}
	}

	return hydra.New(ht, nil)
}

func configureClientWithoutAuth(cmd *cobra.Command) *hydra.OryHydra {
	return configureClientBase(cmd, false)
}

func formatResponse(response interface{}) string {
	out, err := json.MarshalIndent(response, "", "\t")
	cmdx.Must(err, `Command failed because an error ("%s") occurred while prettifying output`, err)
	return string(out)
}

// newTable is table renderer at console
// And defines table layout option
//
// https://github.com/olekukonko/tablewriter
func newTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)

	// render options
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	return table
}

// newEncryptionKey for client secret
func newEncryptionKey(cmd *cobra.Command, client *http.Client) (ek encrypta.EncryptionKey, encryptSecret bool, err error) {
	if client == nil {
		client = http.DefaultClient
	}

	pgpKey := flagx.MustGetString(cmd, "pgp-key")
	pgpKeyURL := flagx.MustGetString(cmd, "pgp-key-url")
	keybaseUsername := flagx.MustGetString(cmd, "keybase")

	if pgpKey != "" {
		ek, err = encrypta.NewPublicKeyFromBase64Encoded(pgpKey)
		encryptSecret = true
		return
	}

	if pgpKeyURL != "" {
		ek, err = encrypta.NewPublicKeyFromURL(pgpKeyURL, encrypta.HTTPClientOption(client))
		encryptSecret = true
		return
	}

	if keybaseUsername != "" {
		ek, err = encrypta.NewPublicKeyFromKeybase(keybaseUsername, encrypta.HTTPClientOption(client))
		encryptSecret = true
		return
	}

	return nil, false, nil
}
