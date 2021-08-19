// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package idutil implements utility functions for generating unique,
// randomized ids.
package idutil

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Next(t *testing.T) {
	now := time.Now()
	g1 := NewGenerator(1, now)
	g2 := NewGenerator(2, now)
	id1 := g1.Next()
	id2 := g1.Next()
	id3 := g2.Next()
	id4 := g2.Next()
	log.Println("id1 = ", id1)
	log.Println("id2 = ", id2)
	log.Println("id3 = ", id3)
	log.Println("id4 = ", id4)
	assert.NotEqual(t, id1, id2)
	assert.NotEqual(t, id3, id4)
}
