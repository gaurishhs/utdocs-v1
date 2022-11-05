package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gaurishhs/utdocs/pkg/config"
	"github.com/gaurishhs/utdocs/pkg/utparser"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the site",
	Long: `Build the site.
This command will build the site and output it to the build directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		parser := utparser.NewParser()
		if len(args) == 0 {
			config.ReadConfig("config.json")
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filepath.WalkDir(dir+"/content", func(path string, d os.DirEntry, err error) error {
				// Generating the Sidebar
				filePath := strings.TrimPrefix(path, dir+"/content/")
				fmt.Print(filePath)
				if filepath.Ext(path) == ".md" {
					file, err := os.Open(path)
					if err != nil {
						fmt.Println(err)
					}
					defer file.Close()
					os.Mkdir(dir+"/build", 0770)
					os.MkdirAll(dir+"/build/"+strings.TrimSuffix(strings.TrimPrefix(path, dir+"/content/"), d.Name()), 0770)
					f, err := os.Create(dir + "/build/" + strings.TrimSuffix(strings.TrimPrefix(path, dir+"/content/"), ".md") + ".html")
					if err != nil {
						fmt.Println(err)
					}
					defer f.Close()
					writer := io.Writer(f)
					utparser.ParseFile(parser, path, nil, writer)
				}
				return nil
			})
			file, err := os.ReadFile(dir + "/styles.css")
			if err != nil {
				fmt.Println(err)
			}
			os.WriteFile(dir+"/build/styles.css", file, 0770)
		} else {
			config.ReadConfig(args[0] + "/config.json")
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			filepath.WalkDir(dir+"/"+args[0]+"/content", func(path string, d os.DirEntry, err error) error {
				if filepath.Ext(path) == ".md" {
					file, err := os.Open(path)
					if err != nil {
						fmt.Println(err)
					}
					defer file.Close()
					os.Mkdir(dir+"/"+args[0]+"/build", 0770)
					os.MkdirAll(dir+"/"+args[0]+"/build/"+strings.TrimSuffix(strings.TrimPrefix(path, dir+"/"+args[0]+"/content/"), d.Name()), 0770)
					f, err := os.Create(dir + "/" + args[0] + "/build/" + strings.TrimSuffix(strings.TrimPrefix(path, dir+"/"+args[0]+"/content/"), ".md") + ".html")
					// f, err := os.Create(dir + "/" + args[0] + "/build/" + strings.TrimSuffix(d.Name(), "md") + "html")
					if err != nil {
						fmt.Println(err)
					}
					defer f.Close()
					writer := io.Writer(f)
					utparser.ParseFile(parser, path, nil, writer)
				}
				return nil
			})
			file, err := os.ReadFile(dir + "/" + args[0] + "/styles.css")
			if err != nil {
				fmt.Println(err)
			}
			os.WriteFile(dir+"/"+args[0]+"/build/styles.css", file, 0770)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
