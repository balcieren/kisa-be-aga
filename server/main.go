package main

import (
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Url struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	ShortId   string    `json:"short_id" gorm:"not null"`
	Url       string    `json:"url" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

type UrlBody struct {
	Url string `json:"url" validate:"required,url"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func (ub *UrlBody) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(ub)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse

			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	db, err := gorm.Open(postgres.Open(os.Getenv("DB")))

	if err != nil {
		log.Fatal("Couldn't connect to database")
	}

	db.AutoMigrate(&Url{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Server has stared."})
	})

	v1 := app.Group("/api/v1/url")

	v1.Get("/:shortid", func(c *fiber.Ctx) error {
		u := Url{}
		result := db.Model(&Url{}).First(&u, "short_id = ?", c.Params("shortid"))

		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Url has not found"})
		}

		return c.Redirect(u.Url)
	})

	v1.Post("/", func(c *fiber.Ctx) error {
		ub := new(UrlBody)

		if err := c.BodyParser(ub); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
		}

		if err := ub.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
		}

		sid, _ := gonanoid.New(5)
		u := Url{ShortId: sid, Url: ub.Url}

		db.Model(&Url{}).Create(&u)

		return c.JSON(u)

	})

	app.Listen(":4000")

}
