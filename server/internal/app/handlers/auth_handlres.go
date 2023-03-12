package handlers

import (
	//"html/template"

	"bytes"
	"log"

	//"fmt"
	"strconv"

	"encoding/json"

	//	"log"

	"time"

	model "github.com/t67y110v/software-engineering/internal/app/model/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// @Summary User Registration
// @Description registration of user
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Registration  true "registration"
// @Success 200 {object} responses.Registration
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/register [post]
func (h *Handlers) Register() fiber.Handler {

	return func(c *fiber.Ctx) error {

		type request struct {
			Email       string `json:"email"`
			Password    string `json:"password"`
			Name        string `json:"name"`
			SeccondName string `json:"seccond_name"`
		}

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle register, status :%d, error :%e", fiber.StatusBadRequest, err)
		}
		log.Println(req)
		u := &model.User{
			Email:       req.Email,
			Password:    req.Password,
			Name:        req.Name,
			SeccondName: req.SeccondName,
		}

		if err := h.pgStore.UserRepository().Create(u); err != nil {
			return err
		}

		u.Sanitize()

		return c.JSON(u)

	}

}

// @Summary User Login
// @Description authentification user in the system
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Login true  "login"
// @Success 200 {object} responses.Login
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/login [post]
func (h *Handlers) Login() fiber.Handler {

	return func(c *fiber.Ctx) error {

		type request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle login, status :%d, error :%e", fiber.StatusBadRequest, err)
		}

		u, err := h.pgStore.UserRepository().FindByEmail(req.Email)
		if err != nil {
			return err
		}

		if u.ID == 0 {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "user not found",
			})
		}

		secret := "11we$*9sd*(@!)"

		minutesCount, _ := strconv.Atoi("15")

		claims := jwt.MapClaims{}

		claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

		claims["id"] = u.ID

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(secret))
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"token": t,
			"name":  u.Name,
			"email": u.Email,
		})
	}

}

// @Summary Check session
// @Description Validation user token
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.CheckToken true  "Check token"
// @Success 200 {object} responses.Login
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/check [post]
func (h *Handlers) CheckJWT() fiber.Handler {

	return func(c *fiber.Ctx) error {

		type request struct {
			Cookie string `json:"token"`
		}

		req := &request{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
		}

		cookie := req.Cookie

		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.JSON(fiber.Map{
				"message": "token id is nil",
			})
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		return c.JSON(u)

	}

}

func (h *Handlers) Logout() fiber.Handler {

	return func(c *fiber.Ctx) error {

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}

}
