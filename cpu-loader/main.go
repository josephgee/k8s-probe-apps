package main

import "time"

func main() {
	loadInterval := 3 * time.Second
	rampUpCpu(loadInterval)
	rampDownCpu(loadInterval)
}

func rampUpCpu(loadInterval time.Duration) {
	for interval := 0; interval <= 100; interval += 10 {
		loadCPU(loadInterval, interval)
	}
}

func rampDownCpu(loadInterval time.Duration) {
	for interval := 90; interval >= 0; interval -= 10 {
		loadCPU(loadInterval, interval)
	}
}


func loadCPU(loadDuration time.Duration, loadPercent int) {
	stopTime := time.Now().Add(loadDuration)
	spinTime := time.Duration(loadPercent)*time.Millisecond
	sleepTime := (100-time.Duration(loadPercent))*time.Millisecond

	for stopTime.Sub(time.Now()).Nanoseconds() > 0 {
		spin(spinTime)
		time.Sleep(sleepTime)
	}
}

func spin(duration time.Duration) {
	waitTime := time.Now().Add(duration)
	now := time.Now()
	for waitTime.Sub(now).Nanoseconds() > 0 {
		now = time.Now()
	}
}
