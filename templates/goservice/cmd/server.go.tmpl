package cmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"{{ .Module }}/apigen"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start your server",
	Long:  `Start your server`,
	Run:   StartServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

type server struct{}

func (s server) GetAuthors(ctx echo.Context) error {
	ctx.JSON(http.StatusOK, &[]apigen.Author{
		{
			Id: "Test Id",
		},
	})

	return nil
}

func StartServer(cmd *cobra.Command, args []string) {
	e := echo.New()
	apigen.RegisterHandlers(e, server{})
	e.Logger.Fatal(e.Start("localhost:3000"))
}
