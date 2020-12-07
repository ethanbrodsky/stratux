#include <ctype.h>
#include <stdio.h>      //printf()
#include <stdlib.h>     //exit()
#include <signal.h>
#include "DEV_Config.h"
#include "MotorDriver.h"

void  Handler(int signo)
{
    //System Exit
    printf("\r\nHandler:Motor Stop\r\n");
//    Motor_Stop(MOTORA);
//    Motor_Stop(MOTORB);
    DEV_ModuleExit();

    exit(0);
}

int main(int argc, char *argv[])
{
    //1.System Initialization
    if(DEV_ModuleInit())
        exit(0);
    

    if (argc != 3)
      {
        printf("SYNTAX: '%s <MOTOR> <SPEED>' \n", argv[0]);
        printf("        <MOTOR>   'A' or 'B' \n");
        printf("        <SPEED>   0..100     \n");
        exit(-1); 
      }

    int motor = MOTORA;
    if (tolower(argv[1][0]) == 'a')
      motor = MOTORA;
    else if (tolower(argv[1][0]) == 'b')
      motor = MOTORB;
    else
      {
        printf("Invalid Motor '%s' - should be A or B \n", argv[0]);
        exit(-1);
      }

    int direction = +1;

    int speed = 0;
    speed = atoi(argv[2]);
    if (speed < 0)
      {
        speed *= -1;
        direction *= -1;
      }

    if (speed > 100)
      speed = 100;

    //1. System Initialization
    if (DEV_ModuleInit())
      exit(0);

    printf("Motor %c \n", (motor == MOTORA) ? 'A' : 'B');
    printf("Dir   %+d \n", direction);
    printf("Speed %d \n", speed);

    //2. Motor Initialization
    Motor_Init();

    printf("Motor_Run\r\n");
    Motor_Run(motor, (direction==+1) ? FORWARD : BACKWARD, speed);

    //3.System Exit
//    DEV_ModuleExit();
    return 0;
}

