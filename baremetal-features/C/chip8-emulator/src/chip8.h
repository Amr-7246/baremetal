#include <stdint.h>

typedef struct {
    uint8_t memory[4096]; // 4KB of RAM 
    uint8_t v[16]; // Registers V0 through VF
    uint16_t i;    // Index register which holds memory addresses
    uint16_t pc;   // The Program Counter which tracks the current instruction

    uint16_t stack; // The Stack (stores return addresses)
    uint8_t sp;     // Stack Pointer (points to the top of the stack)

    uint8_t  delay_timer;  // Decrements at 60Hz
    uint8_t  sound_timer;  // Decrements at 60Hz (beeps when > 0)

    uint8_t display[64 * 32]; // Screen pixels 
    uint8_t keypad[16];       // Current state of the 16 hex keys
} Chip8;