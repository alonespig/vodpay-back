package mq

import (
	"context"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Producer rocketmq.Producer

const (
	OrderConsumerGroup = "vodpay-order_group"
	OrderTopic         = "vodpay-order_topic" // 订单主题
)

func InitProducer() error {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	if err != nil {
		return err
	}

	if err := p.Start(); err != nil {
		return err
	}

	Producer = p
	return nil
}

func SendOrderCreated(orderID int64) error {
	msg := &primitive.Message{
		Topic: OrderTopic,
		Body:  []byte(strconv.FormatInt(orderID, 10)),
	}

	msg.WithTag("order_created") // 订单创建标签

	_, err := Producer.SendSync(context.Background(), msg)
	return err
}

func SendOrderQuery(orderID int64) error {
	msg := &primitive.Message{
		Topic: OrderTopic,
		Body:  []byte(strconv.FormatInt(orderID, 10)),
	}

	msg.WithTag("order_query") // 订单查询标签
	msg.WithDelayTimeLevel(3)

	_, err := Producer.SendSync(context.Background(), msg)
	return err
}
