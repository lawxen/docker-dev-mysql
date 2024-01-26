/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	if len(args) != 2 {
		fmt.Println("Usage: dbimport <database_name> <backup_file_path>")
		return
	}

	dbName := args[0]
	sqlFilePath := args[1]

	// Use mysql to import db
	finalCmd := exec.Command(
		"docker",
		"exec",
		"-i",
		CONTAINER_NAME,
		MYSQL_EXECUTE_NAME,
		"-u"+DBUSER,
		"-p"+MYSQL_ROOT_PASSWORD,
		dbName,
		// "<",
		// sqlFilePath,
	)

	// Open the backup file
	input, err := os.Open(sqlFilePath)
	if err != nil {
		fmt.Printf("Error opening SQL file: %v\n", err)
		return
	}
	defer input.Close()

	// Set the input of the command to the content of the backup file
	finalCmd.Stdin = input

	// Capture the output and error of the command
	output, err := finalCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
	}

	// Print the output of the command
	fmt.Println(strings.TrimSpace(string(output)))

	fmt.Printf("Import database %s completed.\n", dbName)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
