#include "chip8.h"
#include <stdio.h> //? Tell me more about that header usage and contants and from where sould O learn that
#include <stdlib.h> //? should I need that header + Tell me more about that header usage and contants and from where sould O learn that

int load_rom(Chip8 *chip8, const char *filename){
    //& open the file in a binary mode, determine the size of that file, insure that there is a sufficient ram space
        FILE *file = fopen(filename, 'rb'); //? From where should I know the C functions segnature and docs + in the computer operations context, what does opening/closing the file mean
        if(file == NULL){
            fprintf(stderr, "Error: we caould not open the file (%s) in the binary mode amr \n sorry for that ...", filename); //? tell me more about the stderr and the fprintf too
            return -1;
        }
        fseek(file, 0, SEEK_END); //? Explain that function segnature :int fseek(FILE *_File, long _Offset, int _Origin) + what does SEEK_END mean
        long rom_size = ftell(file);
        rewind(file); //? why that step 
        long full_size = sizeof(chip8->memory) - 0x200; //? why we ignore 0X200 from the full memory space, tell me more about that "0x200"
        if (full_size < rom_size)
        {
            fprintf(stderr, "Error: ROM is too big amr (%ld bytes). Max is %ld bytes. \n sorry, try to buy some RAM slots ...", rom_size, full_size); //? tell me more about the stderr and the fprintf too
            return -1;
        }
    //& Allocate the ROM at the memory, clean up and give the contole to the CPU PC 
        size_t rom_bytes = fread(&chip8->memory[0x200], 1, rom_size, file); //? What does size_t mean + explain that seg: size_t fread(void *__restrict__ _DstBuf, size_t _ElementSize, size_t _Count, FILE *__restrict__ _File)
        if (rom_bytes != (size_t)rom_size) {
                fprintf(stderr, "Error: Failed to read the full ROM file.\n");
                fclose(file);
                return -1;
            }
        fclose(file); //? why should I close the file
        printf("Successfully loaded %ld bytes into memory.\n", rom_size);

    return 0;
}