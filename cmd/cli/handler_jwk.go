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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ory/hydra/internal/httpclient/client/admin"
	"github.com/ory/hydra/internal/httpclient/models"
	"github.com/ory/x/pointerx"

	"github.com/mendsley/gojwk"
	"github.com/pborman/uuid"
	"github.com/spf13/cobra"
	jose "gopkg.in/square/go-jose.v2"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/ory/x/josex"
)

type JWKHandler struct{}

func newJWKHandler() *JWKHandler {
	return &JWKHandler{}
}

func (h *JWKHandler) CreateKeys(cmd *cobra.Command, args []string) {
	cmdx.RangeArgs(cmd, args, []int{1, 2})
	m := configureClient(cmd)

	var kid string
	if len(args) == 2 {
		kid = args[1]
	}

	res, err := m.Admin.CreateJSONWebKeySet(admin.NewCreateJSONWebKeySetParams().WithSet(args[0]).WithBody(&models.JSONWebKeySetGeneratorRequest{
		Alg: pointerx.String(flagx.MustGetString(cmd, "alg")),
		Kid: pointerx.String(kid),
		Use: pointerx.String(flagx.MustGetString(cmd, "use")),
	}))
	cmdx.Must(err, "The request failed with the following error message:\n%s", formatSwaggerError(err))
	fmt.Println(formatResponse(res.Payload))
}

func toSDKFriendlyJSONWebKey(key interface{}, kid, use string, public bool) jose.JSONWebKey {
	if jwk, ok := key.(*jose.JSONWebKey); ok {
		key = jwk.Key
		if jwk.KeyID != "" {
			kid = jwk.KeyID
		}
		if jwk.Use != "" {
			use = jwk.Use
		}
	}

	var err error
	var jwk *gojwk.Key
	if public {
		jwk, err = gojwk.PublicKey(key)
		cmdx.Must(err, "Unable to convert public key to JSON Web Key because %s", err)
	} else {
		jwk, err = gojwk.PrivateKey(key)
		cmdx.Must(err, "Unable to convert private key to JSON Web Key because %s", err)
	}

	return jose.JSONWebKey{
		KeyID:     kid,
		Use:       use,
		Algorithm: jwk.Alg,
		Key:       key,
	}
}

func (h *JWKHandler) ImportKeys(cmd *cobra.Command, args []string) {
	cmdx.MinArgs(cmd, args, 2)

	id := args[0]
	use := flagx.MustGetString(cmd, "use")
	client := &http.Client{}

	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: flagx.MustGetBool(cmd, "skip-tls-verify"),
		},
	}

	u := Remote(cmd) + "/keys/" + id
	request, err := http.NewRequest("GET", u, nil)
	cmdx.Must(err, "Unable to initialize HTTP request: %s", err)

	if flagx.MustGetBool(cmd, "fake-tls-termination") {
		request.Header.Set("X-Forwarded-Proto", "https")
	}

	if token := flagx.MustGetString(cmd, "access-token"); token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	response, err := client.Do(request)
	cmdx.Must(err, "Unable to fetch data from %s: %s", u, err)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNotFound {
		cmdx.Fatalf("Expected status code 200 or 404 but got %d while fetching data from %s", response.StatusCode, u)
	}

	var set jose.JSONWebKeySet
	err = json.NewDecoder(response.Body).Decode(&set)
	cmdx.Must(err, "Unable to decode payload to JSON: %s", err)

	for _, path := range args[1:] {
		file, err := ioutil.ReadFile(path)
		cmdx.Must(err, "Unable to read file %s", path)

		if key, privateErr := josex.LoadPrivateKey(file); privateErr != nil {
			key, publicErr := josex.LoadPublicKey(file)
			cmdx.Must(publicErr, `Unable to read key from file %s. Decoding file to private key failed with reason "%s" and decoding it to public key failed with reason: %s`, path, privateErr, publicErr)

			set.Keys = append(set.Keys, toSDKFriendlyJSONWebKey(key, "public:"+uuid.New(), use, true))
		} else {
			set.Keys = append(set.Keys, toSDKFriendlyJSONWebKey(key, "private:"+uuid.New(), use, false))
		}

		fmt.Printf("Successfully loaded key from file: %s\n", path)
	}

	body, err := json.Marshal(&set)
	cmdx.Must(err, "Unable to encode JSON Web Keys to JSON: %s", err)

	request, err = http.NewRequest("PUT", u, bytes.NewReader(body))
	cmdx.Must(err, "Unable to initialize HTTP request: %s", err)

	if flagx.MustGetBool(cmd, "fake-tls-termination") {
		request.Header.Set("X-Forwarded-Proto", "https")
	}

	if token := flagx.MustGetString(cmd, "access-token"); token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = client.Do(request)
	cmdx.CheckResponse(err, http.StatusOK, response)
	defer response.Body.Close()

	fmt.Println("JSON Web Key Set successfully imported!")
}

func (h *JWKHandler) GetKeys(cmd *cobra.Command, args []string) {
	cmdx.ExactArgs(cmd, args, 1)
	m := configureClient(cmd)

	keys, err := m.Admin.GetJSONWebKeySet(admin.NewGetJSONWebKeySetParams().WithSet(args[0]))
	cmdx.Must(err, "The request failed with the following error message:\n%s", formatSwaggerError(err))
	fmt.Printf("%s\n", formatResponse(keys))
}

func (h *JWKHandler) DeleteKeys(cmd *cobra.Command, args []string) {
	cmdx.ExactArgs(cmd, args, 1)
	m := configureClient(cmd)

	_, err := m.Admin.DeleteJSONWebKeySet(admin.NewDeleteJSONWebKeySetParams().WithSet(args[0]))
	cmdx.Must(err, "The request failed with the following error message:\n%s", formatSwaggerError(err))
	fmt.Printf("JSON Web Key Set deleted: %s\n", args[0])
}
