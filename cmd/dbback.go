/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

// dbbackCmd represents the dbback command
var dbbackCmd = &cobra.Command{
	Use:   "dbback",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/** Get parameter from .env **/
		currentPath, _ := os.Getwd()

		if len(args) != 1 {
			fmt.Println("Usage: dbback <database_name>")
			return
		}
		dbName := args[0]

		// Get current time
		currentTime := time.Now()
		timeStamp := currentTime.Format("20060102150405")
		backupFilePath := currentPath + "/db-backup/" + dbName + timeStamp + ".sql"

		finalCmd := exec.Command(
			"docker",
			"exec",
			CONTAINER_NAME,
			MYSQL_EXECUTE_NAME + "-dump",
			"-u"+DBUSER,
			"-p"+MYSQL_ROOT_PASSWORD,
			"--max-allowed-packet=9120M",
			dbName,
		)

		// Redirect standard output to the file
		outfile, err := os.Create(backupFilePath)
		if err != nil {
			fmt.Println("Error creating backup file:", err)
			return
		}
		defer outfile.Close()
		finalCmd.Stdout = outfile

		// Redirect standard error to the console for error visibility
		finalCmd.Stderr = os.Stderr

		// Run the command
		err = finalCmd.Run()
		if err != nil {
			fmt.Println("Error running "+MYSQL_EXECUTE_NAME+"-dump:", err)
			return
		}

		fmt.Printf("Backup of database %s completed. Backup saved to %s\n", dbName, backupFilePath)
	},
}

func init() {
	rootCmd.AddCommand(dbbackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbbackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbbackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
