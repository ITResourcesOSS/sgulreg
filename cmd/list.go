package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	reg "github.com/itross/sgul/registry"
	"github.com/itross/sgulreg/internal/repositories"
	"github.com/itross/sgulreg/internal/services"
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

	serviceRepository := repositories.NewServiceRepository(db)
	registry := services.NewRegistry(serviceRepository)

	var instances []reg.ServiceInfoResponse
	instances, err = registry.DiscoverAll(nil)
	if err != nil {
		panic(err)
	}

	show(instances)
}

func show(instances []reg.ServiceInfoResponse) {
	fmt.Println("\nSgulREG Registered service instances:")
	fmt.Println("=====================================")
	for _, i := range instances {
		fmt.Printf("\n* \"%s\"\n", i.Name)
		for _, j := range i.Instances {
			fmt.Printf("\t- instance id: \t%s\n", j.InstanceID)
			fmt.Printf("\t  host: \t\t%s\n", j.Host)
			fmt.Printf("\t  scheme: \t\t%s\n", j.Schema)
			fmt.Printf("\t  info url: \t\t%s\n", j.InfoURL)
			fmt.Printf("\t  health url: \t\t%s\n", j.HealthCheckURL)
			fmt.Printf("\t  registration: \t%s\n", j.RegistrationTimestamp)
			fmt.Printf("\t  last refresh: \t%s\n", j.LastRefreshTimestamp)
			fmt.Println("\t------------------------------------------------------------")
		}
	}
}
