package tickets

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	Id        int
	Nombre    string
	Email     string
	Destino   string
	HoraVuelo time.Time
	Precio    int
}

type TicketError struct {
	Err   error
	Field string
	Desc  string
}

type Tickets []Ticket
type TicketPointers []*Ticket

// Lista con todos los Tickets en el orden de lectura desde el fichero
var ticketList Tickets

// Mapa con los tickets indexados por Destino
var destinationMap map[string]TicketPointers = make(map[string]TicketPointers)

// Mapa con los tickets indexados por Periodo de tiempo
var periodMap map[string]TicketPointers = make(map[string]TicketPointers)

/******************
*
*  PRIVATE FUNCTIONS
*
*******************/

// Agrega los tickets al listado y los indexa según destino y periodo
func addTicketToList(t Ticket) Tickets {
	ticketList = append(ticketList, t)

	// Llenamos el mapa de tickets indexado por pais con un slice de punteros a los tickets
	destinationMap[t.Destino] = append(destinationMap[t.Destino], &t)
	hour := t.HoraVuelo.Hour()
	period := ""
	switch {
	case hour >= 0 && hour <= 6:
		period = "madrugada"
	case hour >= 7 && hour <= 12:
		period = "mañana"
	case hour >= 13 && hour <= 19:
		period = "tarde"
	case hour >= 20 && hour <= 23:
		period = "noche"
	}
	// fmt.Println(period)

	// Llenamos el mapa de tickets indexado por periodo con un slice de punteros a los tickets
	periodMap[period] = append(periodMap[period], &t)

	return ticketList
}

// Obtiene un slice de tickets basado en su Destino
func getTicketsByDestination(destination string) (TicketPointers, error) {
	ticketsByDestination := destinationMap[destination]
	if ticketsByDestination == nil {
		return nil, fmt.Errorf("no hay ningún ticket para el destino %s", destination)
	}
	return ticketsByDestination, nil
}

// Crea un Time con solo Hora y Minutos, a partir de un String "hh:mm"
func convertStringToTime(timeStr string) (time.Time, error) {

	timePart := strings.Split(timeStr, ":")

	if len(timePart) != 2 {
		return time.Time{}, &TicketError{
			Err:   fmt.Errorf("el valor %s no tiene el formato HH:MM", timeStr),
			Field: "Hora del vuelo",
			Desc:  "no tiene un formato de hora válido",
		}
	}
	hour, err := strconv.Atoi(timePart[0])
	if err != nil {
		return time.Time{}, &TicketError{
			Err:   err,
			Desc:  "La hora no se puede convertir a entero",
			Field: "Hora del vuelo",
		}
	}

	if hour < 0 || hour > 23 {
		return time.Time{}, &TicketError{
			Err:   fmt.Errorf("la hora %d no está en el rango 0 a 23", hour),
			Field: "Hora del vuelo",
			Desc:  "no tiene un formato de hora válido",
		}
	}

	minute, err := strconv.Atoi(timePart[1])
	if err != nil {
		return time.Time{}, &TicketError{
			Err:   err,
			Desc:  "Los minutos no se pueden convertir a entero",
			Field: "Hora del vuelo",
		}
	}

	myTime := time.Date(0, 0, 0, hour, minute, 0, 0, time.Local)
	return myTime, nil

}

// Convierte una línea en Ticket
func convertLineToTicket(line string) (*Ticket, error) {

	// partir la línea por ,
	columns := strings.Split(line, ",")
	//fmt.Println(columns)

	// determinar que tenga la longitud requerida
	if len(columns) != 6 {
		return nil, &TicketError{
			Err:   fmt.Errorf("solo tiene %d columnas de las 6 necesarias", len(columns)),
			Desc:  "está incompleta",
			Field: "La fila",
		}
	}

	Id, err := strconv.Atoi(columns[0])
	if err != nil {
		return nil, &TicketError{
			Err:   err,
			Desc:  "no se puede convertir a entero",
			Field: "Id",
		}
	}
	Nombre := columns[1]
	Email := columns[2]
	Destino := columns[3]
	HoraVuelo, err := convertStringToTime(columns[4])
	if err != nil {
		return nil, err
	}

	Precio, err := strconv.Atoi(columns[5])
	if err != nil {
		return nil, &TicketError{
			Err:   err,
			Desc:  "no se puede convertir a entero",
			Field: "Precio",
		}
	}

	myTicket := &Ticket{
		Id:        Id,
		Nombre:    Nombre,
		Email:     Email,
		Destino:   Destino,
		HoraVuelo: HoraVuelo,
		Precio:    Precio,
	}

	return myTicket, nil

}

/******************
*
*  PUBLIC FUNCTIONS
*
*******************/
// Carga el fichero de disco y registra cada ticket en una lista de tickets
func LoadFile(name string) {
	// Leer el fichero desde disco
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	// Separar cada línea del fichero
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		// recorrer cada línea y convertirla en Ticket
		ticket, err := convertLineToTicket(line)
		if err != nil {
			fmt.Printf("No se procesará la siguiente línea: ´%s´ porque %s\n", line, err)
			continue
		}
		addTicketToList(*ticket)
	}
}

// Obtiene la lista de todos los tickets
func GetTicketList() Tickets {
	return ticketList
}

// Obtiene la lista de todos los paises
func GetDestionations() []string {
	var destinations []string
	for destination := range destinationMap {
		destinations = append(destinations, destination)
	}

	return destinations
}

// Función que calcule cuántas personas viajan a un país determinado.
func GetTotalTickets(destination string) (int, error) {

	ticketsByDestination, err := getTicketsByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(ticketsByDestination), nil
}

// Función que calcula cuántas personas viajan en algún periodo de tiempo
func GetCountByPeriod(time string) (int, error) {

	if periodMap[time] == nil {
		return 0, fmt.Errorf("no existe un período de tiempo llamado %s", time)
	}

	return len(periodMap[time]), nil
}

// Función que calcula el precio promedio de los tiquetes a un destino
// se eliminó el parámetro total de la consigna porque no tenía ningún sentido
func AverageDestination(destination string) (int, error) {

	ticketsByDestination, err := getTicketsByDestination(destination)
	if err != nil {
		return 0, err
	}
	count, sum := 0, 0
	for _, ticket := range ticketsByDestination {
		sum += ticket.Precio
		count++
	}
	return sum / count, nil
}

/******************
*
*  MÉTODOS
*
*******************/

func (t Ticket) String() string {

	var str string = "Ticket {" +
		"\n\tId: %d" +
		"\n\tNombre: %s" +
		"\n\tEmail: %s" +
		"\n\tDestino: %s" +
		"\n\tHora de vuelo: %s" +
		"\n\tPrecio: %d" +
		"\n}"

	hora := fmt.Sprintf("%d:%d", t.HoraVuelo.Hour(), t.HoraVuelo.Minute())
	return fmt.Sprintf(str, t.Id, t.Nombre, t.Email, t.Destino, hora, t.Precio)
}

// Error custom
func (err *TicketError) Error() string {
	return fmt.Sprintf("%s %s con error: %s", err.Field, err.Desc, err.Err.Error())
}
