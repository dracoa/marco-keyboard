#include "Keyboard.h"
#include "Mouse.h"

typedef void (*FuncPtr)(uint8_t p1, uint8_t p2);
FuncPtr controls[]={&echo, &key_write, &key_press, &key_release, &key_releaseAll, &mouse_click, &mouse_press, &mouse_release, &mouse_move};

void setup() {
  Serial.begin(9600);
  Mouse.begin();
  Keyboard.begin();
}

void echo(uint8_t p1, uint8_t p2){
  Serial.print(p1);
}

void key_write(uint8_t p1, uint8_t p2) {
  Keyboard.write(p1);
}

void key_press(uint8_t p1, uint8_t p2) {
  Keyboard.press(p1);
}

void key_release(uint8_t p1, uint8_t p2) {
  Keyboard.release(p1);
}

void key_releaseAll(uint8_t p1, uint8_t p2) {
  Keyboard.releaseAll();
}

void mouse_click(uint8_t p1, uint8_t p2) {
   Mouse.click(p1);
}

void mouse_press(uint8_t p1, uint8_t p2) {
   Mouse.press(p1);
}

void mouse_release(uint8_t p1, uint8_t p2) {
  if (Mouse.isPressed(p1)){
    Mouse.release(p1);
  }
}

void mouse_move(uint8_t p1, uint8_t p2) {
   Mouse.move(p1, p2);
}

void loop() {
  while (Serial.available()) {
    String json = Serial.readStringUntil('\n');
    char cmd = json.charAt(0);
    char p1 = json.charAt(1);
    char p2 = json.charAt(2);
    controls[cmd](p1, p2);
  }
}
