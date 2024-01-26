/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// crCmd represents the cr command
var indbCmd = &cobra.Command{
	Use:   "indb",
	Short: "Into the database",
	Long: `Into the database. For example:

	docker-dev indb
	docker-dev indb <db_name>`,
	Run: func(cmd *cobra.Command, args []string) {

		containerInfo := getFirstContainer()
		containerName := containerInfo["container_name"].(string)
		dbPassword := containerInfo["environment"].(map[string]interface{})["MARIADB_ROOT_PASSWORD"].(string)
		dbUser := "root"

		finalCmd := exec.Command("docker", "exec", "-it", containerName, "mariadb", "-u"+dbUser, "-p"+dbPassword)

		if len(args) != 0 {
			dbName := args[0]
			finalCmd = exec.Command("docker", "exec", "-it", containerName, "mariadb", "-u"+dbUser, "-p"+dbPassword, dbName)
		}

		finalCmd.Stdin = os.Stdin
		finalCmd.Stdout = os.Stdout
		finalCmd.Stderr = os.Stderr

		finalCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(indbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
