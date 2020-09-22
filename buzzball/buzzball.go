// Package buzzball makes the buzzing happen

package buzzball

import (
	"log";
        "math";
)


const (
	MotorL	=	0	// MA
	MotorR	=	1	// MB
	speed	=	1.0
)

func BuzzBall_Init() {
	MotorDriver_Init()

	MotorDriver_SetDirection(MotorL, true);
	MotorDriver_SetDirection(MotorR, true);
}

func SetMotor(motor int, fEnabled bool) {
	if (fEnabled) {
		MotorDriver_SetSpeed(motor, speed)
	} else {
		MotorDriver_SetSpeed(motor, 0.0)
	}
}

func BuzzBall_Close() {
 	MotorDriver_Close()
}

var totaltime float64 = 0
const g = 9.8;  // m/s^2

// Reference Frame:
//   +Roll is bank to right
//  -Slip/Skid bank to the right without right rudder (ball to the L)
//   +Y means aircraft is accelerating to L (ball to the R)
//  AccelY is -g*sin(SlipSkid)

func BuzzBall_Update(dt float64, valid bool, slipskid float64) {
        totaltime += dt;

        var accelY  = -g * math.Sin(slipskid * math.Pi / 180);
        var ballPos = accelY * -math.Sqrt(3)/2; 

        log.Printf("xxxx %.3f %t %+.3f %+.2f %+.2f \n", dt, valid, slipskid, accelY, ballPos);

        // Crudest of crude hacks to make the motors do their thing (not sure if this is flipped L/R yet)
	SetMotor(MotorL, ballPos > +0.375);
	SetMotor(MotorR, ballPos < -0.375);
}
