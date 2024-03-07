package app

import (
	"github.com/leslieleung/gin-application-template/internal/route"
	"github.com/leslieleung/gin-application-template/pkg/database"
	"github.com/leslieleung/gin-application-template/pkg/log"
	"github.com/leslieleung/mjolnir/config"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Run: runServe,
}

func runServe(cmd *cobra.Command, args []string) {
	log.Init()
	configPath, _ := cmd.Flags().GetString("config")
	config.Init(configPath)
	skipDatabase, _ := cmd.Flags().GetBool("skip-database")
	if !skipDatabase {
		database.Init(config.GetConfig().GetString("database.dsn"))
	}
	route.StartRouter()
}

func init() {
	ServeCmd.Flags().StringP("config", "c", "config/example.config.yaml", "config file path")
	ServeCmd.Flags().BoolP("skip-database", "s", false, "skip database initialization")
}
