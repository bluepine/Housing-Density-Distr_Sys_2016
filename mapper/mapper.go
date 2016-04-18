package main

import "fmt"
import "bufio"
import "regexp"
import "os"
import "strconv"


func main() {
	/* Word regular experssion. */
	re, _ := regexp.Compile("[a-zA-Z0-9]+")
	reader := bufio.NewReader(os.Stdin)
	var location string = ""

	//seen : make(map[string]int)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			
			break
		}
		fields := re.FindAll(line, -1)

		data := string(fields[1])+"\t"+string(fields[2])+"\t"+string(fields[3])
		
		latitude := string(fields[9])+"."+string(fields[10])
		longitude := "-"+string(fields[11])+"."+string(fields[12])

		
		
		newLat,_ := strconv.ParseFloat(latitude,64)
		newLong,_ := strconv.ParseFloat(longitude,64)
		
		if newLat > 39.833333 && newLong < -98.583333{
			location = "northwest"
		}else if newLat > 39.833333 && newLong > -98.583333{
			location = "northeast"
		}else if newLat < 39.833333 && newLong < -98.583333{
			location = "southwest"
		}else if newLat < 39.833333 && newLong > -98.583333{
			location = "southeast"
		}

		if len(location) > 1{
		fmt.Printf("%s\t%s\n",location,data)
	}
		
		

	}
}