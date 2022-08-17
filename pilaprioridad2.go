/*Universidad Nacional Autonoma de Honsuras
Estudiante: Carlos Eduardo Ramirez Cortes, #cuenta: 20181005767
Clase: algoritmo y Estructura de Datos
Docente: Jose Enrique Avila

Cola de prioridad implementada segun la edad de un grupo de individuos,
entre mayor sea la edad de los individuos, mayor sera su prioridad en la cola*/

package main

import (
	"container/heap" //paquete pre definido de golan que se utiliza para colas de prioridad
	"fmt"
)

type Item struct {
	Nombre string
	Edad   int

	Index int // indice del elemento en la cola
}

type ColaPrioridad []*Item

func (pq ColaPrioridad) Len() int {
	return len(pq)
}

func (pq ColaPrioridad) Less(i, j int) bool {
	// Pop nos incerta al final de la cola la edad mas baja de las personas
	return pq[i].Edad > pq[j].Edad
}

//implementacion la función predefinida en la interfaz

func (pq *ColaPrioridad) Pop() interface{} {
	posterior := *pq
	n := len(posterior)
	item := posterior[n-1]
	item.Index = -1
	*pq = posterior[0 : n-1]
	return item
}

//funcion para ingresar los datos
func (pq *ColaPrioridad) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

//funcion para intercambiar los datos segun su prioridad
func (pq ColaPrioridad) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func main() {

	listItems := []*Item{
		{Nombre: "Carlos", Edad: 15},
		{Nombre: "Maria", Edad: 40},
		{Nombre: "Luis", Edad: 75},
		{Nombre: "Raul", Edad: 60},
		{Nombre: "Sohamy", Edad: 22},
		{Nombre: "Karla", Edad: 80},
	}
	colaPrioridad := make(ColaPrioridad, len(listItems))

	for i, item := range listItems {
		colaPrioridad[i] = item
		colaPrioridad[i].Index = i
	}

	heap.Init(&colaPrioridad)

	// Impresion de la cola de prioridad
	for colaPrioridad.Len() > 0 {
		item := heap.Pop(&colaPrioridad).(*Item)
		fmt.Printf("Nombre: %s Edad: %d\n", item.Nombre, item.Edad)
	}

}
