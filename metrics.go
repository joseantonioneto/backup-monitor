package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get backup size in bytes
	cmd := exec.Command("du", "-s", "/home/a/Área de Trabalho/backup")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error getting backup size:", err)
		return
	}
	backupSizeStr := strings.Split(string(output), "\t")[0]
	backupSize, err := strconv.Atoi(backupSizeStr)
	if err != nil {
		fmt.Println("Error converting backup size to integer:", err)
		return
	}

	// Convert to gigabytes
	backupSizeGb := float64(backupSize) / (1024 * 1024 * 1024)

	// Get disk usage
	cmd = exec.Command("df", "-h", "/")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error getting disk usage:", err)
		return
	}
	lines := strings.Split(string(output), "\n")
	usageLine := strings.Fields(lines[1])[3] // Assuming the second line contains usage information

	// Get percentage use
	percentUse := strings.Fields(lines[1])[4]
	if err != nil {
		fmt.Println("Error converting percentage use to integer:", err)
		return
	}

	// Get current date
	presentDate := time.Now().Format("2006-05-02 15:04:05")

	// Print results
	fmt.Printf("Backup size: %.2f GB\n", backupSizeGb)
	fmt.Println("Disk usage:", usageLine)
	fmt.Println("Percentage use:", percentUse)
	fmt.Println("Present date:", presentDate)
}
