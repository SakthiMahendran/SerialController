#include "SerialClient.h"

void SerialClient::begin(void) {
    Serial.begin(115200);

    // Initial Handshake
    Serial.print("\n");
    Serial.print("Hello, I am Arduino\n");
}

void SerialClient::end(void) {
    Serial.write(SerialClient::End);
}

void SerialClient::click(uint8_t btn = MOUSE_LEFT) {
    Serial.write(SerialClient::Click);
    Serial.write(btn);
}

void SerialClient::move(int x, int y, signed char wheel = 0) {
    Serial.write(SerialClient::Move);

    Serial.println(x);
    Serial.println(y);

    if (wheel != 0) {
        Serial.write(SerialClient::Scroll);
        Serial.println(wheel);
    }    
}

void SerialClient::press(uint8_t btn = MOUSE_LEFT) {
    Serial.write(SerialClient::Press);
    Serial.write(btn);
}

void SerialClient::release(uint8_t btn = MOUSE_LEFT) {
    Serial.write(SerialClient::Release);
}