package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/config"
	"github.com/yjwong/lark-cli/internal/output"
)

// Version information - set via ldflags at build time
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// SetVersionInfo allows setting version info from main package
func SetVersionInfo(v, c, d string) {
	version = v
	commit = c
	date = d
}

var rootCmd = &cobra.Command{
	Use:   "lark",
	Short: "Lark CLI for Claude Code",
	Long: `A CLI tool to interact with Lark APIs.
Designed for use by Claude Code with JSON output.

All commands output JSON by default.`,
	SilenceUsage:  true,
	SilenceErrors: true,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("lark %s\n", version)
		fmt.Printf("  commit: %s\n", commit)
		fmt.Printf("  built:  %s\n", date)
	},
}

// Execute runs the root command
func Execute() {
	// Initialize config, but don't fail for version command
	if err := config.Init(); err != nil {
		// Allow version command to run without config
		if len(os.Args) >= 2 && os.Args[1] == "version" {
			// Skip config error for version command
		} else {
			output.Fatal("CONFIG_ERROR", err)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		output.Fatal("COMMAND_ERROR", err)
	}
}

func init() {
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(bitableCmd)
	rootCmd.AddCommand(calCmd)
	rootCmd.AddCommand(chatCmd)
	rootCmd.AddCommand(contactCmd)
	rootCmd.AddCommand(docCmd)
	rootCmd.AddCommand(mailCmd)
	rootCmd.AddCommand(minutesCmd)
	rootCmd.AddCommand(msgCmd)
	rootCmd.AddCommand(sheetCmd)
	rootCmd.AddCommand(versionCmd)
}
