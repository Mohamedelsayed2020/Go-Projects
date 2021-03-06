package migration

import (
	"Go-Projects/model"
	"fmt"
	"github.com/spf13/cobra"
	"Go-Projects/model/common"
	"log"
	"os"
)
var command = &cobra.Command{}

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "pg-migrate command is a sub command from root command",
	Long:  "pg-migrate command is a sub command from root command used to migrate postgres db changes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {

		}
		models := args[0]
		switch models {
		case "user":
			common.Conn().AutoMigrate(&model.User{})
		default:
			fmt.Println("This model hasn't created yet :(")
		}
		log.Println("Successfully Migrated :)")
		},
}

func init() {
	command.AddCommand(migrate)
}


func Execute() {
	if err := command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}