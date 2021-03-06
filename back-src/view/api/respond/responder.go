package respond

import (
	"back-src/view/notifications"
)

func Respond(notification notifications.Notification) {
	ctx := notification.Context
	ctx.JSON(notification.StatusCode, notifications.Response{Message: notification.Message, Data: notification.Data})
}
