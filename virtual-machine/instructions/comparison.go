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

// IEQ instruction
type IEQ struct{}

var _ vm.Instruction = IEQ{}

var IEqual = INEQ{}

func (i IEQ) Execute(vm *vm.VirtualMachine) {
	rhsOp := vm.CurrentStackFrame().Pop().(interpreter.IntegerValue)
	lhsOp := vm.CurrentStackFrame().Pop().(interpreter.IntegerValue)

	isEqual := lhsOp.Equal(nil, nil, rhsOp)
	vm.CurrentStackFrame().Push(interpreter.BoolValue(isEqual))
}

func (i IEQ) String() string {
	return fmt.Sprintf("IEQ")
}


// INEQ instruction
type INEQ struct{}

var _ vm.Instruction = INEQ{}

var INotEqual = INEQ{}

func (i INEQ) Execute(vm *vm.VirtualMachine) {
	rhsOp := vm.CurrentStackFrame().Pop().(interpreter.IntegerValue)
	lhsOp := vm.CurrentStackFrame().Pop().(interpreter.IntegerValue)

	isEqual := lhsOp.Equal(nil, nil, rhsOp)
	vm.CurrentStackFrame().Push(interpreter.BoolValue(!isEqual))
}

func (i INEQ) String() string {
	return fmt.Sprintf("INEQ")
}
