/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"testing"

	"gotest.tools/assert"
)

func Test_doTranslate(t *testing.T) {
	type argsType struct {
		args []string
	}

	tests := []struct {
		name string
		args argsType
	}{
		{
			name: "测试英文",
			args: argsType{[]string{"dog"}},
		},
		{
			name: "测试中文",
			args: argsType{[]string{"狗"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := doTranslate(tt.args.args)
			assert.Equal(t, value, 0)
		})
	}
}
