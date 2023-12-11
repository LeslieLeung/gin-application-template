package app

import (
	"strconv"

	"github.com/leslieleung/gin-application-template/internal/config"
	"github.com/leslieleung/gin-application-template/internal/route"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Run: runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	configPath, _ := cmd.Flags().GetString("config")
	cfg := config.New(configPath)
	r := route.RegisterRoute()

	r.Run(":" + strconv.Itoa(cfg.GetInt("app.port")))
}

func init() {
	ServeCmd.Flags().StringP("config", "c", "config/example.config.yaml", "config file path")
}
