package main 
import "fmt";

 func main(){
	// fmt.Println("Hello"); 
	var val, val2 int;
	    fmt.Scan(&val,&val2);
		 if val > val2 {
			 fmt.Println("value", val2, "thele bro", val2, "<",val);
		 }else if val == val2{
			 fmt.Println("value" ,val2,"er sman", val2 ,"==",val);
		 }else{
			 fmt.Println("value",val2, "theke  coto", val2 ,">",val);
		 };
 };