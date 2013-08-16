//Replace values relatively to the current
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	//Read file
	flag.Parse() //Parse command line arguments ie. the file names
	fp1 := strings.Replace(flag.Arg(0), "\\", "/", -1)
	fi1, err := os.Open(fp1)
	if err != nil {
		panic(err)
	}
	defer fi1.Close()
	r1 := bufio.NewReader(fi1)
	
	//Compile regular expressions
	rxValue, _ := regexp.Compile("Line[[0-9]*]")
	rxRel, _ := regexp.Compile("[0-9]+")
	
	//Go over lines
	l1, err := r1.ReadString('\n')
	for err == nil {
		//fmt.Print(l1)
		if rxValue.MatchString(l1) {	
		 //fmt.Print(l1)
		val := rxValue.FindString(l1)
		rel := rxRel.FindString(val)
		relInt, _ := strconv.Atoi(rel)
		new1 := relInt - 1
		new2 := strconv.Itoa(new1)
		fmt.Print(rxValue.ReplaceAllString(l1, "Line[" + new2 + "]"))

		} else {
		fmt.Print(l1)
		}
				l1, err = r1.ReadString('\n')

		
	}
	if err != nil && err != io.EOF {
		fmt.Print(err)
	}

	
		}
		

