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

package codegen

import (
	"github.com/onflow/cadence/runtime/ast"
	vm "github.com/onflow/cadence/virtual-machine"
	"github.com/onflow/cadence/virtual-machine/instructions"
)

type Scope struct {
	indexMapping map[string]int
	nextVarIndex int
	parent       *Scope
}

func NewScope() Scope {
	return Scope{
		indexMapping: make(map[string]int),
		nextVarIndex: 0,
		parent:       nil,
	}
}

func (c *Scope) NextLocalVarIndex() int {
	nexIndex := c.nextVarIndex
	c.nextVarIndex++
	return nexIndex
}

func (c *Scope) Add(name string, varIndex int) {
	c.indexMapping[name] = varIndex
}

func (c *Scope) GetVarIndex(name string) int {
	return c.indexMapping[name]
}

type CodeGenerator struct {
	Instructions      []vm.Instruction
	instructionOffset int
	scope             Scope
}

var _ ast.Visitor = &CodeGenerator{}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (c *CodeGenerator) Generate(program *ast.Program) []vm.Instruction {
	program.Accept(c)
	c.emit(instructions.Print{})
	c.emit(instructions.StopIns)
	return c.Instructions
}

func (c *CodeGenerator) emit(ins vm.Instruction) {
	c.Instructions = append(c.Instructions, ins)
	c.instructionOffset++
}

func (c *CodeGenerator) nextInstruction() int {
	return c.instructionOffset
}

