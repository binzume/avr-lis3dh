CC=avr-gcc
CFLAGS=-std=c99 -O2 -Wall
TARGET=atmega328p
F_CPU=12000000
AVR_WRITER=hidspx
AVR_WRITER_OPT=-ph

APP=sensor
SRCS=sensor.c
ELF=$(APP).elf

all: $(APP).hex

$(APP).hex: $(SRCS)
	$(CC) $(CFLAGS) -mmcu=$(TARGET) -D F_CPU=$(F_CPU) -o $(ELF) $(SRCS)
	avr-objcopy -O ihex -R .eeprom $(ELF) $@
	rm $(ELF)

write: $(APP).hex
	@echo writing $? ...
	$(AVR_WRITER) $(AVR_WRITER_OPT) $?

clean:
	rm *.elf *.hex

