package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "git-tool",
    Short: "Git tool for managing multiple repositories",
    Long:  `A tool for cloning, pulling, replacing content, and pushing changes in multiple Git repositories defined in a YAML file.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.PersistentFlags().StringP("file", "f", "", "YAML file with repository definitions")
    rootCmd.PersistentFlags().StringP("branch", "b", "main", "Branch to use for operations")
    rootCmd.MarkPersistentFlagRequired("file")
}