func (c *CodeGenerator) VisitReturnStatement(returnStatement *ast.ReturnStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitBreakStatement(breakStatement *ast.BreakStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitContinueStatement(continueStatement *ast.ContinueStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitIfStatement(ifStatement *ast.IfStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitSwitchStatement(switchStatement *ast.SwitchStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitWhileStatement(whileStatement *ast.WhileStatement) ast.Repr {
	conditionIndex := c.nextInstruction()

	// emit instruction for condition
	whileStatement.Test.Accept(c)

	// There would be 2 jump instructions after the condition, before the start of body.
	// Therefore, body start index = next instruction + 2
	// IMPORTANT: order matters: nextInstruction should be called before visiting body.
	bodyStartIndex := c.nextInstruction() + 2

	// visit the body first, because condition need to know the end index of the body.
	whileBodyInstructions := c.generateBranch(whileStatement.Block)

	c.emit(instructions.NewJumpIf(bodyStartIndex))

	// body end index = total instructions + 2 jumps (2 for condition, 1 for loop-back)
	bodyEndIndex := c.nextInstruction() + 2
	c.emit(instructions.NewJump(bodyEndIndex))

	c.Instructions = append(c.Instructions, whileBodyInstructions...)

	// go back to condition
	c.emit(instructions.NewJump(conditionIndex))

	return nil
}

func (c *CodeGenerator) generateBranch(block *ast.Block) []vm.Instruction {
	prevInstructions := c.Instructions
	c.Instructions = nil

	block.Accept(c)

	branchInstructions := c.Instructions
	c.Instructions = prevInstructions

	return branchInstructions
}

func (c *CodeGenerator) VisitForStatement(forStatement *ast.ForStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitEmitStatement(emitStatement *ast.EmitStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitVariableDeclaration(variableDeclaration *ast.VariableDeclaration) ast.Repr {
	varIndex := c.scope.NextLocalVarIndex()

	variableDeclaration.Value.Accept(c)

	// store into local variable
	c.emit(instructions.NewIStore(varIndex))

	varName := variableDeclaration.Identifier.Identifier
	c.scope.Add(varName, varIndex)

	return nil
}

func (c *CodeGenerator) VisitAssignmentStatement(assignmentStatement *ast.AssignmentStatement) ast.Repr {
	assignmentStatement.Value.Accept(c)

	var varIndex int

	switch target := assignmentStatement.Target.(type) {
	case *ast.IdentifierExpression:
		varIndex = c.scope.GetVarIndex(target.Identifier.Identifier)
	case *ast.IndexExpression, *ast.MemberExpression:
		panic("Unsupported")
	default:
		panic("Unsupported")
	}

	c.emit(instructions.NewIStore(varIndex))
	return nil
}

func (c *CodeGenerator) VisitSwapStatement(swapStatement *ast.SwapStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitExpressionStatement(expressionStatement *ast.ExpressionStatement) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitBoolExpression(boolExpression *ast.BoolExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitNilExpression(nilExpression *ast.NilExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitIntegerExpression(integerExpression *ast.IntegerExpression) ast.Repr {
	c.emit(instructions.NewIConst(integerExpression.Value))
	return nil
}

func (c *CodeGenerator) VisitFixedPointExpression(fixedPointExpression *ast.FixedPointExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitArrayExpression(arrayExpression *ast.ArrayExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitDictionaryExpression(dictionaryExpression *ast.DictionaryExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitIdentifierExpression(identifierExpression *ast.IdentifierExpression) ast.Repr {
	varName := identifierExpression.Identifier.Identifier
	varIndex := c.scope.GetVarIndex(varName)

	c.emit(instructions.NewLoad(varIndex))

	return nil
}

func (c *CodeGenerator) VisitInvocationExpression(invocationExpression *ast.InvocationExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitMemberExpression(memberExpression *ast.MemberExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitIndexExpression(indexExpression *ast.IndexExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitConditionalExpression(conditionalExpression *ast.ConditionalExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitUnaryExpression(unaryExpression *ast.UnaryExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitBinaryExpression(binaryExpression *ast.BinaryExpression) ast.Repr {
	binaryExpression.Left.Accept(c)
	binaryExpression.Right.Accept(c)

	switch binaryExpression.Operation {
	case ast.OperationPlus:
		c.emit(instructions.IAddIns)
	case ast.OperationNotEqual:
		c.emit(instructions.NotEqualIns)
	default:
		panic("Unsupported")
	}

	return nil
}

func (c *CodeGenerator) VisitFunctionExpression(functionExpression *ast.FunctionExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitStringExpression(stringExpression *ast.StringExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitCastingExpression(castingExpression *ast.CastingExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitCreateExpression(createExpression *ast.CreateExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitDestroyExpression(destroyExpression *ast.DestroyExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitReferenceExpression(referenceExpression *ast.ReferenceExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitForceExpression(forceExpression *ast.ForceExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitPathExpression(pathExpression *ast.PathExpression) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitProgram(program *ast.Program) ast.Repr {
	for _, decl := range program.Declarations() {
		decl.Accept(c)
	}

	return nil
}

func (c *CodeGenerator) VisitFunctionDeclaration(funcDecl *ast.FunctionDeclaration) ast.Repr {
	// start new scope
	parentScope := c.scope
	c.scope = NewScope()

	c.emit(instructions.NewFunction(funcDecl.Identifier.Identifier))

	// TODO: define params

	funcDecl.FunctionBlock.Accept(c)

	// restore the scope
	c.scope = parentScope
	return nil
}

func (c *CodeGenerator) VisitBlock(block *ast.Block) ast.Repr {
	for _, statement := range block.Statements {
		statement.Accept(c)
	}

	return nil
}

func (c *CodeGenerator) VisitFunctionBlock(funcBlock *ast.FunctionBlock) ast.Repr {
	// TODO: pre conditions, post conditions
	funcBlock.Block.Accept(c)
	return nil
}

func (c *CodeGenerator) VisitCompositeDeclaration(declaration *ast.CompositeDeclaration) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitInterfaceDeclaration(declaration *ast.InterfaceDeclaration) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitFieldDeclaration(declaration *ast.FieldDeclaration) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitEnumCaseDeclaration(declaration *ast.EnumCaseDeclaration) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitPragmaDeclaration(declaration *ast.PragmaDeclaration) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitImportDeclaration(declaration *ast.ImportDeclaration) ast.Repr {
	panic("implement me")
}

func (c *CodeGenerator) VisitTransactionDeclaration(declaration *ast.TransactionDeclaration) ast.Repr {
	panic("implement me")
}
