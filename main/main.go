package main

import (
	"fmt"
	"log"
	"marco-keyboard/keyhook"
	"marco-keyboard/microrobot"
	"strings"
)

const Version = "0.0.0"

func init() {}

func main() {
	log.Printf("Version: %s", Version)
	robot := robot()
	if robot == nil {
		log.Fatal("No micro keyboard found.")
	}
	defer robot.Disconnect()

	keyhook.KeyboardHook().ForEach(func(v interface{}) {
		fmt.Printf("%v", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})

	select {}
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
