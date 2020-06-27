package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	//"time"
)


type dat struct{
	Nombre string `json:"name"`
	Departamento string `json:"depto"`
	Edad int `json:"age"`
	FormaContagio string `json:"form"`
	Estado string `json:"state"`
}

func main() {
	go f1(99999999999999)
	go f2(99999999999999)
	go f3(99999999999999)
	go f4(99999999999999)
	go f5(99999999999999)
	go f6(99999999999999)
	go f7(99999999999999)
	go f8(99999999999999)
	go f9(99999999999999)
	go f10(99999999999999)
	go f11(99999999999999)
	go f12(99999999999999)
	go f13(99999999999999)
	go f14(99999999999999)
	go f15(99999999999999)
	go f16(99999999999999)
	go f17(99999999999999)
	go f18(99999999999999)
	//time.Sleep(1000 * time.Millisecond)
	var input string
	fmt.Scanln(&input)
}




func f1(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre1","depto1",10,"form1","est1"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f2(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre2","depto2",20,"form2","est2"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f3(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre3","depto3",30,"form3","est3"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f4(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre1","depto1",4,"form1","est1"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f5(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre2","depto2",5,"form2","est2"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f6(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre3","depto3",6,"form3","est3"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f7(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre1","depto1",7,"form1","est1"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f8(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre2","depto2",8,"form2","est2"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f9(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre3","depto3",9,"form3","est3"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f10(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre1","depto1",10,"form1","est1"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f11(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre2","depto2",11,"form2","est2"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f12(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre3","depto3",12,"form3","est3"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f13(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre1","depto1",13,"form1","est1"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f14(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre2","depto2",14,"form2","est2"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f15(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre3","depto3",15,"form3","est3"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f16(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre1","depto1",16,"form1","est1"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f17(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre2","depto2",17,"form2","est2"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}

func f18(n int) {
	for i:=0; i<n; i++ {
		var jsonData = dat{"nombre3","depto3",18,"form3","est3"};
		jsonValue, _ := json.Marshal(jsonData)
		response, err := http.Post("http://contour.sopes1djgg.tk/datos", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
		//time.Sleep(100 * time.Millisecond)
	}
}