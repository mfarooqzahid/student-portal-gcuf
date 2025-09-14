package main

import (
	"fmt"
	"log"
	"os"
	"student-portal-gcuf/internal/ai"
	"student-portal-gcuf/internal/config"
	"student-portal-gcuf/internal/portal"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found, using system env instead")
	}

	cfg := config.LoadConfig("config/config.yaml")
	genai := ai.NewGenAiClient()
	client := portal.NewPortalClient(&cfg, &genai)

	// TODO: add username and password before running
	// LOGIN TO PORTAL
	username := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	ok, err := client.Login(username, password)
	if err != nil || !ok {
		log.Fatalf("login failed:%v", err)
	}
	log.Println("login successful")

	profile, err := client.GetProfile()
	if err != nil {
		log.Fatalf("failed to get profile: %v", err)
	}

	fmt.Printf("profile: %v\n", profile)

	client.GetAcademics()
}
