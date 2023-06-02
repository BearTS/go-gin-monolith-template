package commands

import (
	"fmt"

	"github.com/BearTS/go-gin-monolith/database"
	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			begin := dbConnection.Begin()

			for i, migrate := range database.AutoMigrate(begin) {
				if err := migrate.Run(begin); err != nil {
					begin.Rollback()
					fmt.Println("[Migrate] Running raw sql schema creation failed")
					panic(err)
				}
				fmt.Println("[", i, "]: ", "Migrate table: ", migrate.TableName)
			}
			begin.Commit()
			fmt.Println("Migration Completed")
			return nil
		},
	}
}
