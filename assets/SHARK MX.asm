; checksum calculation subroutine
ROM0:0EE1 11 DB C3         ld   de,C3DB
ROM0:0EE4 AF               xor  a
ROM0:0EE5 0E 0B            ld   c,0B
ROM0:0EE7 2A               ldi  a,(hl)
ROM0:0EE8 D6 30            sub  a,30
ROM0:0EEA 87               add  a
ROM0:0EEB FE 0A            cp   a,0A
ROM0:0EED 38 02            jr   c,0EF1
ROM0:0EEF D6 09            sub  a,09
ROM0:0EF1 12               ld   (de),a
ROM0:0EF2 13               inc  de
ROM0:0EF3 2A               ldi  a,(hl)
ROM0:0EF4 D6 30            sub  a,30
ROM0:0EF6 12               ld   (de),a
ROM0:0EF7 13               inc  de
ROM0:0EF8 0D               dec  c
ROM0:0EF9 20 EC            jr   nz,0EE7
ROM0:0EFB 21 DB C3         ld   hl,C3DB
ROM0:0EFE AF               xor  a
ROM0:0EFF 0E 16            ld   c,16
ROM0:0F01 86               add  (hl)
ROM0:0F02 23               inc  hl
ROM0:0F03 0D               dec  c
ROM0:0F04 20 FB            jr   nz,0F01
ROM0:0F06 47               ld   b,a
ROM0:0F07 A7               and  a
ROM0:0F08 28 0B            jr   z,0F15
ROM0:0F0A 0E 00            ld   c,00
ROM0:0F0C 0C               inc  c
ROM0:0F0D D6 0A            sub  a,0A
ROM0:0F0F 28 04            jr   z,0F15
ROM0:0F11 30 F9            jr   nc,0F0C
ROM0:0F13 2F               cpl  
ROM0:0F14 3C               inc  a
ROM0:0F15 21 DB C3         ld   hl,C3DB
ROM0:0F18 C6 30            add  a,30
ROM0:0F1A 22               ldi  (hl),a
ROM0:0F1B 79               ld   a,c
ROM0:0F1C D6 0A            sub  a,0A
ROM0:0F1E 30 FC            jr   nc,0F1C
ROM0:0F20 C6 3A            add  a,3A
ROM0:0F22 22               ldi  (hl),a
ROM0:0F23 AF               xor  a
ROM0:0F24 22               ldi  (hl),a
ROM0:0F25 C9               ret  
