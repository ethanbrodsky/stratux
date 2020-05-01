// Package buzzball makes the buzzing happen

package buzzball

import (
	"log";
//        "math";
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

func BuzzBall_Update(dt float64, valid bool, slipskid float64) {
//        totaltime += dt;
//        flag := math.Mod(totaltime, 5) < 2
//        log.Printf("xxxx %.3f %t %+.3f --- %.1f %t \n", dt, valid, slipskid, totaltime, flag);

        log.Printf("xxxx %.3f %t %+.3f \n", dt, valid, slipskid);

        // Crudest of crude hacks to make the motors do their thing (not sure if this is flipped L/R yet)
	SetMotor(MotorL, slipskid > +1)
	SetMotor(MotorR, slipskid < -1)
}
