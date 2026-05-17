package main

import (
	"fmt"
	"os"
	diffpkg "readme/modules/diff/pkg"
	docspkg "readme/modules/docs/pkg"
)

func main() {
	// initialize
	docs := docspkg.NewService()
	diff := diffpkg.NewService()
	errs := []error{}

	// run command
	var command string
	if len(os.Args) >= 2 {
		command = os.Args[1]
	}

	modules := []string{"ci-cd/generation/readme/modules", "engine/modules", "core/modules"}

	switch command {
	case "generate-all":
		for _, module := range modules {
			if errors := docs.GenerateModulesDocs(module); errors != nil {
				errs = append(errs, errors...)
			}
		}
	case "generate-diff":
		modules, err := diff.DiffUncommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := docs.GenerateModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "generate-commit":
		modules, err := diff.DiffCommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := docs.GenerateModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "generate-module":
		modules := os.Args[2:]
		for _, module := range modules {
			if err := docs.GenerateModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "verify-all":
		for _, module := range modules {
			if errors := docs.DiffModulesDocs(module); errors != nil {
				errs = append(errs, errors...)
			}
		}
	case "verify-diff":
		modules, err := diff.DiffCommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := docs.DiffModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "verify-commit":
		modules, err := diff.DiffCommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := docs.DiffModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "verify-module":
		modules := os.Args[2:]
		for _, module := range modules {
			if err := docs.DiffModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "":
		errs = append(errs, fmt.Errorf("missing command"))
	default:
		errs = append(errs, fmt.Errorf("uknown command \"%v\"", command))
	}

	if len(errs) == 0 {
		fmt.Printf("SUCCESS\n")
		return
	}

	for _, err := range errs {
		fmt.Printf("%v\n\n", err.Error())
	}
	fmt.Printf("FAILED\n")
	os.Exit(1)
}
