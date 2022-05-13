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

package tests

import (
	"fmt"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onflow/cadence/runtime/ast"
	"github.com/onflow/cadence/runtime/parser2"

	"github.com/onflow/cadence/virtual-machine"
	"github.com/onflow/cadence/virtual-machine/codegen"
	"github.com/onflow/cadence/virtual-machine/instructions"
)

const LOOP_COUNT = 1000

func TestVM(t *testing.T) {
	ins := []virtual_machine.Instruction{
		instructions.NewFunction("foo"),
		instructions.IConst{big.NewInt(0)},
		instructions.NewStore(0), // result-var index

		instructions.IConst{big.NewInt(0)},
		instructions.NewStore(1), // loop-var index

		// start of loop
		instructions.Load{1},
		instructions.IConst{big.NewInt(LOOP_COUNT)},
		instructions.NotEqual{},
		instructions.JumpIf{10}, // if true, jump to loop body
		instructions.Jump{19},   // jump to end-of-loop

		// loop body

		// update loop variable
		instructions.IConst{big.NewInt(1)}, // load 1
		instructions.Load{1},               // load loop-var
		instructions.IntegerAdd{},          // add 1 to loop-var
		instructions.NewStore(1),           // store loop-var

		// update result variable: increment by 5
		instructions.IConst{big.NewInt(5)}, // load 5
		instructions.Load{0},               // load result-var
		instructions.IntegerAdd{},          // add 1 to result-var
		instructions.NewStore(0),           // store result-var

		instructions.Jump{5}, // go to start of loop (condition)

		// end of loop

		// print result
		instructions.Load{0},
		instructions.Print{},

		instructions.Stop{},
	}

	fmt.Println("Instructions:")
	for _, instruction := range ins {
		fmt.Println(instruction.String())
	}

	fmt.Println("\nExecuting:")
	vm := virtual_machine.NewVirtualMachine()
	vm.Execute(ins)

}

func TestCodeGen(t *testing.T) {
	program, err := parser2.ParseProgram(`
        fun test() {
			var i = 0
			var result = 0
			while i != 1000 {
				i = i + 1
				result = result + 5
			}
	    }
	`)
	assert.NoError(t, err)

	codeGen := codegen.NewCodeGenerator()
	instructions := codeGen.Generate(program)

	fmt.Println("Instructions:")
	for _, instruction := range instructions {
		fmt.Println(instruction.String())
	}

	fmt.Println("\nExecuting:")
	vm := virtual_machine.NewVirtualMachine()
	vm.Execute(instructions)
}

func BenchmarkInstructions(b *testing.B) {
	instructions := []virtual_machine.Instruction{
		instructions.NewFunction("foo"),
		instructions.I64Const{0},
		instructions.NewStore(0), // result-var index

		instructions.I64Const{0},
		instructions.NewStore(1), // loop-var index

		// start of loop
		instructions.Load{0},
		instructions.I64Const{1000},
		instructions.NotEqual{},
		instructions.JumpIf{10}, // if true, jump to loop body
		instructions.Jump{19},   // jump to end-of-loop

		// loop body

		// update loop variable
		instructions.I64Const{1},  // load 1
		instructions.Load{0},      // load loop-var
		instructions.IntegerAdd{}, // add 1 to loop-var
		instructions.NewStore(0),  // store loop-var

		// update result variable
		instructions.I64Const{2},  // load 2
		instructions.Load{1},      // load result-var
		instructions.IntegerAdd{}, // add 1 to result-var
		instructions.NewStore(1),  // store result-var

		instructions.Jump{4}, // go to start of loop (condition)

		// end of loop

		instructions.Stop{},
	}

	vm := virtual_machine.NewVirtualMachine()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		vm.Execute(instructions)
	}
}

func BenchmarkCodeGen(b *testing.B) {
	program, err := parser2.ParseProgram(`
        fun test() {
			var i = 0
			var result = 0
			while i != 1000 {
				i = i + 1
				result = result + 5
			}
	    }
	`)
	assert.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		codeGen := codegen.NewCodeGenerator()
		codeGen.Generate(program)
	}
}

func BenchmarkVM(b *testing.B) {
	program, err := parser2.ParseProgram(`
        fun test() {
			var i = 0
			var result = 0
			while i != 1000 {
				i = i + 1
				result = result + 5
			}
	    }
	`)
	assert.NoError(b, err)

	codeGen := codegen.NewCodeGenerator()
	instructions := codeGen.Generate(program)
	vm := virtual_machine.NewVirtualMachine()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		vm.Execute(instructions)
	}
}

func BenchmarkCodeGenAndVM(b *testing.B) {
	program, err := parser2.ParseProgram(`
        fun test() {
			var i = 0
			var result = 0
			while i != 1000 {
				i = i + 1
				result = result + 5
			}
	    }
	`)
	assert.NoError(b, err)

	b.ResetTimer()

	vm := virtual_machine.NewVirtualMachine()

	for i := 0; i < b.N; i++ {
		codeGen := codegen.NewCodeGenerator()
		instructions := codeGen.Generate(program)
		vm.Execute(instructions)
	}
}

func BenchmarkASTVisitor(b *testing.B) {
	program, err := parser2.ParseProgram(`
        fun test() {
			var i = 0
			var result = 0
			while i != 1000 {
				i = i + 1
				result = result + 5
			}
	    }
	`)
	assert.NoError(b, err)

	block := program.FunctionDeclarations()[0].FunctionBlock.Block

	v := TestASTVisitor{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		block.Accept(v)
	}
}

type TestASTVisitor struct{}

var _ ast.Visitor = TestASTVisitor{}

