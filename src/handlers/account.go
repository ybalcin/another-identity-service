package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	. "github.com/ybalcin/another-identity-service/account"
	"log"
)

//var accountService IAccountService

//func init() {
//	accountRepository := NewUserRepository()
//	accountService = NewAccountService(accountRepository)
//}

type accountHandler struct {
	GetAllHandler func(c *fiber.Ctx) error
}

var getAll = func(c *fiber.Ctx) error {
	accountRepository := NewUserRepository()
	accountService := NewAccountService(accountRepository)

	users, err := accountService.GetUserList()
	if err != nil {
		log.Fatal(err)
		// return err
	}
	byteSlice, e := json.Marshal(users)
	if e != nil {
		return e
	}
	c.SendString(string(byteSlice))
	return nil
}

func NewAccountHandler() *accountHandler {
	return &accountHandler{
		GetAllHandler: getAll,
	}
}
