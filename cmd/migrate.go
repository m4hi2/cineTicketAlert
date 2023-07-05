/*
Copyright © 2023 Mahir Labib Chowdhury dev.mahirchy@gmail.com
*/
package cmd

import (
	_ "github.com/m4hi2/capsule71/db/ddls"
	_ "github.com/m4hi2/capsule71/db/dmls"
	"github.com/m4hi2/capsule71/pkg/dbconn"
	migratorCmd "github.com/m4hi2/capsule71/pkg/migrator/cmd"
	"gorm.io/gorm"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newMigrateCmd())
}
func newMigrateCmd() *cobra.Command {
	options := &migratorCmd.OptionsArg{
		FnGetDBCallback: func() *gorm.DB {
			return getDBConnection()
		},
	}
	return migratorCmd.NewMigrateCmd(options)
}

func getDBConnection() *gorm.DB {
	db, err := dbconn.NewConnection()
	if err != nil {
		log.Fatal("postgres connect error: ", err)
	}
	return db
}
