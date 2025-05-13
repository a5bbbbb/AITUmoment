package consumer

import (
	config "email_service/configs"
	"email_service/internal/logger"
	"email_service/internal/models"
	"email_service/internal/services"
	"encoding/json"
	"sync"

	"github.com/wagslane/go-rabbitmq"
)

var (
	ampqConn *rabbitmq.Conn
	once     sync.Once
)

type EmailConsumer struct {
	consumer *rabbitmq.Consumer
}

func init() {
	logger.GetLogger().Info("Email service: setting up ampq conn....")
	once.Do(establishAmpqConn)
	logger.GetLogger().Info("Email service: Established ampq conn")

}

func establishAmpqConn() {
	conn, err := rabbitmq.NewConn(
		config.RabbitMQURL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		logger.GetLogger().Errorf("Email service, could not establish ampq conn: %v", err.Error())
		panic("Email service,could not establish ampq conn.")
	}

	ampqConn = conn

}

func NewEmailConsumer() (*EmailConsumer, error) {

	logger.GetLogger().Info("Email service: creating new consumer")

	consumer, err := rabbitmq.NewConsumer(
		ampqConn,
		config.RoutingKey,
		rabbitmq.WithConsumerOptionsExchangeName(config.EmailExchange),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsRoutingKey(config.RoutingKey),
		rabbitmq.WithConsumerOptionsLogging,
	)

	if err != nil {
		logger.GetLogger().Errorf("Email service, could not declar the email consumer: %v", err.Error())
		return nil, err
	}

	logger.GetLogger().Info("Email service: created new consumer")

	return &EmailConsumer{consumer: consumer}, nil

}

func (c *EmailConsumer) CloseConsumer() {
	c.consumer.Close()
}

func (c *EmailConsumer) RunConsumer() error {
	logger.GetLogger().Info("Email service: running the consumer....")

	err := c.consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		logger.GetLogger().Infof("Email service, consumed: %v", string(d.Body))

		var msg models.EmailMessage
		if err := json.Unmarshal(d.Body, &msg); err != nil {
			logger.GetLogger().Errorf("Failed to parse message: %v", err)
			return rabbitmq.Ack
		}

		services.SendMail(msg.Email, msg.Subject, msg.Body)

		return rabbitmq.Ack
	})

	if err != nil {
		logger.GetLogger().Errorf("Email service, error during consumer run: %v", err.Error())
		return err
	}

	return nil
}

func CloseConn() {
	logger.GetLogger().Info("Email service: Closing ampq conn...")
	err := ampqConn.Close()

	if err != nil {
		logger.GetLogger().Errorf("Email service: error during ampq conn close: %v", err.Error())
	}
}
