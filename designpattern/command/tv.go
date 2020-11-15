package main

import "fmt"

type tv struct {
	isOn   bool
	volume int
}

func (t *tv) on() {
	t.isOn = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isOn = false
	fmt.Println("Turning tv off")
}

func (t *tv) increaseVolume() {
	if t.isOn {
		if t.volume >= 0 && t.volume < 100 {
			t.volume++
			fmt.Printf("Increase volume to %d\n", t.volume)
		} else {
			fmt.Printf("Max volume.")
		}
	} else {
		fmt.Println("TV is off, cannot change volume")
	}
}

func (t *tv) decreaseVolume() {
	if t.isOn {
		if t.volume > 0 && t.volume <= 100 {
			t.volume--
			fmt.Printf("Decrease volume to %d\n", t.volume)
		} else {
			fmt.Printf("Min volume.")
		}
	} else {
		fmt.Println("TV is off, cannot change volume")
	}
}
