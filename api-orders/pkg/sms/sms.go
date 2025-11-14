package sms

import (
	"api-orders/configs"
	"fmt"
)

type Sms struct {
	config *configs.SmsConfig
}

func New(config *configs.SmsConfig) *Sms {
	return &Sms{
		config: config,
	}
}

func (sms *Sms) Send(phone, message string) error {
	//TODO когда-нибудь здесь будет реальная отправка СМС, а пока что просто пишем в консоль
	fmt.Println("##### Заглушка отправки СМС #####")
	fmt.Println("Номер:", phone)
	fmt.Println("Сообщение:", message)
	fmt.Println()
	return nil
}