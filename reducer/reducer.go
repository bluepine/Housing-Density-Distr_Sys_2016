package main 

import ("fmt"
		"strconv"
		"bufio"
		"os"
		"regexp"
		)

func main() {
	re, _ := regexp.Compile("[a-zA-Z0-9]+")
	reduced := make (map[string]int)
	reader :=bufio.NewReader(os.Stdin)
	var nwCount int = 0;
	var neCount int = 0;
	var swCount int = 0;
	var seCount int = 0;



	for{
		line,_,err := reader.ReadLine()
		if err != nil {
			break
		}

		data := re.FindAll(line,-1)
		
		region := string(data[0])
	    pop := string(data[1])
	    houses := string(data[2])
	    land := string(data[3])
	    pop1,_ := strconv.Atoi(pop)
	    houses1,_ := strconv.Atoi(houses)
	    land1,_ := strconv.Atoi(land)

	  	if houses1 != 0 && pop1 != 0 && land1 != 0{
	    reduced[region] += (houses1 *land1)/(pop1*1000000)
	
	    if region == "northwest"{
	    	nwCount++

	    	}else if region == "northeast"{
	    		neCount++

	    		}else if region == "southwest"{
	    			swCount++


	    			}else if region == "southeast"{
	    				seCount++

	    			}
	    		}

	}

	for key,value := range reduced{
		if key == "southeast"{
		fmt.Printf("%s\t%d\n",key,value/seCount)
		
		}else if key == "southwest"{
			fmt.Printf("%s\t%d\n",key,value/swCount)

			}else if key == "northeast"{
				fmt.Printf("%s\t%d\n",key,value/neCount)


				}else if key == "northwest"{
					fmt.Printf("%s\t%d\n",key,value/nwCount)


				}
	}




}