package main

import (
	"datecounter/utils"
	"flag"
	"fmt"
	"os"
)

func main() {

	//use flag to get args from command
	from := flag.String("from", "01/01/2001", "From is use to start date calculate the distance days")
	to := flag.String("to", "03/01/2001", "To is use to End date  calculate the distance days .")
	flag.Parse()
	//chek args not empty
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	//check form and to is not empty
	if *to == "" || *from == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	//check date is valid
	f := utils.IsValidDate(*from, *to)
	fmt.Printf("%d days", f )

}
