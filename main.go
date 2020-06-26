package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/tidepool-org/devices/config"
	"github.com/tidepool-org/devices/repo"
	"github.com/tidepool-org/devices/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	if err := cfg.LoadFromEnv(); err != nil {
		log.Fatalf("could not load service configuration: %v", err)
	}

	devicesCfg := config.NewDevicesConfig()
	if err := devicesCfg.LoadFromFile(cfg.DevicesConfigFilename); err != nil {
		log.Fatalf("could not load devices configuration: %v", err)
	}

	validate := validator.New()
	if err := devicesCfg.Validate(validate); err != nil {
		log.Fatalf("could not validate devices configuration: %v", err)
	}

	params := &server.Params{
		Cfg:   cfg,
		Cgms:  repo.NewCgmsRepo(devicesCfg.Devices),
		Pumps: repo.NewPumpsRepo(devicesCfg.Devices),
	}

	// listen to signals to stop server
	// convert to cancel on context that server listens to
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func(stop chan os.Signal, cancelFunc context.CancelFunc) {
		<-stop
		log.Print("sigint or sigterm received!!!")
		cancelFunc()
	}(stop, cancelFunc)

	if err := server.ServeAndWait(ctx, params); err != nil {
		log.Fatalln(err.Error())
	}
}
