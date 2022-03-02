package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"yofio/avanzado/initdb"
)

func main() {
	ct := flag.Bool("c", false, "Create tabla")
	dt := flag.Bool("d", false, "Delete tabla")
	flag.Parse()
	ctx := context.TODO()
	cfg, err := initdb.CreateConfig(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	if *ct {
		if err := initdb.CreateTable(ctx, cfg); err != nil {
			fmt.Println(err.Error())
		}
	}
	if *dt {
		if err := initdb.DeleteTable(ctx, cfg); err != nil {
			fmt.Println(err.Error())
		}
	}
}
