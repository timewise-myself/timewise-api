package transport

import "github.com/gofiber/fiber/v2"

type AccountHandlerRegister struct {
	Router  fiber.Router
	Handler *AccountHandler
}

func RegisterAccountHandler(router fiber.Router) {
	handler := NewAccountHandler()
	accountHandler := &AccountHandlerRegister{
		Handler: handler,
	}

	// Register all endpoints here
	router.Get("/user", accountHandler.Handler.getUserInfo)
	router.Get("/user/emails", accountHandler.Handler.getLinkedUserEmails)
	router.Patch("/user", accountHandler.Handler.updateUserInfo)
	router.Post("/user/emails/send", accountHandler.Handler.sendLinkEmailRequest)
	router.Get("/user/emails/link/:token", accountHandler.Handler.actionEmailLinkRequest)
	router.Post("/user/emails/unlink", accountHandler.Handler.unlinkAnEmail)
	router.Post("/user/deactivate", accountHandler.Handler.deactivateAccount)
	router.Get("/user/emails/parent", accountHandler.Handler.getParentLinkedEmails)
	router.Get("/user/emails/clear-rejected", accountHandler.Handler.clearStatusRejectedEmail)
}
