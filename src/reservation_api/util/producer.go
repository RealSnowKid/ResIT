package util

import (
	"github.com/FontysResIT/ResIT/config"
	"github.com/FontysResIT/ResIT/model"
	"github.com/FontysResIT/ResIT/util/kafka"
	log "github.com/sirupsen/logrus"
)

type IProducer interface {
	CreateReservation(model.Reservation)
	CancelReservation(model.Reservation)
}

var _producers []IProducer

func InitProducers() {
	config := config.GetConfig()
	if config.GetBool("kafka.enabled") {
		log.Info("Initializing with a kafka producer")
		_producers = append(_producers, kafka.NewProducer(config))
	}
}

func CreateReservation(reservation model.Reservation) {
	for _, producer := range _producers {
		go producer.CreateReservation(reservation)
	}
}

func CancelReservation(reservation model.Reservation) {
	for _, producer := range _producers {
		go producer.CancelReservation(reservation)
	}
}