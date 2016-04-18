# Housing-Density-Distr_Sys_2016
As you know Hadoop is an open source project for processing large datasets in parallel with the use of low level commodity machines. We have discussed MapReduce during lecture and covered setting up the environment along with an example in recitation. Now it’s time to write your own MapReduce Program.  You are provided with United States Census data, where you can download the Zip Code Tabulation Areas Gazetteer File (1.1MB). Your goal is to find the population-weighted housing density in the northeast, northwest, southeast, and southwest, where the population-weighted housing density is simply the amount of houses per person per square kilometer. The smaller the result, the higher the desnsity. This can be a factor in determining housing price, although we ignore other significant influencing factors such as the local economy, public infrastructure, crime, community, environment, etc. Here, population density is simply based on population count / land area.  For example, for the region identified by the zip code 08854 (Piscataway, NJ), we have:  population density = 1147.5 people / km2 population-weighted housing density = 15.4706 houses / person / km2 For the region identified by the zip code 90822 (Farmington Hills, Michigan), we have::  population density = 1147.5 people / km2 population-weighted housing density = 9.8555 houses / person / km2 We see that Piscataway has a population density that is about the same as Farmington Hills, Michigan but a population-weighted housing density that is approximately 1.6 times greater. It happens that the average house price in Piscataway is $343K while the average house price of Farmington Hills is $214K. The average house price in Piscataway is around 1.57 times greater, although we will not always see this strong correlation between house price and density.  Solve the problem using map-reduce. Along with your program, you should briefly explain how the input is mapped into (key, value) pairs by the map phase. For instance, specify what is the key and what is the associated value in each pair, and how the key(s) and value(s) are computed. Then explain how the (key, value) pairs produced by the map stage are processed by the reduce phase. If the job cannot be done in a single map-reduce pass, describe how it would be structured into two or more map-reduce jobs with the output of the first job becoming input to the next one(s).  This assignment is pretty straightforward, but you should start early and allocate ample time if you have never used Hadoop to develop and run MapReduce jobs.  The write-up focuses on the Java implementation, but you are feel free to do this assignment in Python, Go, or C++. Just anticipate some necessary translation of your code.
