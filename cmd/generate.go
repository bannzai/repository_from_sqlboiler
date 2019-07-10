/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/bannzai/repository_from_sqlboiler/pkg/formatter"
	"github.com/bannzai/repository_from_sqlboiler/pkg/generator"
	"github.com/bannzai/repository_from_sqlboiler/pkg/parser"
	"github.com/bannzai/repository_from_sqlboiler/pkg/reader"
	"github.com/bannzai/repository_from_sqlboiler/pkg/writer"
	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	templateFilePath    string
	sourceFilePath      string
	destinationFilePath string
}

var generateOptions = GenerateOptions{}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		source := generateOptions.sourceFilePath
		destination := generateOptions.destinationFilePath
	tmpl:
		-generateOptions.templateFilePath
		g := generator.GoCodeGenerator{
			DestinationFilePath: destination,
			EntityParser: parser.Entity{
				FilePath: source,
			},
			GoFormatter: formatter.GoFormatter{
				FilePath: destination,
			},
			TemplateReader: reader.Template{
				FilePath: tmpl,
			},
			Writer: writer.File{
				FilePath: destination,
			},
		}
		g.Generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	const invalidGenerateValue = ""
	generateCmd.Flags().StringVar(&generateOptions.sourceFilePath, "source", invalidGenerateValue, "source entity file for generate repository.")
	generateCmd.Flags().StringVar(&generateOptions.destinationFilePath, "destination", invalidGenerateValue, "destination repository file path.")
	generateCmd.Flags().StringVar(&generateOptions.templateFilePath, "template", invalidGenerateValue, "template file path for generate repository source code.")

}
