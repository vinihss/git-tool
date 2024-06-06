package cmd

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
    Use:   "push",
    Short: "Push changes to the remote repositories",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        branch, _ := cmd.Flags().GetString("branch")
        projects := readProjects(file)

        for _, project := range projects {
            fmt.Printf("Pushing project: %s\n", project.Name)
            pushRepo(project.Name, branch)
        }
    },
}

func init() {
    rootCmd.AddCommand(pushCmd)
}

func pushRepo(dir, branch string) {
    cmd := exec.Command("git", "-C", dir, "push", "origin", branch)
    runCommand(cmd, "Pushing changes")
}
