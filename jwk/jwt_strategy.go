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

package jwk

import (
	"context"
	"crypto/rsa"
	"strings"
	"sync"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/ory/fosite/token/jwt"
)

type JWTStrategy interface {
	GetPublicKeyID(ctx context.Context) (string, error)

	jwt.JWTStrategy
}

type RS256JWTStrategy struct {
	sync.RWMutex

	RS256JWTStrategy *jwt.RS256JWTStrategy
	r                InternalRegistry
	c                Configuration
	rs               func() string

	publicKey    *rsa.PublicKey
	privateKey   *rsa.PrivateKey
	publicKeyID  string
	privateKeyID string
}

func NewRS256JWTStrategy(r InternalRegistry, rs func() string) (*RS256JWTStrategy, error) {
	j := &RS256JWTStrategy{r: r, rs: rs, RS256JWTStrategy: new(jwt.RS256JWTStrategy)}
	if err := j.refresh(context.TODO()); err != nil {
		return nil, err
	}
	return j, nil
}

func (j *RS256JWTStrategy) Hash(ctx context.Context, in []byte) ([]byte, error) {
	return j.RS256JWTStrategy.Hash(ctx, in)
}

// GetSigningMethodLength will return the length of the signing method
func (j *RS256JWTStrategy) GetSigningMethodLength() int {
	return j.RS256JWTStrategy.GetSigningMethodLength()
}

func (j *RS256JWTStrategy) GetSignature(ctx context.Context, token string) (string, error) {
	return j.RS256JWTStrategy.GetSignature(ctx, token)
}

func (j *RS256JWTStrategy) Generate(ctx context.Context, claims jwt2.Claims, header jwt.Mapper) (string, string, error) {
	if err := j.refresh(ctx); err != nil {
		return "", "", err
	}

	return j.RS256JWTStrategy.Generate(ctx, claims, header)
}

func (j *RS256JWTStrategy) Validate(ctx context.Context, token string) (string, error) {
	if err := j.refresh(ctx); err != nil {
		return "", err
	}

	return j.RS256JWTStrategy.Validate(ctx, token)
}

func (j *RS256JWTStrategy) Decode(ctx context.Context, token string) (*jwt2.Token, error) {
	if err := j.refresh(ctx); err != nil {
		return nil, err
	}

	return j.RS256JWTStrategy.Decode(ctx, token)
}

func (j *RS256JWTStrategy) GetPublicKeyID(ctx context.Context) (string, error) {
	if err := j.refresh(ctx); err != nil {
		return "", err
	}

	return j.publicKeyID, nil
}

func (j *RS256JWTStrategy) refresh(ctx context.Context) error {
	keys, err := j.r.KeyManager().GetKeySet(ctx, j.rs())
	if err != nil {
		return err
	}

	public, err := FindKeyByPrefix(keys, "public")
	if err != nil {
		return err
	}

	private, err := FindKeyByPrefix(keys, "private")
	if err != nil {
		return err
	}

	if strings.Replace(public.KeyID, "public:", "", 1) != strings.Replace(private.KeyID, "private:", "", 1) {
		return errors.New("public and private key pair kids do not match")
	}

	if k, ok := private.Key.(*rsa.PrivateKey); !ok {
		return errors.New("unable to type assert key to *rsa.PublicKey")
	} else {
		j.Lock()
		j.privateKey = k
		j.RS256JWTStrategy.PrivateKey = k
		j.Unlock()
	}

	if k, ok := public.Key.(*rsa.PublicKey); !ok {
		return errors.New("unable to type assert key to *rsa.PublicKey")
	} else {
		j.Lock()
		j.publicKey = k
		j.publicKeyID = public.KeyID
		j.Unlock()
	}

	j.RLock()
	defer j.RUnlock()
	if j.privateKey.PublicKey.E != j.publicKey.E ||
		j.privateKey.PublicKey.N.String() != j.publicKey.N.String() {
		return errors.New("public and private key pair fetched from store does not match")
	}

	return nil
}
