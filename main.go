package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// TODO: Set up structs for
// Config. Config will read from ~/.config/zet_config
// ZET_DIR
type Config struct {
	Path    string
	ZetHome string
}

func (c *Config) getConfigPath() error {
	userProfile := os.Getenv("USERPROFILE")

	if userProfile == "" {
		return fmt.Errorf("USERPROFILE env is not set: %s", userProfile)
	}

	stdPath := filepath.Join(userProfile, ".config", "zet_config")
	c.Path = stdPath

	return nil
}

func (c *Config) getZetHome() error {

	err := c.getConfigPath()

	if err != nil {
		return fmt.Errorf("Error getting config path: %v", err)
	}

	stdPath := c.Path
	f, err := os.Open(stdPath)
	defer f.Close()
	if err != nil {
		return err
	}

	scan := bufio.NewScanner(f)

	for scan.Scan() {
		line := scan.Text()

		//check if line starts with ZET_HOME
		if strings.HasPrefix(line, "ZET_HOME") {
			parts := strings.Split(line, "=")
			partNoSpace := strings.TrimSpace(parts[1])
			partNoQuote := strings.Trim(partNoSpace, `"`)

			c.ZetHome = convertPath(partNoQuote)
			return nil
		}
	}
	return nil
}

func convertPath(path string) string {
	if strings.HasPrefix(path, "~") {
		userProfile := os.Getenv("USERPROFILE")
		if userProfile != "" {
			// Normalize the path for both \ and / after ~
			if suffix, found := strings.CutPrefix(path, "~\\"); found {
				return filepath.Join(userProfile, suffix)
			} else if suffix, found := strings.CutPrefix(path, "~/"); found {
				return filepath.Join(userProfile, suffix)
			} else if path == "~" {
				return userProfile // If the path is exactly "~"
			}
		}

	}
	return path
}

// Change ZET_HOME = "path" in ~/.config/zet_config to path
func (c *Config) changeZetHome(path string) error {
	f, err := os.Open(c.Path)

	defer f.Close()

	if err != nil {
		return err
	}

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {

		line := s.Text()

		if strings.HasPrefix(line, "ZET_HOME") {
			line = fmt.Sprintf(`ZET_HOME = "%s"`, path)
		}
		lines = append(lines, line)
	}

	if err := s.Err(); err != nil {
		return err
	}
	f, err = os.OpenFile(c.Path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func generateFilename(home string) string {
	return fmt.Sprintf("%s/%s.md", home, time.Now().Format("20060102_150405"))
}

func add(home string) {
	filename := generateFilename(home)
	cmd := exec.Command("nvim", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error opening neovim: %s", err)
	}
}

func main() {
	c := Config{}
	err := c.getConfigPath()
	if err != nil {
		fmt.Printf("Error getting config Path")
		os.Exit(1)
	}
	err = c.getZetHome()

	if err != nil {
		fmt.Printf("Error gettting directory: %v", err)
	}

	configCmd := flag.NewFlagSet("config", flag.ExitOnError)
	zetHomeFlag := configCmd.String("home", "", `Change zet Home`)

	switch os.Args[1] {

	case "config":
		configCmd.Parse(os.Args[2:])
		if *zetHomeFlag != "" {
			err := c.changeZetHome(*zetHomeFlag)
			if err != nil {
				fmt.Errorf("Error changing zet home: %v", err)
			}
		}
	case "add":
		add(c.ZetHome)
	}

}