func (t TestASTVisitor) VisitReturnStatement(statement *ast.ReturnStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitBreakStatement(statement *ast.BreakStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitContinueStatement(statement *ast.ContinueStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitIfStatement(statement *ast.IfStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitSwitchStatement(statement *ast.SwitchStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitWhileStatement(statement *ast.WhileStatement) ast.Repr {
	for i := 0; i < LOOP_COUNT; i++ {
		statement.Test.Accept(t)
		statement.Block.Accept(t)
	}

	return true
}

func (t TestASTVisitor) VisitForStatement(statement *ast.ForStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitEmitStatement(statement *ast.EmitStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitVariableDeclaration(declaration *ast.VariableDeclaration) ast.Repr {
	return true
}

func (t TestASTVisitor) VisitAssignmentStatement(statement *ast.AssignmentStatement) ast.Repr {
	return true
}

func (t TestASTVisitor) VisitSwapStatement(statement *ast.SwapStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitExpressionStatement(statement *ast.ExpressionStatement) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitBoolExpression(expression *ast.BoolExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitNilExpression(expression *ast.NilExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitIntegerExpression(expression *ast.IntegerExpression) ast.Repr {
	return interpreter.NewIntValue(expression.Value, sema.Int64Type)
}

func (t TestASTVisitor) VisitFixedPointExpression(expression *ast.FixedPointExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitArrayExpression(expression *ast.ArrayExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitDictionaryExpression(expression *ast.DictionaryExpression) ast.Repr {
	panic("implement me")
}

var tmp = interpreter.NewIntValueFromInt64(2)

func (t TestASTVisitor) VisitIdentifierExpression(expression *ast.IdentifierExpression) ast.Repr {
	return tmp
}

func (t TestASTVisitor) VisitInvocationExpression(expression *ast.InvocationExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitMemberExpression(expression *ast.MemberExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitIndexExpression(expression *ast.IndexExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitConditionalExpression(expression *ast.ConditionalExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitUnaryExpression(expression *ast.UnaryExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitBinaryExpression(expression *ast.BinaryExpression) ast.Repr {
	lhs := expression.Left.Accept(t).(interpreter.IntegerValue)
	rhs := expression.Right.Accept(t).(interpreter.IntegerValue)

	switch expression.Operation {
	case ast.OperationPlus:
		lhs.Plus(rhs)
	case ast.OperationNotEqual:
		lhs.Equal(nil, nil, rhs)
	default:
		panic("Unsupported")
	}

	return true
}

func (t TestASTVisitor) VisitFunctionExpression(expression *ast.FunctionExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitStringExpression(expression *ast.StringExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitCastingExpression(expression *ast.CastingExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitCreateExpression(expression *ast.CreateExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitDestroyExpression(expression *ast.DestroyExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitReferenceExpression(expression *ast.ReferenceExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitForceExpression(expression *ast.ForceExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitPathExpression(expression *ast.PathExpression) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitProgram(program *ast.Program) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitFunctionDeclaration(declaration *ast.FunctionDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitBlock(block *ast.Block) ast.Repr {
	for _, statement := range block.Statements {
		statement.Accept(t)
	}

	return true
}

func (t TestASTVisitor) VisitFunctionBlock(block *ast.FunctionBlock) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitCompositeDeclaration(declaration *ast.CompositeDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitInterfaceDeclaration(declaration *ast.InterfaceDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitFieldDeclaration(declaration *ast.FieldDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitEnumCaseDeclaration(declaration *ast.EnumCaseDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitPragmaDeclaration(declaration *ast.PragmaDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitImportDeclaration(declaration *ast.ImportDeclaration) ast.Repr {
	panic("implement me")
}

func (t TestASTVisitor) VisitTransactionDeclaration(declaration *ast.TransactionDeclaration) ast.Repr {
	panic("implement me")
}

func BenchmarkIntAdd(b *testing.B) {
	var x interpreter.NumberValue = interpreter.Int64Value(1)
	var y interpreter.NumberValue = interpreter.Int64Value(2)

	for i := 0; i < b.N; i++ {
		for i := 0; i < LOOP_COUNT; i++ {
			_ = y.Plus(x)
			_ = y.Plus(x)
			_ = y.Equal(nil, nil, x)
		}
	}
}

func BenchmarkContext(b *testing.B) {

	b.Run("push", func(b *testing.B) {
		vm := virtual_machine.NewVirtualMachine()
		ctx := vm.NewContext()
		ctx.CallStack.PushNew()

		sf := ctx.CurrentStackFrame()
		sf.Push(interpreter.Int64Value(1))

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			for i := 0; i < LOOP_COUNT; i++ {
				sf.Push(interpreter.Int64Value(1))
				sf.Push(interpreter.Int64Value(1))
				sf.Push(interpreter.Int64Value(1))
			}
		}

	})

	b.Run("pop", func(b *testing.B) {
		vm := virtual_machine.NewVirtualMachine()
		ctx := vm.NewContext()
		ctx.CallStack.PushNew()

		sf := ctx.CurrentStackFrame()
		sf.Push(interpreter.Int64Value(1))

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			for i := 0; i < LOOP_COUNT; i++ {
				sf.Push(interpreter.Int64Value(1))
				sf.Push(interpreter.Int64Value(1))
				sf.Pop()
				sf.Pop()
			}
		}

	})
}

func BenchmarkInstruction(b *testing.B) {
	ins := make([]virtual_machine.Instruction, 0, LOOP_COUNT)

	for i := 0; i < LOOP_COUNT; i++ {
		ins = append(ins, instructions.NewFunction("foo"))
	}

	ins = append(ins, instructions.Stop{})

	vm := virtual_machine.NewVirtualMachine()

	for i := 0; i < b.N; i++ {
		vm.Execute(ins)
	}
}
