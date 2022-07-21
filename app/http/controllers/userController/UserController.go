package userController

import (
	userInterface2 "authMicroservice/app/http/interfaces/userInterface"
	"authMicroservice/app/http/services"
	"authMicroservice/app/model/UserModel"
	"authMicroservice/app/utils/handlers"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	Time "time"
)

func GetUser(c *fiber.Ctx) error {
	userId := c.Locals("userInfo")
	id, ok := userId.(string)
	if ok != true {
		return handlers.ErrorJsonOutput(c, nil, 500, "Internal server error")
	}
	strToInt, _ := strconv.Atoi(id)
	res, err := UserModel.FindById(strToInt)
	res.Gender = UserModel.GetGender(res.Gender)
	return handlers.JsonOutputOrError(c, res, err, 200, 500)
}
func UserDet(c *fiber.Ctx) error {
	res := c.Locals("userInfo")
	return c.Status(200).JSON(fiber.Map{
		"ok":   true,
		"data": res,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var signUp userInterface2.SignUpBody
	if err := c.BodyParser(&signUp); err != nil {
		fmt.Println("Error", err.Error())
		return handlers.ErrorJsonOutput(c, err, 400)
	}
	if err := signUp.Validate(); err != nil {
		return handlers.ErrorJsonOutput(c, err, 400)

	}
	_, err := UserModel.SignUpByLocal(signUp)
	return handlers.JsonOutputOrError(c, "Registration Successful", err, 201, 500)

}

func Login(c *fiber.Ctx) error {
	userDet := new(userInterface2.LoginBody)
	if err := c.BodyParser(&userDet); err != nil {
		fmt.Println("Error", err.Error())
		return handlers.ErrorJsonOutput(c, errors.New("input error"), 400, "only accepts json format")
	}
	if err := userDet.Validate(); err != nil {
		fmt.Println("Error", err.Error())
		return handlers.ErrorJsonOutput(c, errors.New("invalid Credentials"), 401)
	}
	//TODO check login
	result, err := UserModel.LoginByEmail(userDet.Email, userDet.Pass, c.IP())
	if err != nil {
		return handlers.ErrorJsonOutput(c, err, 401)
	}
	//if all Okay generate token
	token, _ := services.GenerateJWTToken(result.ID)

	//set cookie
	time := Time.Minute * 60
	cookie := handlers.SetCookie("Authorization", token, time, false, false)
	data := fiber.Map{
		"token":   token,
		"message": "login Successful",
		"expire":  time.Seconds(),
	}
	c.Cookie(cookie)
	return handlers.JsonOutputOrError(c, data, nil, 200, 500)
}

func Logout(c *fiber.Ctx) error {
	//userId := c.Locals("userInfo")
	time := Time.Minute * -1
	cookie := handlers.SetCookie("Authorization", "", time, false, false)
	c.Cookie(cookie)
	return handlers.JsonOutputOrError(c, "logout success", nil, 200, 500)
}
