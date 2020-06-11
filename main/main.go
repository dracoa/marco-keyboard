package main

import (
	"log"
	"marco-keyboard/keyhook"
	"marco-keyboard/microrobot"
	"strings"
)

const Version = "0.0.0"

func init() {}

func main() {
	log.Printf("Version: %s", Version)
	keyhook.Start()

	select {}

	robot := robot()
	if robot == nil {
		log.Fatal("No micro keyboard found.")
	}
	defer robot.Disconnect()
}

func robot() *microrobot.MicroRobot {
	var robot *microrobot.MicroRobot
	for com, name := range ListSerial() {
		if strings.Contains(name, "USB") {
			log.Printf("Testing connect to %s: %s", com, name)
			robot = microrobot.Connect(com)
		}
		if robot != nil {
			return robot
		}
	}
	return nil
}
