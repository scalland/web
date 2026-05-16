package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	"/internal/handlers"
	"/internal/routes"
	"/internal/worker"
	"/pkg/utils"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := utils.GetUtils()

		// Connect databases
		u.ConnectWriter()
		u.ConnectReader()
		defer u.CloseDB()

		// Start email worker
		w := worker.NewWorker(10)
		go w.Start()

		// Setup handlers
		h := handlers.NewWebHandlers(u, w)

		// Setup router
		router := routes.SetupRouter(h, u, webFS)

		addr := fmt.Sprintf(":%s", u.Config.App.ServerPort)
		u.Logger.Infof("Starting server on %s", addr)

		srv := &http.Server{
			Addr:         addr,
			Handler:      router,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
