package handlers

import (
	//"html/template"

	"bytes"

	//"fmt"
	"strconv"

	"encoding/json"

	//	"log"

	"time"

	model "github.com/t67y110v/software-engineering/internal/app/model/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handlres) Register() fiber.Handler {

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

func (h *Handlres) Login() fiber.Handler {

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

func (h *Handlres) CheckJWT() fiber.Handler {

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

func (h *Handlres) Logout() fiber.Handler {

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
