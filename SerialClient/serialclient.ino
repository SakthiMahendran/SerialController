class SerialClient {
    public:
        void begin() {
            Serial.begin(115200);
            this->initHandShake();
        }

        void mouseLeftClick() {
            Serial.write(SerialClient::LeftClick);
        }

        void mouseRightClick() {
            Serial.write(SerialClient::RightClick);
        }

        void mouseScroll(int x, int y) {
            Serial.write(SerialClient::Scroll);

            Serial.println(x);
            Serial.println(y);
        }

        void mouseScrollRelative(int x, int y) {
            Serial.write(SerialClient::Scroll);

            Serial.println(x);
            Serial.println(y);
        }

        void mouseMove(int x, int y) {
            Serial.write(SerialClient::Move);

            Serial.println(x);
            Serial.println(y);
        }

        void mouseMoveRelative(int x, int y) {
            Serial.write(SerialClient::MoveRelative);

            Serial.println(x);
            Serial.println(y);
        }

    private:
        const byte LeftClick      = 1;
	    const byte RightClick     = 2;
	    const byte Scroll         = 3;
	    const byte ScrollRelative = 4;
	    const byte Move           = 5;
	    const byte MoveRelative   = 6;

        void initHandShake() {
            Serial.print("\n");
            Serial.print("Hello, I am Arduino\n");
        }
};

SerialClient sc = SerialClient();
void setup() {
    sc.begin();
}

void loop() {
    // delay(2000);
    // sc.mouseLeftClick();
    delay(2000);
    sc.mouseMove(0, 0);
    delay(2000);
    sc.mouseMoveRelative(5, 5);
    delay(2000);
    sc.mouseRightClick();
    delay(2000);
    sc.mouseScroll(9, 10);
    delay(2000);
    sc.mouseScrollRelative(10, 10);
}
