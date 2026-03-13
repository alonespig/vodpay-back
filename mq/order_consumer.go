package mq

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func StartOrderConsumer() error {

	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName(OrderConsumerGroup),
		consumer.WithConsumeMessageBatchMaxSize(1),
	)

	if err != nil {
		return err
	}

	err = c.Subscribe(OrderTopic, consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

			for _, msg := range msgs {

				orderID := string(msg.Body)

				fmt.Println("收到订单消息:", orderID)
			}

			return consumer.ConsumeSuccess, nil
		})

	if err != nil {
		return err
	}

	err = c.Start()
	if err != nil {
		return err
	}

	fmt.Println("Order consumer started")

	return nil
}
