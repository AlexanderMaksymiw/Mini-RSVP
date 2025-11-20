package main

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
    config := &serial.Config{
        Name: "COM7", // change this
        Baud: 9600,   // some OBD2 dongles want 9600 or 38400
        ReadTimeout: time.Second * 2,
    }

    port, err := serial.OpenPort(config)
    if err != nil {
        log.Fatal("open error:", err)
    }

    // Send a reset command. Every ELM327 clone supports this.
    _, err = port.Write([]byte("ATZ\r"))
    if err != nil {
        log.Fatal("write error:", err)
    }

    buf := make([]byte, 256)
    n, err := port.Read(buf)
    if err != nil {
        log.Fatal("read error:", err)
    }

    log.Printf("Response: %s", string(buf[:n]))
}
