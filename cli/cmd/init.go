package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long: `Initialize a new project in the current directory.
This command will create a new directory with the name of the project.
It will also create a new git repository and initialize it with a README.md file.`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime := time.Now()
		if len(args) == 0 {
			fmt.Printf("Creating a new UTDocs project in the current directory\n")
			configFile, err := os.Create("config.json")
			if err != nil {
				log.Fatal("Failed to create config file", err)
				os.Exit(1)
			}
			defer configFile.Close()
			configFile.WriteString(`{"$schema": "https://api.npoint.io/f041af3a8d6befad626f", "sitename": "My Site" }`)
			configFile.Close()
			contentErr := os.Mkdir("content", 0755)
			if contentErr != nil {
				log.Fatal("Failed to create content directory", err)
				os.Exit(1)
			}
			resp, err := http.Get("https://raw.githubusercontent.com/gaurishhs/utdocs/main/styles.css")
			if err != nil {
				log.Fatal("Failed to download styles.css", err)
				os.Exit(1)
			}
			defer resp.Body.Close()
			stylesFile, err := os.Create("styles.css")
			if err != nil {
				log.Fatal("Failed to create styles.css", err)
				os.Exit(1)
			}
			_, err = io.Copy(stylesFile, resp.Body)
			if err != nil {
				log.Fatal("Failed to copy styles.css", err)
				os.Exit(1)
			}
			stylesFile.Close()
			log.Printf("Project created successfully, Time taken: %s\n", time.Since(startTime).String())
		} else {
			fmt.Printf("Creating a new UTDocs project in the specified directory\n")
			err := os.Mkdir(args[0], 0755)
			if err != nil {
				log.Fatal("Failed to create project directory", err)
				os.Exit(1)
			}
			os.Chdir(args[0])
			configFile, err := os.Create("config.json")
			if err != nil {
				log.Fatal("Failed to create config file\n", err)
				os.Exit(1)
			}
			configFile.WriteString(`{"$schema": "https://api.npoint.io/f041af3a8d6befad626f", "sitename": "My Site" }`)
			configFile.Close()
			contentErr := os.Mkdir("content", 0755)
			if contentErr != nil {
				log.Fatal("Failed to create content directory", err)
				os.Exit(1)
			}
			resp, err := http.Get("https://raw.githubusercontent.com/gaurishhs/utdocs/main/styles.css")
			if err != nil {
				log.Fatal("Failed to download styles.css", err)
				os.Exit(1)
			}
			defer resp.Body.Close()
			stylesFile, err := os.Create("styles.css")
			if err != nil {
				log.Fatal("Failed to create styles.css", err)
				os.Exit(1)
			}
			_, err = io.Copy(stylesFile, resp.Body)
			if err != nil {
				log.Fatal("Failed to copy styles.css", err)
				os.Exit(1)
			}
			stylesFile.Close()
			log.Printf("Project created successfully, Time taken: %s\n", time.Since(startTime).String())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
