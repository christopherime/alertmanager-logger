package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	// check if the log directory exists
	// if not, create it
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", 0755)
	}

	// check if the log file exists
	// if not, create it
	if _, err := os.Stat("log/alerts.log"); os.IsNotExist(err) {
		os.Create("log/alerts.log")
	}

}

func writeAlertLog(alertManagerMessage AlertManagerNotificationObject) {
	log.Println("Alert received")
	for _, alert := range alertManagerMessage.Alerts {
		log.Println(alert.Labels.Alertname, alert.Labels.Instance, alert.Status)
	}
}

func main() {

	// defined default log file
	f, err := os.OpenFile("log/alerts.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("Starting Alert Manager Logger on port 9095")

	app := fiber.New()

	app.Post("/logger", func(c *fiber.Ctx) error {

		var alertManagerMessage AlertManagerNotificationObject

		c.BodyParser(&alertManagerMessage)

		writeAlertLog(alertManagerMessage)

		return c.SendString("Notification received")
	})

	app.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusForbidden).SendString("Forbidden")
	})

	err = app.Listen(":9095")
	if err != nil {
		panic(err)
	}
}
