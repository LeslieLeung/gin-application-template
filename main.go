package main

import (
	"github.com/leslieleung/gin-application-template/cmd"
	"github.com/leslieleung/gin-application-template/internal/log"
)

func main() {
	log.Init()
	cmd.Execute()
}
