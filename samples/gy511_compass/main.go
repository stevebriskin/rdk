package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/viamrobotics/robotcore/sensor/compass/gy511"
	"github.com/viamrobotics/robotcore/serial"

	"github.com/edaniels/golog"
)

func main() {
	var calibrate bool
	flag.BoolVar(&calibrate, "calibrate", false, "calibrate compass")
	flag.Parse()

	devices, err := serial.SearchDevices(serial.SearchFilter{Type: serial.DeviceTypeArduino})
	if err != nil {
		golog.Global.Fatal(err)
	}
	if len(devices) == 0 {
		golog.Global.Fatal("no applicable device found")
	}
	sensor, err := gy511.New(devices[0].Path)
	if err != nil {
		golog.Global.Fatal(err)
	}

	if calibrate {
		sensor.StartCalibration()
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		sensor.StopCalibration()
	}

	for {
		time.Sleep(100 * time.Millisecond)
		readings, err := sensor.Readings(context.Background())
		if err != nil {
			golog.Global.Errorw("failed to get sensor reading", "error", err)
			continue
		}
		golog.Global.Infow("readings", "data", readings)
		heading, err := sensor.Heading()
		if err != nil {
			golog.Global.Errorw("failed to get sensor heading", "error", err)
			continue
		}
		golog.Global.Infow("heading", "data", heading)
	}
}
