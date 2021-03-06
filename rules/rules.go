// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rules contains implementations of rules that apply to all Google APIs.
package rules

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/rules/aip0131"
	"github.com/googleapis/api-linter/rules/aip0132"
	"github.com/googleapis/api-linter/rules/aip0140"
	"github.com/googleapis/api-linter/rules/aip0158"
	"github.com/googleapis/api-linter/rules/aip0191"
)

func init() {
	aip0131.AddRules(coreRules)
	aip0132.AddRules(coreRules)
	aip0140.AddRules(coreRules)
	aip0158.AddRules(coreRules)
	aip0191.AddRules(coreRules)
}

var coreRules, _ = lint.NewRuleRegistry()

// Rules returns all rules registered in this package.
func Rules() lint.RuleRegistry {
	return coreRules.Copy()
}
