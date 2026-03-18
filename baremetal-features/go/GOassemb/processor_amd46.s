// High-Performance Cryptographic Hashing Tool
// Go has a special feature called Go Assembly (Plan 9). It allows us to write functions in .s files and call them directly from .go files.
// This is how the Go standard library makes math and crypto so fast . . . SO this is very simple project (as an entry point) to implement it


// textflag.h contains constants for the linker
#include "textflag.h"

// func ProcessBlock(data []byte)
// A slice in Go is 24 bytes: [Pointer(8), Len(8), Cap(8)] //? what does that mean

TEXT ·ProcessBlock(SB), NOSPLIT, $0-24 //? what does (TEXT, loop, done) mean
    MOVQ data_base+0(FP), SI  //! Load the Pointer to the array into SI
    MOVQ data_len+8(FP), CX //! Load the Length into CX (the loop counter)

loop :
    CMPQ CX, $0 //! check if length is Zero
    JE done

    MOVB (SI), AX  //! Load byte at SI into AX register //? what does the () represent, why not just MOVB SI, AX
    ADDB $1, AX //! increment the byte (here is the logic)
    MOVB AX, (SI) //Store it back

    ADDQ $1, SI //to move it to the next memory byte
    SUBQ $1, CX //Decrement the counter
    JMP loop

done :
    RET
