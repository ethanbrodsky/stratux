// Package motordriver provides control of WaveShare 2x3A PCA9685-based motor driver hat (15364)

package buzzball

import (
        "log"
        "math"
	"github.com/kidoman/embd"
//	"github.com/kidoman/embd/controller/pca9685"  USED OUR OWN VERSION
	"../pca9685"
)

const (
	i2cAddr    = 0x40        // default base address (configurable from 0x40..0x5F)
        pwmFreq    = 100         // Hz

	motor1_speed = 0
	motor1_dirR  = 1
	motor1_dirF  = 2
	
	motor2_speed = 5
	motor2_dirR  = 3
	motor2_dirF  = 4
)

type MotorRegs struct {
  speed int
  dirF  int
  dirR  int
}

var motor_regs = []MotorRegs	{ 
					MotorRegs { speed:motor1_speed, dirF:motor1_dirF, dirR:motor1_dirR, }, 
					MotorRegs { speed:motor2_speed, dirF:motor2_dirF, dirR:motor2_dirR, },
				}


type MotorDriver struct {
	*pca9685.PCA9685
}

var motordriver *MotorDriver

func MotorDriver_Init() {
	motordriver = &MotorDriver{pca9685.New(embd.NewI2CBus(1), i2cAddr)}
	motordriver.Freq = pwmFreq

	// ~100 Hz
	motordriver.SetPrescaler(67);
}


func MotorDriver_Close() {

	// Stop both motors
	motordriver.SetPwm(0, 0, 0)
	motordriver.SetPwm(1, 0, 0)

	motordriver.Close()
}

// motor: 0 or 1;  speed: 0..1
func MotorDriver_SetSpeed(motor int, speed float64) {
	var speedword int = int(math.Round(speed*4095))
  
	err := motordriver.SetPwm(motor_regs[motor].speed, 0, speedword)

        log.Printf("xxxx SetPwm: %d %d \n", motor_regs[motor].speed, speedword);

	if err != nil {
		log.Printf("xxxx ERROR %s!", err)
	}
}

func MotorDriver_SetDirection(motor int, fForward bool) {
	if (fForward) {
		motordriver.SetPwm(motor_regs[motor].dirF, 0, 4095);
		motordriver.SetPwm(motor_regs[motor].dirR, 0, 0);
	} else {
		motordriver.SetPwm(motor_regs[motor].dirF, 0, 0);
		motordriver.SetPwm(motor_regs[motor].dirR, 0, 4095);
	}
}
