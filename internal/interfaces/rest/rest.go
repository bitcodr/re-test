package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	shipmentservice "github.com/bitcodr/re-test/internal/domain/service/shipment"
	"github.com/bitcodr/re-test/internal/infrastructure/config"
)

// Rest Service we can add all our services in here and pass it to our rest patterns
type Rest struct {
	ShipmentService shipmentservice.IShipment

	// register your services in the restful transport layer here
}

// InitTransport InitCommand to initialise http apis
// it is possible in a project that use multiple rest like grpc, http, commandline, etc
func InitTransport(rest *Rest, config *config.Service) error {

	http.HandleFunc("orders/calculate", rest.Calculate)

	http.HandleFunc("packets/update", rest.UpdatePacket)

	err := http.ListenAndServe(config.PORT, nil)
	if err != nil {
		return err
	}

	return nil
}

type OrderRequestDTO struct {
	Items uint `json:"items"`
}

// Calculate Show @Summary get a packet
// @Description crawl
// @Tags Shipment
func (r *Rest) Calculate(res http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	var order OrderRequestDTO

	if err = json.Unmarshal(b, &order); err != nil {
		fmt.Println(err)
	}

	if err = req.Body.Close(); err != nil {
		fmt.Println(err)
	}

	result, err := r.ShipmentService.Calculate(req.Context(), order.Items)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Add("Content-Type", "application/json")
	_, err = res.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}

type UpdatePacketRequestDTO struct {
	Packets []uint
}

func (r *Rest) UpdatePacket(res http.ResponseWriter, req *http.Request) {

	b, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}

	var packets UpdatePacketRequestDTO

	if err = json.Unmarshal(b, &packets); err != nil {
		fmt.Println(err)
	}

	if err = req.Body.Close(); err != nil {
		fmt.Println(err)
	}

	result, err := r.ShipmentService.UpdatePacket(req.Context(), packets.Packets)
	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Add("Content-Type", "application/json")
	_, err = res.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
