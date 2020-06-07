#include <ArduinoJson.h>
#include "Keyboard.h"
#include "Mouse.h"

void setup() {
  // initialize serial:
  Serial.begin(9600);
  Mouse.begin();
  Keyboard.begin();
}

void loop() {
  while (Serial.available()) {
    String json = Serial.readStringUntil('\n');
    Serial.println(json);
    StaticJsonDocument<200> doc;
    DeserializationError error = deserializeJson(doc, json);
    // Test if parsing succeeds.
    if (error) {
      Serial.print(F("deserializeJson() failed: "));
      Serial.println(error.c_str());
      return;
    } 
    const String cmd = doc["cmd"].as<String>();
    const uint8_t c = doc["char"];
    const uint8_t cmd_delay = doc["delay"];
    delay(cmd_delay);
    if (cmd == "key_write"){
      Keyboard.write(c);
    }
    if (cmd == "key_press"){
      Keyboard.press(c);
    }
    if (cmd == "key_release"){
      Keyboard.release(c);
    }
    if (cmd == "release_all"){
      Keyboard.releaseAll();
    }
    if (cmd == "left_click"){
      Mouse.click(MOUSE_LEFT);
    }
    if (cmd == "right_click"){
      Mouse.click(MOUSE_RIGHT);
    }
    if (cmd == "mouse_press"){
      Mouse.press(c);
    }
    if (cmd == "mouse_release"){
      if (Mouse.isPressed(c)){
        Mouse.release(c);
      }
    }
    if (cmd == "mouse_move"){
      const char x = doc["x"];
      const char y = doc["y"];      
      Mouse.move(x, y);
    }    
  }
}
