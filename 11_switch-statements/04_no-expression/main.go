package main

import "fmt"

func main() {
	stop := false
	for stop == false{
		var myFriendsName string
		fmt.Scanln(&myFriendsName)

		switch {
		case len(myFriendsName) == 2:
			fmt.Println("Wassup my friend with name of length 2")
		case myFriendsName == "Tim":
			fmt.Println("Wassup Tim")
		case myFriendsName == "Jenny":
			fmt.Println("Wassup Jenny")
		case myFriendsName == "Marcus", myFriendsName == "Medhi":
			fmt.Println("Your name is either Marcus or Medhi")
		case myFriendsName == "Julian":
			fmt.Println("Wassup Julian")
		case myFriendsName == "Sushant":
			fmt.Println("Wassup Sushant")
		case myFriendsName == "Quit":
			fmt.Println("Quitting...")
			stop = true
		default:
			fmt.Println("nothing matched; this is the default")
		}
	}
}

/*
	expression not needed
	-- if no expression provided, go checks for the first case that evals to true
	-- makes the switch operate like if/if else/else
	cases can be expressions
 */