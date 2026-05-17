package main

import (
	cicdpkg "cicd/pkg"
	"cicd/world"
	"fmt"
	"os"

	"github.com/ogiusek/ioc/v2"
)

func main() {
	// initialize
	c := ioc.NewContainer(cicdpkg.Pkg)
	world := ioc.Get[world.CICDWorld](c)
	errs := []error{}

	// run command
	var command string
	if len(os.Args) >= 2 {
		command = os.Args[1]
	}

	modules := []string{"cicd/modules", "engine/modules", "core/modules"}

	switch command {
	case "setup":
		if err := world.Hooks().Setup(); err != nil {
			errs = append(errs, err)
		}
	case "hook":
		if err := world.Hooks().Handle(os.Args[2]); err != nil {
			errs = append(errs, err)
		}
	case "generate-all":
		for _, module := range modules {
			if errors := world.Docs().GenerateModulesDocs(module); errors != nil {
				errs = append(errs, errors...)
			}
		}
	case "generate-diff":
		modules, err := world.Git().DiffUncommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := world.Docs().GenerateModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "generate-commit":
		modules, err := world.Git().DiffCommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := world.Docs().GenerateModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "generate-module":
		modules := os.Args[2:]
		for _, module := range modules {
			if err := world.Docs().GenerateModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "verify-all":
		for _, module := range modules {
			if errors := world.Docs().DiffModulesDocs(module); errors != nil {
				errs = append(errs, errors...)
			}
		}
	case "verify-diff":
		modules, err := world.Git().DiffCommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := world.Docs().DiffModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "verify-commit":
		modules, err := world.Git().DiffCommited()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := world.Docs().DiffModuleDocs(module); err != nil {
				errs = append(errs, err)
			}
		}
	case "verify-module":
		modules := os.Args[2:]
		for _, module := range modules {
			if err := world.Docs().DiffModuleDocs(module); err != nil {
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
