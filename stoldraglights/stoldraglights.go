package stoldraglights

import (
	"log";
	"math";
	"../ws2811";
)

const (
	ledPin        = 10
	ledCount      = 4
	ledBrightness = 255
)

const (
	ColorRed   = 0x00200000
	ColorGreen = 0x00002000
	ColorBlue  = 0x00000020
	ColorOff   = 0x00000000
)

func LEDs_AllOff() {
	for ii := 0; ii < ledCount; ii++ {
		ws2811.SetLed(ii, ColorOff)
	}

	ws2811.Render()
}

func LED_Red(fEnabled bool) {
	if (fEnabled) {
		ws2811.SetLed(0, ColorRed)
	} else {
		ws2811.SetLed(0, ColorOff)
	}
}

func LED_Green(fEnabled bool) {
	if (fEnabled) {
		ws2811.SetLed(1, ColorGreen)
	} else {
		ws2811.SetLed(1, ColorOff)
	}
}

func StolDragLights_Init() {

	defer StolDragLights_Close()

	err := ws2811.Init(ledPin, ledCount, ledBrightness)
	if err != nil {
		log.Println(err)
	}

//	LEDs_AllOff()

}

func StolDragLights_Close() {

	ws2811.Fini()

}

var totaltime float64 = 0

var stoptime float64 = 0

func StolDragLights_Update(dt float64, gpsGroundSpeed float64, gpsTrueCourse float64, gpsTurnRate float64, /*gpsLastTime Time,*/ ahrsPitch float64, ahrsRoll float64, ahrsHeading float64, ahrsTurnRate float64, ahrsGLoad float64, ahrsSlipSkid float64, ahrsStatus uint8 ) {
        totaltime += dt;

        var x = ahrsSlipSkid

        var fStopped = math.Abs(x) < 5

	if (fStopped) {
		stoptime += dt
	} else {
		stoptime = 0
	}
        
        var fFullStop = stoptime > 2

	if (totaltime > 30) {
//		LED_Green( fStopped )
//        	LED_Red  ( fFullStop )
	}

        log.Printf(
	           "StolDragLights %.1f %.1f * %.2f %.2f %.2f * %+.2f %+.2f %.2f  %+.2f %+.2f %+.2f  %d  %t %t %.1f \n", 
	           dt, totaltime, 
	           gpsGroundSpeed, gpsTrueCourse, gpsTurnRate, /*gpsLastTime,*/
	           ahrsPitch, ahrsRoll, ahrsHeading,   ahrsTurnRate, ahrsGLoad, ahrsSlipSkid,   ahrsStatus,
	           fStopped, fFullStop, stoptime);
}
