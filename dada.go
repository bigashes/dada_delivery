package dada_delivery

type Config struct {
	AppKey    string
	AppSecret string
	SourceId  string
}

type Client struct {
	config Config
	Debug  bool
}

func NewClient(config Config) *Client {
	return &Client{
		config: config,
	}
}
