package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temperature, err := getCPUtemperature()
		if err != nil {
			http.Error(w, "Failed to read CPU temperature", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "CPU Temperature: %s\n", temperature)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func getCPUtemperature() (string, error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	temperature := strings.TrimSpace(string(output))
	// Convert millidegrees Celsius to degrees Celsius
	temperatureC := fmt.Sprintf("%.2f", float64(mustParseInt(temperature))/1000.0)
	return temperatureC + "°C", nil
}

func mustParseInt(s string) int {
	i := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			i = i*10 + int(c-'0')
		}
	}
	return i
}
