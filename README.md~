# Housing-Density-Distr_Sys_2016
As you know Hadoop is an open source project for processing large datasets in parallel with the use of low level commodity machines. We have discussed MapReduce during lecture and covered setting up the environment along with an example in recitation. Now it’s time to write your own MapReduce Program.  You are provided with United States Census data, where you can download the Zip Code Tabulation Areas Gazetteer File (1.1MB). Your goal is to find the population-weighted housing density in the northeast, northwest, southeast, and southwest, where the population-weighted housing density is simply the amount of houses per person per square kilometer. The smaller the result, the higher the desnsity. This can be a factor in determining housing price, although we ignore other significant influencing factors such as the local economy, public infrastructure, crime, community, environment, etc. Here, population density is simply based on population count / land area.  For example, for the region identified by the zip code 08854 (Piscataway, NJ), we have:  population density = 1147.5 people / km2 population-weighted housing density = 15.4706 houses / person / km2 For the region identified by the zip code 90822 (Farmington Hills, Michigan), we have::  population density = 1147.5 people / km2 population-weighted housing density = 9.8555 houses / person / km2 We see that Piscataway has a population density that is about the same as Farmington Hills, Michigan but a population-weighted housing density that is approximately 1.6 times greater. It happens that the average house price in Piscataway is $343K while the average house price of Farmington Hills is $214K. The average house price in Piscataway is around 1.57 times greater, although we will not always see this strong correlation between house price and density.

#Testing Locally

	Build executables
	      go build reducer.go
	      go build mapper.go
	Pipe Data between the programs
	     Data_File.txt | ./mapper | ./reducer


#Testing on Hadoop Cluster
 	The executables(goLang code) and Hadoop-Streaming-Jar File sit on the local machine and the data file(s) are on the hdfs.  
	

	