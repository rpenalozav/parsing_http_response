package cli

import (
	"github.com/spf13/cobra"
	. "parsing_http_response/internal"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const IUser string = "login"

// initialize users command
func InitUserCmd(RemoteRepository UserRepo, CsvRepository UserRepo) *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "github",
		Short: "Print data about github user",
		Run:   runUserFn(RemoteRepository, CsvRepository),
	}

	userCmd.Flags().StringP(IUser, "u", "", "user in github")
	return userCmd
}

func runUserFn(remoteRepository UserRepo, csvRepository UserRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString(IUser)
		res, err := remoteRepository.GetUser(user)
		if err != nil {
			panic("user not exist")
		}
		err = csvRepository.AddUser(res)
	}
}
