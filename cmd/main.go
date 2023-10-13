package main

import (
	"context"
	"log"

	"github.com/bitcodr/re-test/internal/domain/service/shipment"
	"github.com/bitcodr/re-test/internal/infrastructure/config"
	shipmentrepo "github.com/bitcodr/re-test/internal/infrastructure/repository/memory/packet"
	"github.com/bitcodr/re-test/internal/interfaces/rest"
)

/*
main to know about how to run a project, you can check README.md
The structure of project is Hexagonal
- interface -> that is rest and framework layer
- service -> it is domain model that contains all relation business models and logic
- repository -> that contains the source of data we have
*/
func main() {
	ctx := context.Background()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	repo, err := shipmentrepo.InitRepo(ctx, cfg.Connections[config.MEMORY])
	if err != nil {
		log.Fatalln(err.Error())
	}

	shipmentService := shipment.InitService(ctx, repo)

	err = rest.InitTransport(&rest.Rest{
		ShipmentService: shipmentService,
	}, &cfg.Service)
	if err != nil {
		return
	}

}
