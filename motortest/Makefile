DIR_SRC = ./src
DIR_BIN = ./bin

SRC_C = $(wildcard ${DIR_SRC}/*.c)
SRC_O = $(patsubst %.c,${DIR_BIN}/%.o,$(notdir ${SRC_C}))

TARGET = motortest

CC = gcc -std=c99

DEBUG = -g -O0 -Wall
CFLAGS += $(DEBUG)

LIB = -lbcm2835 -lm

${TARGET} : ${SRC_O}
	$(CC) $(CFLAGS) $(SRC_O) -o $@ $(LIB)

${DIR_BIN}/%.o : $(DIR_SRC)/%.c
	$(CC) $(CFLAGS) -c  $< -o $@ $(LIB)

	
clean :
	rm $(DIR_BIN)/*.* 
	rm $(TARGET) 
