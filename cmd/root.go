package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tahsinrahman/booklist-api/api"
)

// NewRootCmd is the root command for our cli app
func NewRootCmd() *cobra.Command {
	port := 8080

	rootCmd := &cobra.Command{
		Use:     "booklist-api",
		Short:   "A simple booklist api",
		Long:    "A simple booklist api that supports CRUD operations",
		Example: "booklist-api --port=8080",
		Run: func(cmd *cobra.Command, args []string) {
			api.StartServer(port)
		},
	}

	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "specify the port in which the server will run")

	return rootCmd
}
