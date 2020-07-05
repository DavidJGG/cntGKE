package main
import(
	"bufio"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
	"math/rand"
	"bytes"
	"log"
	"net/http"
)

//----------- STRUCT PARA LEER LOS DATOS DEL JSON -----------
type Case struct {
	Name 	string 	`json:"Nombre"`
	Depto 	string 	`json:"Departamento"`
	Age		int		`json:"Edad"`
	Form	string	`json:"Forma de contagio"`
	State	string	`json:"Estado"`
} 

//----------- STRUCT PARA ENVIAR LOS DATOS A LOS LOAD BALANCER -----------
type Caso struct {
	Name 	string 	`json:"name"`
	Depto 	string 	`json:"depto"`
	Age		int		`json:"age"`
	Form	string	`json:"form"`
	State	string	`json:"state"`
}

func main(){
	sum := 1
	reader := bufio.NewReader(os.Stdin)
	for sum < 1000 {
		sum += sum
		fmt.Print("Send Data (y/n): ")
		siEnvio, err1 := reader.ReadString('\n')
		siEnvio = strings.Replace(siEnvio, "\n","", -1)
		if err1 != nil {
			fmt.Println("Entries error...!!!")
			break
		} else if siEnvio == "y" {
			sendCases()
		} else {
			fmt.Println("")
			break
		}
	}
}

func sendCases(){
	//	/home/huriel/Documents/Proyecto2_Client/P2Go/in/entrada.json
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the file path:")
	rutaFile, err1 := reader.ReadString('\n')
	fmt.Println("Enter the number of threads:")
	nHilos, err2 := reader.ReadString('\n')
	fmt.Println("Enter the number of cases:")
	nCasos, err3 := reader.ReadString('\n')
	fmt.Println("enter the url of the load balancer:")
	urlLB, err4 := reader.ReadString('\n')
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("Entries error...!!!")
	} else {
		readEntry(rutaFile,	nHilos,	nCasos,	urlLB)
	}
}


func readEntry(rutaFile string,	stringHilos string, stringCasos string, urlLB string){
	
	//---------- LIMPIANDO CADENAS DE LOS SALTOS DE LINEA ----------
	rutaFile = strings.Replace(rutaFile, "\n","", -1) 
	stringHilos = strings.Replace(stringHilos, "\n","", -1)
	stringCasos = strings.Replace(stringCasos, "\n","", -1)
	urlLB = strings.Replace(urlLB, "\n","", -1)

	//---------- VERIFICANDO QUE EL NUMERO DE HILOS SEA UN ENTERO ----------
	intHilos, err := strconv.Atoi(stringHilos)
	if err != nil {
		fmt.Println(stringHilos, "Threads is not an integer.")
	}
	//fmt.Println(intHilos)

	//---------- VERIFICANDO QUE EL NUMERO DE CASOS SEA UN ENTERO ----------
	intCasos, err := strconv.Atoi(stringCasos)
	if err != nil {
		fmt.Println(stringCasos, "Cases is not an integer.")
	}
	//fmt.Println(intCasos)

	//--------- OBTENIENDO ARREGLO DE CASOS DENTRO DEL JSON ---------
	cases := getCases(rutaFile)

	//-------------- ENVIANDO DATA A LOS LOAD BALANCER --------------
	sendData(cases, intHilos, intCasos, urlLB)

}

func getCases(ruta string) []Case {
    cases := make([]Case, 3)
    raw, err := ioutil.ReadFile(ruta)
    if err != nil {
		fmt.Println("Error en la lectura del archivo...!!!")
        fmt.Println(err.Error())
        os.Exit(1)
    }
    json.Unmarshal(raw, &cases)
    return cases
}

// Metodo con las go rutinas
func sendData(data []Case, nHilos int, nSolicitudes int, url string){
	fmt.Println("Send Data")
	nDatos := len(data)
	count := 0	//Contador de datos enviados

	for s := 0; s < nSolicitudes; s=s+nHilos {
		for i := 0; i < nHilos; i++ {
			posDato := rand.Intn(nDatos-1)
			count++
			nuevo := Caso{data[posDato].Name, data[posDato].Depto, data[posDato].Age, data[posDato].Form, data[posDato].State}
			go send(nuevo, url)
			time.Sleep(10 * time.Millisecond)
		}
	}
	//time.Sleep(100 * time.Millisecond)
	//time.Sleep(5+time.Second)
	//fmt.Println(count)	
}


// Post a la api
func send(caseCovid Caso, url string){
	/*
		Entrada de prueba:
		path:	entrada.json
		#hilos: 50
		#Casos: 20000
		url:	http://contour.sopes1djgg.tk/datos
	*/
	jsonReq, err := json.Marshal(caseCovid)
	if err != nil {
		log.Fatalln(err)
	}
	//time.Sleep(1000)
	//fmt.Println(caseCovid.Name)
	todoBien := false
	for todoBien==false{
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
		if err != nil {
			fmt.Print("Error in Post: ")
			fmt.Print(caseCovid.Name+", ")
			fmt.Println(err)
			todoBien=false
		}else {
			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			bodyString := string (bodyBytes)
			fmt.Println("Body Response: ")
			fmt.Println(bodyString)
			todoBien=true;
		}
	}	
}