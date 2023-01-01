#define MOUSE_LEFT 1
#define MOUSE_RIGHT 2
#define MOUSE_MIDDLE 3

class SerialClient {
private:
    const uint8_t Click   = 1;
    const uint8_t Move    = 2;
    const uint8_t Press   = 3;
    const uint8_t Release = 4;
    const uint8_t End     = 5;
    const uint8_t Scroll  = 6;

public:
  SerialClient(void);
  void begin(void);
  void end(void);
  void click(uint8_t b = MOUSE_LEFT);
  void move(int x, int y, signed char wheel = 0);
  void press(uint8_t b = MOUSE_LEFT);   // press LEFT by default
  void release(uint8_t b = MOUSE_LEFT); // release LEFT by default
//   bool isPressed(uint8_t b = MOUSE_LEFT); Will be implemented soon.
};
