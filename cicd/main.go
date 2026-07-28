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

	var cloudCmd = &cobra.Command{
		Use:   "cloud [arg]",
		Short: "Run cloud",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var arg string
			if len(args) > 0 {
				arg = args[0]
			}
			return w.Pipe().Cloud(arg)
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

	rootCmd.AddCommand(setupCmd, syncCmd, fixCmd, cloudCmd, verifyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v\n\nFAILED\n", err)
		os.Exit(1)
	}
}
