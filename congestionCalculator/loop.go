package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("./darknet", "detect", "cfg/yolov3.cfg", "yolov3.weights", "../busImages/posted_bus.png").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
