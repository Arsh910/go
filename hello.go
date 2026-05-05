package main

import "fmt"

func sayHello(comapany string) string{
    return "hello , i am macantosh"
}

func mulReturn(s1 int , s2 int)(int,int){
    return s1 , s2;
}

// named return
func coordinates(s string)(x , y int){
    x = 20
    y = 40

    return
}

// gaurd clause
func divide(s int , d int)(int, error){
    
    if d == 0{
        return 0 , fmt.Errorf("can't divide by zero");
    } 
    val := s/d
    
    return val , nil
}

func main(){
    user , product := "hello world" , "mac" 

    const company = "apple"

    fmt.Println(user , product , "by" , company);
    fmt.Printf("hello %d \n",10);

    
    msg := fmt.Sprintf("hello %v " ,10);
    fmt.Println(msg);

    if angle := 90.1; angle > 90 {
        fmt.Println("its real")
    }
    
    fmt.Println(sayHello(company));
 
    // _ ignore the return value
    x , _ := mulReturn(2,3)
    fmt.Println(x)
  
    // named return , means naming the return value along with the return types of the function 
    x1, y1 := coordinates("hello")
    fmt.Println(x1,y1)

    // gaurd clawses , Early returns
    x , err := divide(1 , 0);
    if err != nil {
        fmt.Println("error : ", err)
    } else {
        fmt.Println("val : ", x)
    }

    // structs 
    
    type Hello struct{
        Name string
        Age int
        Height float64
    }

    p := Hello{
        Name : "Mac",
        Age : 30,
        Height : 20.3,
    }
    fmt.Println(p)

    // struct in struct
    
    type Wheel  struct {
        Radius float64
        Material string
    }

    type Car struct {
        Name string
        Engine string
        FrontWheel Wheel
        BackWheel Wheel
    }
    car1 := Car{
        Name : "etios",
        Engine : "D4D",
        FrontWheel : {Radius : 20.4 , Material : "soft",},
        BackWheel : {Radius : 22.8 , Material : "hard",},
    }

    fmt.Println(car1)

    // anonymus struct
	MyCar := struct {
		Name string
		Brand string	
	}{
		Name : "Avanza",
		Brand : "Toyota",
	}	

	fmt.Println(MyCar)

	// embedded struct

	type Car struct {
		Name string
		Brand string
	}

	type Employee struct {
		Name string
		Age int
	}

	type CarEmployee struct {
		Car
		Employee
	}

	MyCarEmployee := CarEmployee{
		Car : Car{
			Name : "Avanza",
			Brand : "Toyota",
		},
		Employee : Employee{
			Name : "John Doe",
			Age : 30,
		},
	}

	// the difference between embedded and nested  
	// embedded struct we can access the fields directly without using the name of the embedded struct, 
	
	// nested struct we have to use the name of the nested struct to access the field of the nested struct.

	fmt.Println(MyCarEmployee)	


	// interface 

}



