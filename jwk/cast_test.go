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

package jwk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustRSAPrivate(t *testing.T) {
	//if testing.Short() {
	//	t.SkipNow()
	//}

	keys, err := new(RS256Generator).Generate("foo", "sig")
	assert.Nil(t, err)

	_, err = ToRSAPrivate(&keys.Key("private:foo")[0])
	assert.Nil(t, err)

	MustRSAPrivate(&keys.Key("private:foo")[0])

	_, err = ToRSAPublic(&keys.Key("public:foo")[0])
	assert.Nil(t, err)
	MustRSAPublic(&keys.Key("public:foo")[0])
}
