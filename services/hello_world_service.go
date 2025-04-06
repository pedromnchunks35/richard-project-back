package services

type HelloWorldServiceImpl struct{}

func (h HelloWorldServiceImpl) HelloWorld(msg string) string {
	return "Hello, the following message is about to be presented: " + msg
}

func NewHelloWorldImpl() *HelloWorldServiceImpl {
	return &HelloWorldServiceImpl{}
}
