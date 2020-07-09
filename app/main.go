package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pskclub/tendermint-domain-manament/app/domain"
	"os"
	"os/signal"
	"syscall"

	abciserver "github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/libs/log"
)

var socketAddr string

func init() {
	flag.StringVar(&socketAddr, "socket-addr", "tcp://0.0.0.0:26658", "Unix domain socket address")
}

func main() {
	db, err := gorm.Open("mysql", "my_user:my_password@(db:3306)/my_database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	defer db.Close()

	app := domain.NewApplication(db)
	flag.Parse()

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	server := abciserver.NewSocketServer(socketAddr, app)
	server.SetLogger(logger)
	if err := server.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error starting socket server: %v", err)
		os.Exit(1)
	}
	defer server.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}
