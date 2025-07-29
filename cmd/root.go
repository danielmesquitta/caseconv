/*
Copyright © 2025 Daniel Mesquita <daniel.mesquita@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ettle/strcase"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "caseconv",
	Short: "Convert text to different cases",
	Long: `A CLI tool to convert text to different cases using various conversion functions.
	
Examples:
  caseconv camel "hello world"
  caseconv snake "HelloWorld" 
  echo "hello world" | caseconv pascal`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Helper function to get input text from args or stdin
func getInputText(cmd *cobra.Command, args []string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	// Check if there's input from stdin
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return strings.Join(lines, "\n"), nil
	}

	return "", fmt.Errorf("no input provided. Please provide text as an argument or pipe it via stdin")
}

func init() {
	rootCmd.AddCommand(camelCmd)
	rootCmd.AddCommand(goCamelCmd)
	rootCmd.AddCommand(pascalCmd)
	rootCmd.AddCommand(goPascalCmd)
	rootCmd.AddCommand(snakeCmd)
	rootCmd.AddCommand(upperSnakeCmd)
	rootCmd.AddCommand(kebabCmd)
	rootCmd.AddCommand(upperKebabCmd)
}

// Camel case command
var camelCmd = &cobra.Command{
	Use:   "camel [text]",
	Short: "Convert text to camelCase",
	Long:  "Convert text to camelCase (e.g., 'hello world' -> 'helloWorld')",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToCamel(text))
		return nil
	},
}

// Go Camel case command
var goCamelCmd = &cobra.Command{
	Use:   "gocamel [text]",
	Short: "Convert text to Go-style camelCase",
	Long:  "Convert text to Go-style camelCase with special handling for Go identifiers",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToGoCamel(text))
		return nil
	},
}

// Pascal case command
var pascalCmd = &cobra.Command{
	Use:   "pascal [text]",
	Short: "Convert text to PascalCase",
	Long:  "Convert text to PascalCase (e.g., 'hello world' -> 'HelloWorld')",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToPascal(text))
		return nil
	},
}

// Go Pascal case command
var goPascalCmd = &cobra.Command{
	Use:   "gopascal [text]",
	Short: "Convert text to Go-style PascalCase",
	Long:  "Convert text to Go-style PascalCase with special handling for Go identifiers",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToGoPascal(text))
		return nil
	},
}

// Snake case command
var snakeCmd = &cobra.Command{
	Use:   "snake [text]",
	Short: "Convert text to snake_case",
	Long:  "Convert text to snake_case (e.g., 'hello world' -> 'hello_world')",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToSnake(text))
		return nil
	},
}

// Kebab case command
var kebabCmd = &cobra.Command{
	Use:   "kebab [text]",
	Short: "Convert text to kebab-case",
	Long:  "Convert text to kebab-case (e.g., 'hello world' -> 'hello-world')",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToKebab(text))
		return nil
	},
}

// Upper Snake case command (SCREAMING_SNAKE_CASE)
var upperSnakeCmd = &cobra.Command{
	Use:   "uppersnake [text]",
	Short: "Convert text to UPPER_SNAKE_CASE",
	Long:  "Convert text to UPPER_SNAKE_CASE (e.g., 'hello world' -> 'HELLO_WORLD')",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToSNAKE(text))
		return nil
	},
}

// Upper Kebab case command (SCREAMING-KEBAB-CASE)
var upperKebabCmd = &cobra.Command{
	Use:   "upperkebab [text]",
	Short: "Convert text to UPPER-KEBAB-CASE",
	Long:  "Convert text to UPPER-KEBAB-CASE (e.g., 'hello world' -> 'HELLO-WORLD')",
	RunE: func(cmd *cobra.Command, args []string) error {
		text, err := getInputText(cmd, args)
		if err != nil {
			return err
		}
		fmt.Println(strcase.ToKEBAB(text))
		return nil
	},
}
