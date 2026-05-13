package port

import (
	"OrderApp/service/notification/gmail/model"
)

type MailServicePort interface {
	SendSuccessfullyOrderPlayed(to string, data model.SuccessfullyOrderPlayedData)
}
