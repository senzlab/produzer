package main

import (
    "fmt"
    //"time"
	"log"
    "os"
	"github.com/Shopify/sarama"
)

const (
    kafkaConn = "10.4.1.29:9092"
    topic = "test"
)

func main() {
    // create producer
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)
    producer, err := initProducer()
    if err != nil {
        fmt.Println("Error producer:", err.Error())
        os.Exit(1)
    }

    // publish data
    //producer.Input() <- &sarama.ProducerMessage{
    //    Topic: topic,
    //    Partition: 0,
    //    Value: sarama.StringEncoder("helllll"),
    //}
    msg := &sarama.ProducerMessage {
        Topic: "test",
        Partition: 0,
        Value: sarama.StringEncoder("era"),
    }
    p, o, err := producer.SendMessage(msg)
    if err != nil {
        fmt.Println("Error producer:", err.Error())
    }

    fmt.Println("partition :", p)
    fmt.Println("offset :", o)
}

func initProducer()(sarama.SyncProducer, error) {
	config := sarama.NewConfig()
    //config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
    //prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)
    prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)
    if err != nil {
        fmt.Println("Error producer:", err.Error())
        return nil, err
    }

    return prd, nil
}
