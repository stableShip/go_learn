package main

import "fmt"
import "testing"


func buildArray(t testing.T){
	var myArray []string = []string{"1", "2", "3"};
	fmt.Println(myArray)
	if (len(myArray) != 3) {
		t.Error("myArray length should be 3")
	}
	if(myArray[0] != "1"){
		t.Error("myArray first value should be '1'")
	}
}