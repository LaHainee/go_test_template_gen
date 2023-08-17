package main

import (
	"flag"

	"github.com/LaHainee/go_test_template_gen/internal/presenter"
	"github.com/LaHainee/go_test_template_gen/internal/repository/functions"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast"
	astFile "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/file"
	astFunctions "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/functions"
	astFunctionArguments "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/functions/source/arguments"
	astFunctionReceiver "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/functions/source/receiver"
	astImports "github.com/LaHainee/go_test_template_gen/internal/repository/parse/facade/ast/source/imports"
	"github.com/LaHainee/go_test_template_gen/internal/repository/parse/file"
	"github.com/LaHainee/go_test_template_gen/internal/repository/test/create"
	createSourceImports "github.com/LaHainee/go_test_template_gen/internal/repository/test/create/source/imports"
	createSourcePackageName "github.com/LaHainee/go_test_template_gen/internal/repository/test/create/source/package_name"
	createSourceTest "github.com/LaHainee/go_test_template_gen/internal/repository/test/create/source/test"
	"github.com/LaHainee/go_test_template_gen/internal/usecase/codegen"
	"github.com/LaHainee/go_test_template_gen/internal/usecase/codegen/files"
	"github.com/LaHainee/go_test_template_gen/internal/usecase/codegen/files/by_dirpath"
	"github.com/LaHainee/go_test_template_gen/internal/usecase/codegen/files/by_filepath"
)

func main() {
	path := flag.String("path", "", "Path to file or directory for test template generation")
	flag.Parse()

	functionSourceReceiver := astFunctionReceiver.NewSource()
	functionSourceArguments := astFunctionArguments.NewSource()

	sourceFunctions := astFunctions.NewSource(functionSourceReceiver, functionSourceArguments)
	sourceImports := astImports.NewSource()
	sourceFile := astFile.NewSource()
	astFacade := ast.NewFacade(sourceFunctions, sourceImports, sourceFile)

	fileParser := file.NewParser(astFacade)
	functionsRepo := functions.NewRepository()

	filesByDirpath := by_dirpath.NewSource(fileParser)
	filesByFilepath := by_filepath.NewSource(fileParser)
	filesGetter := files.NewGetter(functionsRepo, filesByDirpath, filesByFilepath)

	testsPresenter := presenter.NewPresenter(presenter.NewFactory())

	// <! Create test files chain of responsibility
	sourceCreateImports := createSourceImports.NewSource(astFacade)
	sourceCreatePackageName := createSourcePackageName.NewSource()
	sourceCreateTest := createSourceTest.NewSource()
	sourceCreatePackageName.SetNext(sourceCreateImports)
	sourceCreateImports.SetNext(sourceCreateTest)
	// Create test files chain of responsibility !>

	testCreateRepository := create.NewRepository(sourceCreatePackageName)

	usecase := codegen.NewUseCase(filesGetter, testsPresenter, testCreateRepository)

	err := usecase.Execute(*path)
	if err != nil {
		panic(err)
	}
}
