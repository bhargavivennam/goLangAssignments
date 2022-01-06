package main

import (
	DataTypePackage "assignments/VarDeclarationProject/DataTypesModule"
	scope "assignments/VarDeclarationProject/ScopesDeclarationModule"
	VarDeclaration "assignments/VarDeclarationProject/VarDeclarationModule"

	"fmt"
)

func main() {
	VarDeclaration.PrintDeclarationTypes() //calling the function using package name
	scope.ScopeDeclaration()
	scope.SampleFunc()
	fmt.Println("Global level var", scope.GlobalVar) //for global level scope it is packagename.variablename
	DataTypePackage.DataType()
	DataTypePackage.ArithmeticOperation()
}
