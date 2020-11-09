package main

import "time"

func main() {
	loadInterval := time.Minute
	for {
		println("Start of Cycle")

		println("Ramping up")
		rampUpCpu(loadInterval)

		println("Ramping down")
		rampDownCpu(loadInterval)
		
		println("End of Cycle")
	}
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
