package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RepoURL  string `yaml:"repo_url"`
	CloneDir string `yaml:"clone_dir"`
}

func main() {
	// Handle Flag for accepting config.yaml
	configFile := flag.String("c", "config.yaml", "path to config YAML file")
	flag.Parse()
	// Handle Config File
	data, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshaling YAML:", err)
		return
	}

	// Runtime variables
	repoURL := config.RepoURL
	cloneDir := config.CloneDir

	fmt.Println("Repo URL:", repoURL)
	fmt.Println("Clone Dir:", cloneDir)

	// Clone the repository
	err = cloneRepo(repoURL, cloneDir)
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		return
	}

	// Change to the cloned directory
	fmt.Println("Entering the release directory")
	err = os.Chdir(cloneDir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Run commands
	listFiles()

	fmt.Println("Deployment successful")
}

func cloneRepo(repoURL, cloneDir string) error {
	dirName, err := exec.Command("date", "+%Y-%m-%d_%H-%M-%S").Output()
	if err != nil {
		log.Fatal(err)
	}
	relDir := string(dirName)
	cmd := exec.Command("git", "clone", repoURL, cloneDir+relDir)
	return cmd.Run()
}

func listFiles() string {
	out, err := exec.Command("ls", "-la").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
	return ""
}

func listProcess() string {
	out, err := exec.Command("ps", "-ef").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(out))
	return ""
}
