//Test program for testing the SerialClient
#include "SerialClient.h"

SerialClient sc = SerialClient();
void setup() {
    sc.begin();
}

void loop() {
    delay(2000);
    sc.move(100, 100);
}
