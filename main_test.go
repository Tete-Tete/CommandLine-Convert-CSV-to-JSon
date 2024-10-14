package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Testing monitor CPU memory
	cputest, err := os.Create("cpu_profile.prof")
	if err != nil {
		t.Fatalf("Sry, couldn't create CPU profile: %v, plz check again", err)
	}
	pprof.StartCPUProfile(cputest) // Start CPU profiling
	defer pprof.StopCPUProfile()   // End CPU profiling when this function is done

	// Create a temporary CSV file for testing
	testFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Sry, couldn't create temporary CSV file for testing: %v, plz check again", err)
	}
	defer os.Remove(testFile.Name()) // Delete the temporary CSV file when this function is done

	// Write random data to the temporary CSV File for testing
	_, err = testFile.WriteString("name,group,age\nYat,5,22\nChris,5,23\n")
	if err != nil {
		t.Fatalf("Sry, couldn't write to CSV file: %v, plz check again", err)
	}
	testFile.Close()

	// Create a temporary output file
	outputFile, err := os.CreateTemp("", "output.json")
	if err != nil {
		t.Fatalf("Sry, couldn't create output file: %v, plz check again", err)
	}
	defer os.Remove(outputFile.Name()) // Delete the output file when this function is done

	// Run main function with test arguments
	os.Args = []string{"cmd", testFile.Name(), outputFile.Name()}
	main()

	// Log memory usage after running the main function
	logMemoryUsage()

	// Check the output file
	fileCheck, err := os.Stat(outputFile.Name())
	if err != nil {
		t.Fatalf("Sry, couldn't get output file info: %v, plz check again", err)
	}
	if fileCheck.Size() == 0 {
		t.Fatalf("Sry, the output file is empty, plz check again")
	}
}

// Log memory usage function
func logMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	log.Printf("TotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	log.Printf("Sys = %v MiB", m.Sys/1024/1024)
}
