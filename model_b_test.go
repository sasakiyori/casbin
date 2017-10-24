// Copyright 2017 The casbin Authors. All Rights Reserved.
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

package casbin

import (
	"fmt"
	"testing"
)

func rawEnforce(sub string, obj string, act string) bool {
	policy := [2][3]string{{"alice", "data1", "read"}, {"bob", "data2", "write"}}
	for _, rule := range policy {
		if sub == rule[0] && obj == rule[1] && act == rule[2] {
			return true
		}
	}
	return false
}

func BenchmarkRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawEnforce("alice", "data1", "read")
	}
}

func BenchmarkBasicModel(b *testing.B) {
	e := NewEnforcer("examples/basic_model.conf", "examples/basic_policy.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkRBACModel(b *testing.B) {
	e := NewEnforcer("examples/rbac_model.conf", "examples/rbac_policy.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "data2", "read")
	}
}

func BenchmarkRBACModelLarge(b *testing.B) {
	e := NewEnforcer("examples/rbac_model.conf")
	// 1000 users, 100 roles, 10 resources.
	for i := 0; i < 100; i++ {
		e.AddPolicy(fmt.Sprintf("group%d", i), fmt.Sprintf("data%d", i / 10), "read")
	}

	for i := 0; i < 1000; i++ {
		e.AddGroupingPolicy(fmt.Sprintf("user%d", i), fmt.Sprintf("group%d", i / 10))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("user501", "data9", "read")
	}
}

func BenchmarkRBACModelWithResourceRoles(b *testing.B) {
	e := NewEnforcer("examples/rbac_model_with_resource_roles.conf", "examples/rbac_policy_with_resource_roles.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkRBACModelWithDomains(b *testing.B) {
	e := NewEnforcer("examples/rbac_model_with_domains.conf", "examples/rbac_policy_with_domains.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "domain1", "data1", "read")
	}
}

func BenchmarkABACModel(b *testing.B) {
	e := NewEnforcer("examples/abac_model.conf")
	data1 := newTestResource("data1", "alice")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", data1, "read")
	}
}

func BenchmarkKeyMatchModel(b *testing.B) {
	e := NewEnforcer("examples/keymatch_model.conf", "examples/keymatch_policy.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "/alice_data/resource1", "GET")
	}
}

func BenchmarkRBACModelWithDeny(b *testing.B) {
	e := NewEnforcer("examples/rbac_model_with_deny.conf", "examples/rbac_policy_with_deny.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "data1", "read")
	}
}

func BenchmarkPriorityModel(b *testing.B) {
	e := NewEnforcer("examples/priority_model.conf", "examples/priority_policy.csv")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Enforce("alice", "data1", "read")
	}
}