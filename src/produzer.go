package main

import (
    "fmt"
	"log"
    "os"
    "bufio"
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
        fmt.Println("Error producer: ", err.Error())
        os.Exit(1)
    }

    // read command line input
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter msg: ")
        msg, _ := reader.ReadString('\n')

        // publish
        go publish(msg, producer)
    }
}

func initProducer()(sarama.SyncProducer, error) {
	config := sarama.NewConfig()
    config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

    // async producer
    //prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

    // sync producer
    prd, err := sarama.NewSyncProducer([]string{kafkaConn}, config)

    return prd, err
}

func publish(message string, producer sarama.SyncProducer) {
    // publish sync
    msg := &sarama.ProducerMessage {
        Topic: topic,
        Value: sarama.StringEncoder(message),
    }
    p, o, err := producer.SendMessage(msg)
    if err != nil {
        fmt.Println("Error publish: ", err.Error())
    }

    // publish async
    //producer.Input() <- &sarama.ProducerMessage{

    fmt.Println("Partition: ", p)
    fmt.Println("Offset: ", o)
}
