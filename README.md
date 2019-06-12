#***GRPC***

**The project requires the following configuration to run:**

1)Goland IDE version 2019.1.3
2)Requires Go version 1.6 or above
3)Protocol Buffers V3
4)GRPC
5)Intellij Idea Ultimate
6)gradle version 3.0+ 

The program performs the following tasks using grpc:

Calculate average of a number.
Check whether a number is  Armstrong number or not.
Calculate prime factors of a number.
Calculate average of a given set of numbers.
Determine whether given numbers are odd or even.


This project is based on GRPC communication between:
a)GO SERVER AND GO CLIENT
b)GO SERVER AND JAVA CLIENT

***To run the program with Go server and Go client***

Open Factorial directory as project in GoLand IDE.
Run main.go file
Note:-> Before running the program please chech the imports in the following go files first 1.) client.go in Golang -> client directory 2.) server.go and calculator_handler in Golang -> server directory

***OUTPUT***
![grpc1](https://user-images.githubusercontent.com/40175918/59329274-24522d00-8d0c-11e9-9f0b-6a86d6f9822c.png)

***To run the program with Go server and Java client***

Open Golang directory as project in Goland IDE.
Run only grpc server from main.go file
Open grpcjava directory as project in IntelliJ IDE.
After opening the project, open the project directory in terminal and run the following command "gradle clean build"
After running gradle build, in grpcjava navigate to src/main/java and run client.java



***OUTPUT***
![grpc2](https://user-images.githubusercontent.com/40175918/59329871-9119f700-8d0d-11e9-869a-dc36537e77a0.png)





