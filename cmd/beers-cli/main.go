package main

import (
	"flag"

	"parsing_http_response/internal/storage/ontario"

	"parsing_http_response/internal/storage/csv"

	beerscli "parsing_http_response/internal"
	"parsing_http_response/internal/cli"

	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = ontario.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	rootCmd.Execute()
}
