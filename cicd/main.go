package main

import (
	cicdpkg "cicd/pkg"
	"cicd/world"
	"fmt"
	"os"

	"github.com/ogiusek/ioc/v2"
	"github.com/spf13/cobra"
)

func main() {
	c := ioc.NewContainer(cicdpkg.Pkg)
	w := ioc.Get[world.CICDWorld](c)

	// var err error
	// switch os.Args[1] {
	// case "setup":
	// 	err = w.Pipe().Setup()
	// case "sync":
	// 	err = w.Pipe().Sync()
	// case "fix":
	// 	err = w.Pipe().Fix()
	// case "verify":
	// 	var arg string
	// 	if len(os.Args) > 3 {
	// 		arg = os.Args[2]
	// 	}
	// 	err = w.Pipe().Verify(arg)
	// default:
	// 	err = errors.New("uknown command")
	// }
	//
	// if err != nil {
	// 	fmt.Printf("%v\n\n", err.Error())
	// 	fmt.Printf("FAILED\n")
	// 	os.Exit(1)
	// 	return
	// }
	// fmt.Printf("SUCCESS\n")
	var rootCmd = &cobra.Command{
		Use:           "cicd",
		Short:         "CICD tool",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("SUCCESS")
		},
	}

	var setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Run setup",
		RunE: func(cmd *cobra.Command, args []string) error {
			return w.Pipe().Setup()
		},
	}

	var syncCmd = &cobra.Command{
		Use:   "sync",
		Short: "Run sync",
		RunE: func(cmd *cobra.Command, args []string) error {
			return w.Pipe().Sync()
		},
	}

	var fixCmd = &cobra.Command{
		Use:   "fix",
		Short: "Run fix",
		RunE: func(cmd *cobra.Command, args []string) error {
			return w.Pipe().Fix()
		},
	}

	var verifyCmd = &cobra.Command{
		Use:   "verify [arg]",
		Short: "Run verify",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var arg string
			if len(args) > 0 {
				arg = args[0]
			}
			return w.Pipe().Verify(arg)
		},
	}

	rootCmd.AddCommand(setupCmd, syncCmd, fixCmd, verifyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n\nFAILED\n", err)
		os.Exit(1)
	}
}
