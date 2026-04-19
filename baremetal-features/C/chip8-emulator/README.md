## Description ....
A CHIP-8 emulator which is the "Hello World" of systems programming. 
The main goal is bridging the gap between high-level logic and how hardware actually handles data.

## Task Flow ....

    1. Set the State Structure: Define a struct that holds the registers, 4KB of memory, the Program Counter (PC), the stack and so on.
    2. Building The File Loader (consider building a Relocating Loader to support ASLR and multitasking): Write a function to read a CHIP-8 ROM file into memory starting at address 0x200.
    3. The Fetch-Decode Loop: Create a while loop that reads two bytes (one 16-bit opcode), shifts them into a single variable, and uses a switch statement to identify the instruction.
    4. Core Instructions: Implement the "math" and "jump" opcodes first (e.g., 7XNN for addition, 1NNN for jumps).
    5. Graphics (SDL2): Integrate a library like SDL2 to create a window. Map your internal display[64 * 32] array to the screen.
    6. The Timers: Implement the Delay and Sound timers that decrement at 60Hz.
    7. Input: Map your keyboard keys to the 16-key hex keypad (0-F).

## Project Structure ....

``
chip8-emulator/
├── src/
│   ├── main.c          # Entry point, SDL2 initialization, and main loop
│   ├── chip8.c         # Logic for opcodes, timers, and memory
│   └── chip8.h         # The State struct and function prototypes
├── roms/               # Folder to store your .ch8 files
├── Makefile            # To compile with: gcc src/*.c -lSDL2 -o chip8
└── README.md
``