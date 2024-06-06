package cmd

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "os/exec"

    "gopkg.in/yaml.v2"

    "github.com/spf13/cobra"
)

type Project struct {
    Name   string `yaml:"name"`
    Repo   string `yaml:"repo"`
    Branch string `yaml:"branch"`
}

var cloneCmd = &cobra.Command{
    Use:   "clone",
    Short: "Clone repositories defined in the YAML file",
    Run: func(cmd *cobra.Command, args []string) {
        file, _ := cmd.Flags().GetString("file")
        branch, _ := cmd.Flags().GetString("branch")
        projects := readProjects(file)

        for _, project := range projects {
            fmt.Printf("Cloning project: %s\n", project.Name)
            cloneRepo(project.Repo, project.Name)
            checkoutBranch(project.Name, branch)
        }
    },
}

func init() {
    rootCmd.AddCommand(cloneCmd)
}

func readProjects(fileName string) []Project {
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }

    var projects []Project
    if err := yaml.Unmarshal(data, &projects); err != nil {
        log.Fatalf("Failed to parse YAML: %v", err)
    }

    return projects
}

func cloneRepo(repo, dir string) {
    cmd := exec.Command("git", "clone", repo, dir)
    runCommand(cmd, "Cloning repository")
}

func checkoutBranch(dir, branch string) {
    cmd := exec.Command("git", "-C", dir, "checkout", branch)
    runCommand(cmd, "Checking out branch")
}

func runCommand(cmd *exec.Cmd, action string) {
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        log.Fatalf("%s failed: %v\n%s", action, err, stderr.String())
    }
    fmt.Printf("%s succeeded:\n%s\n", action, out.String())
}
