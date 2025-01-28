package main 
import "fmt";
 func main(){
	//   var  i int;
	//    var sum int  = 0;
	//    for i = 1 ;i<=10;i++{
    //         fmt.Println(i);
	// 		//  sum += sum+i;
	// 		sum = sum+i;
	//    }; 
	//    fmt.Println(sum);
	 var i int;
	  var sum int  = 0;
	   for i = 1;i<=10;i++{
		 if i%2 == 0{
			 fmt.Println(i);
			  sum = sum+i;
		 };
	   };
	    fmt.Println(sum);

        var sum2 int = 0;
	 for x := 1;  x<=10;x++{
		 if x%2  !=  0{
			  fmt.Println(x);
			   sum2 = sum2 +x;
		 };
	 };
	  fmt.Println(sum2);
	   
	  for y := 1;y<=100;y++{
		     fmt.Println(y);
		    if(y == 50){ 
				//  continue;
				break;
			};
	  };
	//   leep year ber korar program 
	 var  year int;
	  fmt.Scan(&year);
	  if(year %4 == 0){
		 fmt.Println("Leep Year");
	  }else{
		 fmt.Println("NOT Leep Year");
	  }
 }