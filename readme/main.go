package main

import (
	"fmt"
	"os"
	diffpkg "readme/modules/diff/pkg"
	docspkg "readme/modules/docs/pkg"
)

// lychee
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

	switch command {
	case "all":
		if errors := docs.GenerateModulesDocs("modules"); errors != nil {
			errs = append(errs, errors...)
		}
		if errors := docs.GenerateModulesDocs("../engine/modules"); errors != nil {
			errs = append(errs, errors...)
		}
		if errors := docs.GenerateModulesDocs("../core/modules"); errors != nil {
			errs = append(errs, errors...)
		}
	case "diff":
		modules, err := diff.GetModifiedModules()
		if err != nil {
			errs = append(errs, err)
		}
		for _, module := range modules {
			if err := docs.GenerateModuleDocs(module); err != nil {
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
