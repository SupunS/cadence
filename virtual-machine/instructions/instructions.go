/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package instructions

import (
	"fmt"

	vm "github.com/onflow/cadence/virtual-machine"
)

// Print instruction.
// This is a temporary instruction for testing purpose.
type Print struct{}

var _ vm.Instruction = Print{}

func (i Print) Execute(ctx *vm.ExecutionContext) {
	fmt.Println(ctx.CurrentStackFrame().Pop())
}

func (i Print) String() string {
	return "PRINT"
}

// Stop instruction
type Stop struct{}

var _ vm.Instruction = Stop{}

var StopIns = Stop{}

func (i Stop) Execute(ctx *vm.ExecutionContext) {
	ctx.NextIndex = vm.NO_OP
}

func (i Stop) String() string {
	return "STOP"
}
