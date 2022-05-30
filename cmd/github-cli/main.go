package main

import (
	"github.com/spf13/cobra"
	"parsing_http_response/internal"
	"parsing_http_response/internal/cli"
	"parsing_http_response/internal/storage/csv"
	"parsing_http_response/internal/storage/github"
)

func main() {

	var remoteRepository internal.UserRepo
	remoteRepository = github.NewRemoteRepository()

	var csvRepository internal.UserRepo
	csvRepository = csv.NewCsvRepository()

	rootCmd := &cobra.Command{Use: "github-cli"}
	rootCmd.AddCommand(cli.InitUserCmd(remoteRepository, csvRepository))
	rootCmd.Execute()
}
