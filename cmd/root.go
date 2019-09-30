package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/itross/sgul"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sgulreg",
	Short: "SGUL Service Registry",
	Long: `---------------------------------------------
sgul Service Registry v. 0.1.0
---------------------------------------------`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initialize)
}

func initialize() {
	//conf = sgul.GetConfiguration()
	logger = sgul.GetLogger().Sugar()
	//initRegistry()

	env := strings.ToLower(os.Getenv("ENV"))
	if env != "dev" && env != "development" && env != "" {
		configureGentleShutdown()
	}
}

func configureGentleShutdown() {
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		logger.Infof("[%+v] SIGNAL CAUGHT", sig)

		logger.Info("wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)

		// Close BoltDB??????
		//db.Close()
		//logger.Info("DB connection closed")

		logger.Info("service goes down")
		logger.Info("Bye!")
		logger.Sync()
		os.Exit(0)
	}()

	logger.Info("service gentle shutdown hook activated")
}
