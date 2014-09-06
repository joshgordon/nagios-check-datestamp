package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {


	if len(os.Args) != 4 { 
		fmt.Println("UNKNOWN - Invalid command line args")
		os.Exit(4)
	}
	//Read in the file specified by the first command line arg. 
	filename := os.Args[1]
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		//Provide an error message if we can't read in the file. 
		fmt.Println("UNKNOWN - Error reading date file")
		fmt.Println(err)
		os.Exit(4)
	}

	//Current unix time. 
	timenow := time.Now().Unix()


	//Parse the file. This expects a newline at the end of the file, so it
	//cuts that off. 
	lastRun, err := strconv.ParseInt(string(file[0:len(file)-1]), 10, 64)
	if err != nil {

		//provide an error if we can't parse the date file. 
		fmt.Println("UNKNOWN - Error parsing date file")
		fmt.Println(err)
		os.Exit(4)
	}

	//set the warning and critical levels.
	var warning int64
	var critical int64

	//parse the command line args for number of seconds to warn and
	//go critical. 
	warning, errw := strconv.ParseInt(os.Args[2], 10, 64)
	critical, errc := strconv.ParseInt(os.Args[3], 10, 64)
	
	//Check for errors with command line args. 
	if errw != nil && errc != nil { 
		fmt.Println("UNKNOWN - Error with command line args")
		fmt.Println(errc)
		fmt.Println(errw)
		os.Exit(4)
	}

	//compute the time since the last run.
	timeSinceLastRun := timenow - lastRun

	//format the string for telling the time since the last run.
	//It's ugly, but it works. 
	lastRunStr := fmt.Sprintf("%d seconds (%d days, %02d:%02d:%02d) since last run", timeSinceLastRun, timeSinceLastRun/86400, (timeSinceLastRun%86400)/3600, (timeSinceLastRun%3600)/60, timeSinceLastRun%60)

	// Critical level. 
	if timeSinceLastRun > critical {
		fmt.Println("CRITICAL - " + lastRunStr)
		os.Exit(2)
	}

	//Warning level. 
	if timeSinceLastRun > warning {
		fmt.Println("WARNING - " + lastRunStr)
		os.Exit(1)
	}

	//if we get here - we're all good.
	fmt.Println("OK - " + lastRunStr)
	os.Exit(0)

}
