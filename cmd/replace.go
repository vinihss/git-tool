package cmd

import (
    "fmt"
    "os/exec"

    "github.com/spf13/cobra"
)

var oldStr, newStr string

var replaceCmd = &cobra.Command{
    Use:   "replace",
    Short: "Replace text in the defined repositories",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        projects := readProjects(file)

        for _, project := range projects {
            fmt.Printf("Replacing text in project: %s\n", project.Name)
            replaceInFiles(project.Name, oldStr, newStr)
        }
    },
}

func init() {
    rootCmd.AddCommand(replaceCmd)
    replaceCmd.Flags().StringVar(&oldStr, "old", "", "Old string to be replaced")
    replaceCmd.Flags().StringVar(&newStr, "new", "", "New string to replace with")
    replaceCmd.MarkFlagRequired("old")
    replaceCmd.MarkFlagRequired("new")
}

func replaceInFiles(dir, old, new string) {
    cmd := exec.Command("find", dir, "-type", "f", "-exec", "sed", "-i", "", fmt.Sprintf("s/%s/%s/g", old, new), "{}", "+")
    runCommand(cmd, "Replacing text in files")
}
