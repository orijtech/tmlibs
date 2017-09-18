// Copyright 2017 Tendermint. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tmlibs/pubsub/query"
)

func TestEmptyQueryMatchesAnything(t *testing.T) {
	q := query.Empty{}
	assert.True(t, q.Matches(map[string]interface{}{}))
	assert.True(t, q.Matches(map[string]interface{}{"Asher": "Roth"}))
	assert.True(t, q.Matches(map[string]interface{}{"Route": 66}))
	assert.True(t, q.Matches(map[string]interface{}{"Route": 66, "Billy": "Blue"}))
}
                                                                                                                                                              