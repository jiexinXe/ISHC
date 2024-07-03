package config

import (
	"database/sql"
	"fmt"
	"github.com/IBM/sarama"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

var (
	KafkaProducer sarama.SyncProducer
	KafkaConsumer sarama.ConsumerGroup
)

func InitDB() {
	var err error
	dsn := "old_care:Tz5ckZHNEDmtJEZr@tcp(47.93.76.253:3306)/old_care"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully")
}

func InitKafka() {
	kafkaBrokers := []string{"47.102.213.168:9092"}
	groupId := "test-consumer-group"

	// 生产者配置
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(kafkaBrokers, producerConfig)
	if err != nil {
		log.Fatalf("无法启动 Sarama 生产者: %v", err)
	}
	KafkaProducer = producer

	// 消费者配置
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Return.Errors = true
	consumerConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumer, err := sarama.NewConsumerGroup(kafkaBrokers, groupId, consumerConfig)
	if err != nil {
		log.Fatalf("无法启动 Sarama 消费者: %v", err)
	}
	KafkaConsumer = consumer

	log.Println("Kafka 生产者和消费者已初始化")
}
