/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		containerInfo := getFirstContainer()
		dbPassword := containerInfo["environment"].(map[string]interface{})["MARIADB_ROOT_PASSWORD"].(string)
		containerPort := containerInfo["ports"].([]interface{})[0].(map[string]interface{})["published"].(string)
		dbUser := "root"

		if len(args) != 1 {
			fmt.Println("Usage: delete <database_name>")
			return
		}
		dbName := args[0]

		db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp(127.0.0.1:"+containerPort+")/")
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
		if err != nil {
			panic(err)
		}

		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
