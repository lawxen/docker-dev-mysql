/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		containerInfo := getFirstContainer()
		containerName := containerInfo["container_name"]
		dbPassword := containerInfo["password"]
		dbUser := "root"

		if len(args) != 1 {
			fmt.Println("Usage: dbback <database_name>")
			return
		}
		dbName := args[0]

		// Define SQL statement to drop all tables
		dropTablesSQL := fmt.Sprintf("mariadb -u%s -p%s -D %s -e 'SHOW TABLES;' | tail -n +2 | xargs -I{} mariadb -u%s -p%s -D %s -e 'DROP TABLE {}'", dbUser, dbPassword, dbName, dbUser, dbPassword, dbName)

		finalCmd := exec.Command("docker", "exec", "-i", containerName, "bash", "-c", dropTablesSQL)

		finalCmd.Stdin = os.Stdin
		finalCmd.Stdout = os.Stdout
		finalCmd.Stderr = os.Stderr

		if err := finalCmd.Run(); err != nil {
			fmt.Printf("Error running command: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("All tables in database %s have been dropped.\n", dbName)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
