package main

import "fmt"

// IProcess interface
type IProcess interface {
	process()
}

// ProcessClass struct
type ProcessClass struct{}

// ProcessClass method process
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// ProcessDecorator struct
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// ProcessDecorator class method process
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Println("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

func main() {
	process := &ProcessClass{}
	decorator := &ProcessDecorator{}
	decorator.process()

	decorator.processInstance = process
	decorator.process()
}
