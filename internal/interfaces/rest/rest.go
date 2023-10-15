package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bitcodr/re-test/internal/domain/model"
	shipmentservice "github.com/bitcodr/re-test/internal/domain/service/shipment"
	"github.com/bitcodr/re-test/internal/infrastructure/config"
	"github.com/bitcodr/re-test/internal/infrastructure/helper"
)

// Rest Service we can add all our services in here and pass it to our rest patterns
type Rest struct {
	ShipmentService shipmentservice.IShipment

	// register your services in the restful transport layer here
}

// InitTransport to initialise http apis
// it is possible in a project that use multiple rest like grpc, http, commandline, etc
func InitTransport(ctx context.Context, rest *Rest, config *config.Service) error {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Hello, world!\n")
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.HandleFunc("/orders/calculate", rest.Calculate)

	http.HandleFunc("/packets/update", rest.UpdatePacket)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.PORT
	}

	srv := http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: config.ReadTimeout,
	}

	log.Println("Server is running on port: ", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, config.IdleTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

type OrderRequestDTO struct {
	Items int `json:"items"`
}

// Calculate order items
//
//	@Summary		Order Items
//	@Description	Get correct order items
//	@Tags			order, items, calculate
//	@Accept			json
//	@Produce		json
//	@Param			items	body		int	false	"items to calculate"	minimum(1)
//	@Success		200		{object}	model.Order
//	@Failure		400		{string}	string	"error"
//	@Failure		404		{string}	string	"error"
//	@Failure		500		{string}	string	"error"
//	@Router			/orders/calculate [post]
func (r *Rest) Calculate(res http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		helper.ResponseError(res, "cannot read payload", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			helper.ResponseError(res, "cannot close payload", err)
			return
		}
	}(req.Body)

	var order OrderRequestDTO

	if err = json.Unmarshal(b, &order); err != nil {
		helper.ResponseError(res, "cannot decode payload", err)
		return
	}

	if order.Items <= 0 {
		helper.ResponseError(res, "items can't be empty or 0", nil)
		return
	}

	result, err := r.ShipmentService.Calculate(req.Context(), order.Items)
	if err != nil {
		helper.ResponseError(res, "error in calculate order items", err)
		return
	}

	helper.ResponseSuccess[*model.Order](res, result)
}

type UpdatePacketRequestDTO struct {
	Packets []int `json:"packets"`
}

// UpdatePacket update available packs
//
//	@Summary		update packs
//	@Description	update available packs for order
//	@Tags			order, items, update
//	@Accept			json
//	@Produce		json
//	@Param			items	body		array	false	"name search by q"	Format(email)
//	@Success		200		{array}		int
//	@Failure		400		{string}	string	"error"
//	@Failure		404		{string}	string	"error"
//	@Failure		500		{string}	string	"error"
//	@Router			/packets/update [post]
func (r *Rest) UpdatePacket(res http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		helper.ResponseError(res, "cannot read payload", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			helper.ResponseError(res, "cannot close payload", err)
			return
		}
	}(req.Body)

	var packets UpdatePacketRequestDTO

	if err = json.Unmarshal(b, &packets); err != nil {
		helper.ResponseError(res, "cannot decode payload", err)
		return
	}

	if packets.Packets == nil {
		helper.ResponseError(res, "packets cannot be empty", nil)
		return
	}

	items, err := r.ShipmentService.UpdatePacket(req.Context(), packets.Packets)
	if err != nil {
		helper.ResponseError(res, "cannot update packet", err)
		return
	}

	helper.ResponseSuccess[[]int](res, items)
}
