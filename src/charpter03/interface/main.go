package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name  string
	Brand string
}

func (p Phone) Start() {
	fmt.Println("Phone", p.Name, "is starting")
}
func (p Phone) Stop() {
	fmt.Println("Phone", p.Name, "is stopping")
}
func (p Phone) Call() {
	fmt.Println("Ring A Ring A Ring聆听爱的和弦")
}

type Camera struct {
	Name  string
	Brand string
}

func (c Camera) Start() {
	fmt.Println("Camera", c.Name, "is starting")
}
func (c Camera) Stop() {
	fmt.Println("Camera", c.Name, "is stopping")
}

type Computer struct {
	Name  string
	Brand string
}

func (c Computer) Working(usb Usb) {
	fmt.Println("Computer", c.Name, "is working with", usb)
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

}
func (c Computer) Stop(usb Usb) {
	fmt.Println("Computer", c.Name, "is stopping with", usb)
	usb.Stop()
}

func main() {
	var usb [3]Usb
	usb[0] = Phone{Name: "iPhone", Brand: "Apple"}
	usb[1] = Camera{Name: "EOS", Brand: "Canon"}
	usb[2] = Phone{Name: "Galaxy", Brand: "Samsung"}
	computer := Computer{Name: "ThinkPad", Brand: "Lenovo"}
	for _, device := range usb {
		computer.Working(device)
		computer.Stop(device)
		fmt.Println("-----")
	}
}
