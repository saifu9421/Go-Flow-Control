package main

import (
	"fmt"
);

func found(ar [] int, n int, x int) bool{  // bool return type;
	  var falg bool = false;
   for i:=0;i<n;i++{
	     if ar[i] == x{
			   falg = true;
			   break;
		 };
   };
    return falg; 
};
 
  func main(){
	//  ar := [5] int {1,2,3,4,5};
	// //  for i := 0; i<5;i++{
	// // 	 fmt.Println(ar[i]);
	//  };

	//  var n int;
	//   fmt.Scanln(&n);
	// //    var ar int;
	//    ar := make([]int ,n);
	//   for i:= 0; i<n;i++{
	// 	    fmt.Scan(&ar[i]);
	//   };
	//    fmt.Println("Array Print");
	//   for i:=0;i<n;i++{
	// 	 fmt.Println(ar[i]);
	//   };
	var n int;
	 fmt.Scan(&n);
	  ar := make([]int,n);
	  for i:=0;i<n;i++{
		 fmt.Scan(&ar[i]);
	  };
	   fmt.Println("Array Print");
	 for i:=0;i<n;i++{
		 fmt.Println(ar[i]);
	 };
	  
	 var fl bool;
	  fl = found(ar,n,10);
	  if fl == true{
		  fmt.Println("FOUND");
	  }else{
		 fmt.Println("NOT FOUND");
	  };
	  
  }