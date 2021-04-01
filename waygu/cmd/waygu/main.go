package main

import (
	"io"
	"log"
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
)

// https://stackoverflow.com/questions/11356330/how-to-get-cpu-usage
func taskManager() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")

	if err != nil {
		return
	}

	lines := strings.Split(string(contents), "\n")

	for _, line := range(lines) {
		fields := strings.Fields(line)

		if fields[0] == "cpu" {
			numFields := len(fields)

			for i := 1; i<numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)

				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}

				total += val 
				if i == 4 {
					idle = val
				}
			}

			return
		}
	}

	return
}

func remindScheduler() {
	re := regexp.MustCompile("([0-9]+)\\s?([A-Za-z]+)")
	
}

func main() {
	remindHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "reminder endpoint\n")
	}

	http.HandleFunc("/remind", remindHandler)
		log.Println("Listing for requests at http://localhost:8000/remind")
	
	
	cpuHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "task manager endpoint\n")

		idle0, total0 := taskManager()
		time.Sleep(3 * time.Second)
		idle1, total1 := taskManager()

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)
		cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks 

		fmt.Printf("%f\n", cpuUsage)
		msg := "CPU usage is "
		res := fmt.Sprint(msg, float64(cpuUsage))
		
		// out := "CPU usage is " + float64(cpuUsage) + "[busy: " + float64(totalTicks-idleTicks)
		// 	+ ", total: " + float64(totalTicks) + "\n]"
		io.WriteString(w, res)
	}

	http.HandleFunc("/cpu", cpuHandler)
		log.Println("Listing for requests at http://localhost:8000/cpu")

	log.Fatal(http.ListenAndServe(":8000", nil))
}