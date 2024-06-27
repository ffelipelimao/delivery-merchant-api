package config

type Config struct {
	DBUrl    string
	QueueUrl string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	c.DBUrl = "postgres://adminfiap:adminfiap@rds-delivery.czdped39fibl.us-east-1.rds.amazonaws.com:5432/delivery?sslmode=disable"
	c.QueueUrl = "https://sqs.us-east-1.amazonaws.com/399351524471/cs_order_production_0.fifo"

	return nil
}
