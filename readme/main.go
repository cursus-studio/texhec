package main

import (
	"fmt"
	"os"
	docspkg "readme/docs/pkg"
)

func main() {
	s := docspkg.NewService()
	errs := []error{}
	if err := s.GenerateModuleDocs("docs/"); err != nil {
		errs = append(errs, err)
	}
	if errors := s.GenerateModulesDocs("../engine/modules"); errors != nil {
		errs = append(errs, errors...)
	}
	if errors := s.GenerateModulesDocs("../core/modules"); errors != nil {
		errs = append(errs, errors...)
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
