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
	"github.com/onflow/cadence/runtime/interpreter"
	vm "github.com/onflow/cadence/virtual-machine"
)

// Contains all jump instructions.

// Jump instruction
type Jump struct {
	Instruction int
}

var _ vm.Instruction = Jump{}

func NewJump(index int) Jump {
	return Jump{
		Instruction: index,
	}
}

func (i Jump) Execute(ctx *vm.ExecutionContext) {
	ctx.NextIndex = i.Instruction
}

func (i Jump) String() string {
	return fmt.Sprintf("JUMP %d", i.Instruction)
}

// JumpIf instruction
type JumpIf struct {
	Instruction int //Instruction int // instruction to jump to, if true
}

var _ vm.Instruction = JumpIf{}

func NewJumpIf(index int) JumpIf {
	return JumpIf{
		Instruction: index,
	}
}

func (i JumpIf) Execute(ctx *vm.ExecutionContext) {
	condition := ctx.CurrentStackFrame().Pop().(interpreter.BoolValue)
	if condition {
		ctx.NextIndex = i.Instruction
	}
}

func (i JumpIf) String() string {
	return fmt.Sprintf("JUMPIF %d", i.Instruction)
}
