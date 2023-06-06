package limo

import (
	"context"
	"testing"

	"github.com/edaniels/golog"
	"go.viam.com/test"

	"go.viam.com/rdk/resource"
)

func TestLimoBaseConstructor(t *testing.T) {
	ctx := context.Background()
	logger := golog.NewTestLogger(t)

	c := make(chan []uint8, 100)

	_, err := CreateLimoBase(context.Background(), resource.Config{ConvertedAttributes: &Config{}}, logger)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, err.Error(), test.ShouldStartWith, "drive mode must be defined")

	cfg := &Config{
		DriveMode: "ackermann",
		TestChan:  c,
	}

	baseBase, err := CreateLimoBase(context.Background(), resource.Config{ConvertedAttributes: cfg}, logger)
	test.That(t, err, test.ShouldBeNil)
	base, ok := baseBase.(*limoBase)
	test.That(t, ok, test.ShouldBeTrue)

	test.That(t, controllers[defaultSerial], test.ShouldNotBeNil)
	width, _ := base.Width(ctx)
	test.That(t, width, test.ShouldEqual, 172)
	base.Close(ctx)

	test.That(t, controllers, test.ShouldBeEmpty)
}
