package main

import (
	"fmt"
	"log"

	"github.com/hugoivankm/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Unable to read configuration file")
	}
	cfg.SetUser("hugo")
	fmt.Printf("%+v\n", cfg)
}
