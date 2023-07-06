/*
Copyright © 2023 Mahir Labib Chowdhury dev.mahirchy@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/m4hi2/cineTicketAlert/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server for cineTicketAlert",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("server.port")
		Serve(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Serve(port string) {
	app, err := api.New()
	if err != nil {
		log.Fatalln(err)
	}

	addr := fmt.Sprintf(":%s", port)

	err = app.Listen(addr)
	if err != nil {
		log.Fatalln(err)
	}

}
