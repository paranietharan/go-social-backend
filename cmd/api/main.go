package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}
	app := &application{
		cfg,
	}

	log.Printf("Server has started at %s\n", app.config.addr)

	mux := app.mount()
	log.Fatal(app.run(mux))
}
