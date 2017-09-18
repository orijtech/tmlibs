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

package process

import (
	. "github.com/tendermint/tmlibs/common"
)

// Runs a command and gets the result.
func Run(dir string, command string, args []string) (string, bool, error) {
	outFile := NewBufferCloser(nil)
	proc, err := StartProcess("", dir, command, args, nil, outFile)
	if err != nil {
		return "", false, err
	}

	<-proc.WaitCh

	if proc.ExitState.Success() {
		return string(outFile.Bytes()), true, nil
	} else {
		return string(outFile.Bytes()), false, nil
	}
}
                                                                                                                                                 