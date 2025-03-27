package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Input file
	inputFile := "domains.txt"

	// Output file for IP addresses
	outputFile := "output.txt"

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create or open the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" {
			continue
		}

		ips, err := net.LookupIP(domain)
		if err != nil {
			fmt.Fprintf(writer, "%s: Failed to resolve\n", domain)
			continue
		}

		for _, ip := range ips {
			fmt.Fprintf(writer, "%s: %s\n", domain, ip.String())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	writer.Flush()
	fmt.Println("IP addresses have been saved to output.txt")
}
