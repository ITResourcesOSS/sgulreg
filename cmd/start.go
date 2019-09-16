package cmd

import (
	"fmt"

	"github.com/ITResourcesOSS/sgulreg/internal"
	"github.com/ITResourcesOSS/sgulreg/internal/registry"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "starts the API Gateway",
	Long:  "This command configures and starts the API Gateway",
	Run: func(cmd *cobra.Command, args []string) {
		start(args)
	},
}

func init() {
	RootCmd.AddCommand(startCommand)
}

func start(args []string) {
	logger.Info("Starting Service Registry...")
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

	r := registry.New(db)
	internal.NewApp(r).Start()
}
