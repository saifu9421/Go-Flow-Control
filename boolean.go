package main 
import "fmt";
func main(){
	    var boolTrue bool = true
		 var boolFalse bool = false;
		  fmt.Println("This is boolean value",boolTrue,"or",boolFalse);
		//   number := 12;
		//    number2 := 20;
		//      result := number < number2;
		// 	 result2 := number > number2;
		// 	  fmt.Println(result);
		// 	   fmt.Println(result2);

		var number1  int ;
		  var number2 int;
		 fmt.Scan(&number1,&number2);
		  var  result bool;
		  result = (number1 == number2);
		  fmt.Println(result);
		  result = number1 < number2;
		  fmt.Println(result);
		   result = number1 != number2;
		   fmt.Println(result);
		   
		   var x,y int;
		      fmt.Scan(&x,&y);
			  var re bool;
			   re = (x == y) && (x < y);
			   fmt.Println(re);
			    re = (x < y) || (x == y);
				fmt.Println(re);

}