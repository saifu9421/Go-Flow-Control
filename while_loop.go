package main 
import (
	 "fmt";
);
  func multiplication (i int ,j int,x int){
     for i<=j {
		  fmt.Println(x,"*",i,"=", x*i);
		  i++;
	 };
  };
   
 func main(){
	 i := 1; 
	  for i<= 5{
		  fmt.Println(i);
		  i++;
	  };
	    x :=1;
	  for  true {
		if x == 51{
			break;
	   };
		   fmt.Println(x);
		   x++;
	  };

	  //Create multiplication table using while loop
	    var m,j,y int;
		m  = 1;
		j = 10;
		y = 5;
	  multiplication(m,j,y);
	   
 }