package repositories

type RabbitMQRepository interface {
	CloudApprove(url []string, _type string) error
}
