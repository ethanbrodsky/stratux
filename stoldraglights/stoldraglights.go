package stoldraglights

import (
	"fmt";
	"log";
	"math";
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go";
	"os/exec";
)

const (
	ColorRed   = 0x00200000
	ColorGreen = 0x00002000
	ColorBlue  = 0x00000020
	ColorOff   = 0x00000000
)

type wsType struct {
	ws2811 *ws281x.WS2811
}

func (ws *wsType) init() error {
	err := ws.ws2811.Init()
	if err != nil {
		return err
	}

	return nil
}

func (ws *wsType) close() {
	ws.ws2811.Fini()
}

func (ws *wsType) AllOff() {
	for ii := 0; ii < ledCount; ii++ {
		ws.ws2811.Leds(0)[ii] = ColorOff
	}

	ws.ws2811.Render()
}

func (ws *wsType) SetRed(fEnabled bool) {
	if (fEnabled) {
		ws.ws2811.Leds(0)[0] = ColorRed
	} else {
		ws.ws2811.Leds(0)[0] = ColorOff
	}
}

func (ws *wsType) SetGreen(fEnabled bool) {
	if (fEnabled) {
		ws.ws2811.Leds(0)[1] = ColorGreen
	} else {
		ws.ws2811.Leds(0)[1] = ColorOff
	}

}

func (ws *wsType) Render() {
	ws.ws2811.Render()
	ws.ws2811.Wait()
}

var ws *wsType

const (
	ledPin        = 10
	ledCount      = 4
	ledBrightness = 255
)

func StolDragLights_Init() {
//DONE BY USER	defer StolDragLights_Close()

/*
	opt := ws281x.DefaultOptions
	opt.Channels[0].Brightness = ledBrightness
	opt.Channels[0].LedCount   = ledCount
	opt.Channels[0].GpioPin    = ledPin

        log.Printf("xxx1 \n");

	ws2811_new, err := ws281x.MakeWS2811(&opt)
	if err != nil {
		log.Println(err)
	}

        log.Printf("xxx2 \n");

	ws := wsType{
		ws2811: ws2811_new,
	}

        log.Printf("xxx3 \n");

	err = ws.init()
	if err != nil {
		log.Println(err)
	}

        log.Printf("xxx4 \n");
	
//	ws.AllOff()
*/
        log.Printf("xxx5 \n");
}

func StolDragLights_Close() {
//	ws.close()
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
//		ws.SetRed  ( fStopped )
//        	ws.SetGreen( fFullStop )
//		ws.Render()
	}

	bit_r := 0
	bit_g := 0

	if (fStopped) {
		bit_r = 1
	}

	if (fFullStop) {
		bit_g = 1
	}

        str := fmt.Sprintf("%d%d", bit_r, bit_g)
        cmd := exec.Command("/usr/local/bin/ledset", str)
        err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

        log.Printf(
	           "StolDragLights %.1f %.1f * %.2f %.2f %.2f * %+.2f %+.2f %.2f  %+.2f %+.2f %+.2f  %d  %t %t %.1f %s \n", 
	           dt, totaltime, 
	           gpsGroundSpeed, gpsTrueCourse, gpsTurnRate, /*gpsLastTime,*/
	           ahrsPitch, ahrsRoll, ahrsHeading,   ahrsTurnRate, ahrsGLoad, ahrsSlipSkid,   ahrsStatus,
	           fStopped, fFullStop, stoptime, str);
}
