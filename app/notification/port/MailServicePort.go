package port

import "OrderApp/notification/gmail/model"

type MailServicePort interface {
	SendSuccessfullyOrderPlayed(to string, data model.SuccessfullyOrderPlayedData)
}
