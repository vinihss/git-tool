package cmd

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
    Use:   "pull",
    Short: "Pull latest changes for the defined repositories",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        branch, _ := cmd.Flags().GetString("branch")
        projects := readProjects(file)

        for _, project := range projects {
            fmt.Printf("Pulling project: %s\n", project.Name)
            pullRepo(project.Name, branch)
        }
    },
}

func init() {
    rootCmd.AddCommand(pullCmd)
}

func pullRepo(dir, branch string) {
    cmd := exec.Command("git", "-C", dir, "pull", "origin", branch)
    runCommand(cmd, "Pulling latest changes")
}
