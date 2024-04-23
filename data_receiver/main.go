package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tolling/ctypes"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

var kafkaTopic = "obudata"

type DataReceiver struct {
	msgch chan ctypes.OBUData
	conn  *websocket.Conn
	prod  *kafka.Producer
}

func main() {

	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}

	// Start another goroutine to check if we have delevered the data
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &DataReceiver{
		msgch: make(chan ctypes.OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data ctypes.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// topic := kafkaTopic
	err = dr.prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kafkaTopic,
			Partition: kafka.PartitionAny},
		Value: b,
	}, nil)
	return err
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU client connected")
	for {
		var data ctypes.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}
		fmt.Println("received msg:", data)
		if err := dr.produceData(data); err != nil {
			fmt.Println("kafka produce error:", err)
		}
	}
}
