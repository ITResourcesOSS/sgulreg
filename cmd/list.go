package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var listServicesCommand = &cobra.Command{
	Use:   "list",
	Short: "list all registered services",
	Long:  "This commando prints out the list of all the registered servcies information",
	Run: func(cmd *cobra.Command, args []string) {
		listServices(args)
	},
}

func init() {
	RootCmd.AddCommand(listServicesCommand)
}

func listServices(args []string) {
	db, err := bolt.Open("registry.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("services"))
		return err
	})
	if err != nil && err.Error() != "bucket already exists" {
		panic(fmt.Errorf("error creating 'services' bucket: %s", err))
	}
	logger.Info("internal service registry database (BoltDB instance) inizialized")

	// serviceRepository := repositories.NewServiceRepository(db)
	// registry := services.NewRegistry(serviceRepository)

}
