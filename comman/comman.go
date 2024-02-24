package comman

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Getenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file" + err.Error())
		return err
	}
	return nil
}

func Searlize(data string) string {
	return fmt.Sprintf("%x", data)
}

func DeSerilizeData(s string) string {
	var response string
	_, err := fmt.Sscanf(s, "%x", &response)
	if err != nil {
		fmt.Println("Error reverting:", err)
	}
	return response
}
