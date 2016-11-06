/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flags

import (
	"strconv"
	"strings"
)

const Version = "0.11.3"

type version []int

func ParseVersion(s string) (version, error) {
	v := make(version, 0)
	ps := strings.Split(s, ".")
	for _, p := range ps {
		i, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}

		v = append(v, i)
	}

	return v, nil
}

func (v version) Lte(u version) bool {
	lv := len(v)
	lu := len(u)

	for i := 0; i < lv; i++ {
		// Everything up to here has been equal and v has more elements than u.
		if i >= lu {
			return false
		}

		// Move to next digit if equal.
		if v[i] == u[i] {
			continue
		}

		return v[i] < u[i]
	}

	// Equal.
	return true
}
