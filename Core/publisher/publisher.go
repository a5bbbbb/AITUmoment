package publisher

import (
	"aitu-moment/logger"
	"aitu-moment/utils"
	"encoding/json"
	"sync"

	"github.com/wagslane/go-rabbitmq"
)

var (
	ampqConn *rabbitmq.Conn
	once     sync.Once
)

type EmailPublisher struct {
	publisher *rabbitmq.Publisher
}

func init() {
	logger.GetLogger().Info("Core: setting up ampq conn....")
	once.Do(establishAmpqConn)
	logger.GetLogger().Info("Core: Established ampq conn")

}

func establishAmpqConn() {
	conn, err := rabbitmq.NewConn(
		utils.GetFromEnv("RABBITMQ_URL", "amqp://guest:guest@localhost"),
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		logger.GetLogger().Errorf("Core, could not establish ampq conn: %v", err.Error())
		panic("Core,could not establish ampq conn.")
	}

	ampqConn = conn

}

func NewEmailPublisher() (*EmailPublisher, error) {

	publisher, err := rabbitmq.NewPublisher(
		ampqConn,
		rabbitmq.WithPublisherOptionsExchangeName(utils.GetFromEnv("EmailExchange", "email_exchange")),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsLogging,
	)

	if err != nil {
		logger.GetLogger().Errorf("Core, could not declar the email publisher: %v", err.Error())
		return nil, err
	}

	return &EmailPublisher{publisher: publisher}, nil

}

func (c *EmailPublisher) ClosePublisher() {
	c.publisher.Close()
}

func (c *EmailPublisher) PublishEmailSendMessage(email, subject, body string) error {

	msg := map[string]string{"email": email, "subject": subject, "body": body}
	payload, _ := json.Marshal(msg)

	err := c.publisher.Publish(
		payload,
		[]string{utils.GetFromEnv("SendEmailQueueRoutingKey", "send_email")},
		rabbitmq.WithPublishOptionsContentType("application/json"),
	)

	if err != nil {
		logger.GetLogger().Errorf("Core, error during publishing email send message: %v", err.Error())
		return err
	}

	return nil
}

func CloseConn() {
	logger.GetLogger().Info("Core: Closing ampq conn...")
	err := ampqConn.Close()

	if err != nil {
		logger.GetLogger().Errorf("Core: error during ampq conn close: %v", err.Error())
	}
}
