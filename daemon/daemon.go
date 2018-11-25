package daemon

import (
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "go-webapp/db"
    "go-webapp/model"
    "go-webapp/ui"
)

type Config struct {
    ListenSpec string
    Db db.Config
    Ui ui.Config
}

func Run(cfg *Config) error {
    log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

    db, err := db.InitDb(cfg.Db)
    if err != nil {
        log.Printf("Error initializing database: %v\n", err)
        return err
    }

    m := model.New(db)

    l, err := net.Listen("tcp", cfg.ListenSpec)
    if err != nil {
        log.Printf("Error creating listener: %v\n", err)
        return err
    }

    ui.Start(cfg.Ui, m, l)

    waitForSignal()

    return nil
}

func waitForSignal() {
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    s := <-ch
    log.Printf("Got signal: %v, exiting.", s)
}