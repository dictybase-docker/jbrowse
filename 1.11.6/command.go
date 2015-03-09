package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
	"gopkg.in/codegangsta/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "jbrowse"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "A http static file server for jbrowse",
			Action: ServeAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "jbrowse-folder, jf",
					Usage:  "Location of jbrowse source folder",
					EnvVar: "JBROWSE_FOLDER",
				},
				cli.StringFlag{
					Name:  "port, p",
					Usage: "http port, default is 9595",
					Value: "9595",
				},
				cli.StringFlag{
					Name:  "log-folder, f",
					Usage: "Folder where the log file will be kept",
					Value: "/log/jbrowse",
				},
				cli.StringFlag{
					Name:  "log-file, l",
					Usage: "Name of the log file, the logging is done in apache combined log format",
					Value: "frontend.log",
				},
				cli.BoolFlag{
					Name:  "no-stderr",
					Usage: "Turn off stderr logging, on by default",
				},
			},
		},
	}
	app.Run(os.Args)
}

func ServeAction(c *cli.Context) {
	// create log folder
	err := os.MkdirAll(c.String("log-folder"), 0744)
	if err != nil {
		log.Fatal(err)
	}
	var w io.Writer
	fw, err := os.Create(filepath.Join(c.String("log-folder"), c.String("log-file")))
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()
	if c.Bool("no-stderr") {
		w = fw
	} else {
		w = io.MultiWriter(fw, os.Stderr)
	}
	fs := http.FileServer(http.Dir(c.String("jbrowse-folder")))
	http.Handle("/", handlers.CombinedLoggingHandler(w, fs))
	port := fmt.Sprintf(":%s", c.String("port"))
	log.Printf("listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
