package olly

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAPIKey() string {
    // First, check for environment variable
    apiKey := os.Getenv("OLLY_API_KEY")
    if apiKey != "" {
        return apiKey
    }

    // If not found, check for .olly file in $HOME
    home, err := os.UserHomeDir()
    if err != nil {
        return ""
    }

    configPath := filepath.Join(home, ".olly")
    content, err := os.ReadFile(configPath)
    if err != nil {
        return ""
    }

    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "OLLY_API_KEY=") {
            return strings.TrimPrefix(line, "OLLY_API_KEY=")
        }
    }

    return ""
}
