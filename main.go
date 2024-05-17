package main

import (
	"fmt"
	"strings"
)

// 1. Árbol de búsqueda de prefijos
/*
const AlphabetSize = 26 // Numero de posibles caracteres (a-z)

// Nodo representa cada nodo en el Trie
type Nodo struct {
	hijos          [AlphabetSize]*Nodo
	esFinDePalabra bool
}

// Trie representa el Trie y tendra un nodo raiz
type Trie struct {
	raiz *Nodo
}

// Funcion que crea un nuevo Trie
func inicializaTrie() *Trie {
	return &Trie{raiz: &Nodo{}}
}

// Funcion para agregar una palabra al Trie
func (t *Trie) Insertar(palabra string) {
	nodo := t.raiz
	for _, caracter := range palabra {
		indice := caracter - 'a'
		if nodo.hijos[indice] == nil {
			nodo.hijos[indice] = &Nodo{}
		}
		nodo = nodo.hijos[indice]
	}
	nodo.esFinDePalabra = true
}

// Verifica si una palabra esta en el Trie
func (t *Trie) buscar(palabra string) bool {
	nodo := t.raiz
	for _, caracter := range palabra {
		indice := caracter - 'a'
		if nodo.hijos[indice] == nil {
			return false
		}
		nodo = nodo.hijos[indice]
	}
	return nodo != nil && nodo.esFinDePalabra
}

// Busca si un prefijo esta presente en el Trie
func (t *Trie) buscarPrefijo(prefijo string) bool {
	nodo := t.raiz
	for _, caracter := range prefijo {
		indice := caracter - 'a'
		if nodo.hijos[indice] == nil {
			return false
		}
		nodo = nodo.hijos[indice]
	}
	return true
}

func main() {
	trie := inicializaTrie()

	palabra := []string{"hola", "mundo"}
	for _, palabra := range palabra {
		trie.Insertar(palabra)
	}

	// Prueba de busqueda
	fmt.Println(trie.buscar("hola"))       // Deberia devolver true
	fmt.Println(trie.buscar("mundo"))      // Deberia devolver true
	fmt.Println(trie.buscar("hol"))        // Deberia devolver false
	fmt.Println(trie.buscarPrefijo("hol")) // Deberia delvolver true
}
*/

// 2. Estructura de datos para auot-completar

/*
const AlphabetSize = 26 // Numero de posibles caracteres (a-z)

// Nodo representa cada nodo en el Trie
type Nodo struct {
	hijos          [AlphabetSize]*Nodo
	esFinDePalabra bool
}

// Trie representa el Trie y tendra un nodo raiz
type Trie struct {
	raiz *Nodo
}

// Funcion que crea un nuevo Trie
func inicializaTrie() *Trie {
	return &Trie{raiz: &Nodo{}}
}

// Funcion para agregar una palabra al Trie
func (t *Trie) Insertar(palabra string) {
	nodo := t.raiz
	for _, caracter := range palabra {
		indice := caracter - 'a'
		if nodo.hijos[indice] == nil {
			nodo.hijos[indice] = &Nodo{}
		}
		nodo = nodo.hijos[indice]
	}
	nodo.esFinDePalabra = true
}

// Verifica si una palabra esta en el Trie
func (t *Trie) buscar(palabra string) bool {
	nodo := t.raiz
	for _, caracter := range palabra {
		indice := caracter - 'a'
		if nodo.hijos[indice] == nil {
			return false
		}
		nodo = nodo.hijos[indice]
	}
	return nodo != nil && nodo.esFinDePalabra
}

// Busca si un prefijo esta presente en el Trie
// Busca si un prefijo esta presente en el Trie y devuelve la palabra que coincide
// Busca la palabra completa con el prefijo dado en el Trie
func (t *Trie) buscarPalabraConPrefijo(prefijo string) (string, bool) {
	nodo := t.raiz
	for _, caracter := range prefijo {
		indice := caracter - 'a'
		if nodo.hijos[indice] == nil {
			return "", false
		}
		nodo = nodo.hijos[indice]
	}
	return t.encontrarPalabraDesdeNodo(nodo, prefijo)
}

// Encuentra la palabra desde el nodo dado
func (t *Trie) encontrarPalabraDesdeNodo(nodo *Nodo, prefijo string) (string, bool) {
	if nodo.esFinDePalabra {
		return prefijo, true
	}
	for i, hijo := range nodo.hijos {
		if hijo != nil {
			caracter := rune('a' + i)
			palabra, encontrado := t.encontrarPalabraDesdeNodo(hijo, prefijo+string(caracter))
			if encontrado {
				return palabra, true
			}
		}
	}
	return "", false
}

func main() {
	trie := inicializaTrie()

	palabras := []string{"avion", "banana", "carro"}
	for _, palabra := range palabras {
		trie.Insertar(palabra)
	}

	// Prueba de búsqueda de palabra completa con prefijo
	fmt.Print("Escriba una palabra para auto-completar: ")
	var prefijo string
	fmt.Scanln(&prefijo)
	palabraCompleta, encontrado := trie.buscarPalabraConPrefijo(prefijo)
	if encontrado {
		fmt.Printf("Escribio la palabra '%s' Quizas desea escribir: '%s
n", prefijo, palabraCompleta)
	} else {
		fmt.Printf("No se encontró una palabra con el prefijo '%s'\n", prefijo)
	}
}
*/

// 3. Árbol de segmentos

/*
// Estructura del Arbol Segmentado
type ArbolSegmentado struct {
	arbol []int // Arbol para almacenar los segmentos
	n     int   // Tamaño del arreglo subyacente
}

// Constructor del nuevo Arbol Segmentado
func nuevoArbolSegmentado(tamanio int) *ArbolSegmentado {
	as := &ArbolSegmentado{
		arbol: make([]int, 4*tamanio), // Inicializa el Arbol con un tamaño 4 veces el del arreglo
		n:     tamanio,
	}
	return as
}

// Funcion para crear el Arbol Segmentado a partir de un arreglo
func (as *ArbolSegmentado) crear(arr []int, i, al, ar int) {
	if al == ar {
		as.arbol[i] = arr[al] // Si es una hoja, almacena el valor en el árbol
	} else {
		tm := (al + ar) / 2                           // Calcula el punto medio
		as.crear(arr, i*2, al, tm)                    // Construye el subarbol izquierdo
		as.crear(arr, i*2+1, tm+1, ar)                // Construye el subarbol derecho
		as.arbol[i] = as.arbol[i*2] + as.arbol[i*2+1] // Suma los valores de los subarboles
	}
}

// Funcion para actualizar un valor en un índice específico
func (as *ArbolSegmentado) actualizar(i, al, ar, pos, nuevoVal int) {
	if al == ar {
		as.arbol[i] = nuevoVal // Actualiza el valor en la hoja
	} else {
		pm := (al + ar) / 2 // Calcula el punto medio
		if pos <= pm {
			as.actualizar(i*2, al, pm, pos, nuevoVal) // Actualiza el subarbol izquierdo
		} else {
			as.actualizar(i*2+1, pm+1, ar, pos, nuevoVal) // Actualiza el subarbol derecho
		}
		as.arbol[i] = as.arbol[i*2] + as.arbol[i*2+1] // Recalcula el valor del nodo padre
	}
}

// Funcion para calcular la suma de un rango de valores
func (as *ArbolSegmentado) sumar(i, al, ar, l, r int) int {
	if l > r {
		return 0 // Si el rango es invalido, retorna 0
	}
	if l == al && r == ar {
		return as.arbol[i] // Si el rango coincide con el nodo, retorna su valor
	}
	tm := (al + ar) / 2 // Calcula el punto medio
	// Calcula la suma de los rangos de los subarboles y los suma
	return as.sumar(i*2, al, tm, l, min(r, tm)) + as.sumar(i*2+1, tm+1, ar, max(l, tm+1), r)
}

// Funciones auxiliares para calcular el minimo y maximo
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}      // Arreglo de ejemplo
	as := nuevoArbolSegmentado(len(arr)) // Crea un nuevo arbol de segmentos
	as.crear(arr, 1, 0, len(arr)-1)      // Construye el arbol con el arreglo

	// Imprime la suma de valores en el rango(1, 3)
	fmt.Println("Suma de valores en el rango(1, 3):", as.sumar(1, 0, len(arr)-1, 1, 3))

	// Actualiza el valor en el indice 2 a 10
	as.actualizar(1, 0, len(arr)-1, 2, 10)

	// Imprime la suma de valores en el rango(1, 3) despues de la actualizacion
	fmt.Println("Suma de valores en el rango(1, 3) despues de la actualizacion:", as.sumar(1, 0, len(arr)-1, 1, 3))
}
*/

// 4. Estructura de datos para el juego del ahorcado

/*
// Struct del ahorcado
type JuegoAhorcado struct {
	palabra       string   // Palabra a adivinar
	letras        []string // Letras adivinadas correctamente
	intentos      int      // número de intentos restantes
	intentosMax   int      // Número maximo de intentos
	palabraOculta string   //Represantacion de la palabra oculta con guiones

}

//crear funcion para el juego del ahorcado

func NuevoJuegoAhorcado(palabra string, intentosMax int) *JuegoAhorcado {

	palabra = strings.ToUpper(palabra) // Convertir las palabras en mayusculas permitiendo que las validaciones sean insensibles a letras minusculas
	palabraOculta := strings.Repeat("_ ", len(palabra))
	return &JuegoAhorcado{

		palabra:       palabra,
		letras:        make([]string, 0),
		intentos:      intentosMax,
		intentosMax:   intentosMax,
		palabraOculta: palabraOculta,
	}
}

//funcion para adivinar las letras

func (j *JuegoAhorcado) AdivinarLetra(letra string) {

	letra = strings.ToUpper(letra) //Convertimos a mayuscula
	if !strings.Contains(j.palabra, letra) {
		j.intentos--
	} else {
		j.letras = append(j.letras, letra)
		//Actualizacion de la represantacion de la palabra oculta
		palabraOculta := ""
		for _, char := range j.palabra {
			if contains(j.letras, string(char)) { //Crear funcion contains
				palabraOculta += string(char) + " "
			} else {
				palabraOculta += "_ "
			}
		}

		j.palabraOculta = palabraOculta
	}
}

//Funcion de verificacion de victoria

func (j *JuegoAhorcado) HaGanado() bool {
	return j.palabraOculta == j.palabra
}

//Funcion de verificacion de derrota

func (j *JuegoAhorcado) HaPerdido() bool {
	return j.intentos <= 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	//Crear un nuevo juego con la palabra "GOL"

	juego := NuevoJuegoAhorcado("GOL", 5)

	//Ciclo de juego

	for {
		fmt.Println("Intentos restantes: ", juego.intentos)
		fmt.Println("Palabra: ", juego.palabraOculta)

		//solicitud de jugador para adivinar una letra

		var letra string
		fmt.Print("Adivina una letra: ")
		fmt.Scanln(&letra)

		//Adivinar la letra

		juego.AdivinarLetra(letra)

		//verificar si se ha ganado o perdido

		if juego.HaGanado() {
			fmt.Println("¡Has ganado! La palabra era: ", juego.palabra)
			break
		} else if juego.HaPerdido() {
			fmt.Println("¡Has perdido! La palabra era: ", juego.palabra)
			break
		} else if juego.HaGanado() {
			fmt.Println("¡Has ganado! La palabra era: ", juego.palabra)
			break
		}
	}
}
*/

// 5. Implementacion creativo (Arbol Comprimido)

type Nodo struct {
	hijos     map[string]*Nodo
	esPalabra bool
}

func NuevoNodo() *Nodo {
	return &Nodo{hijos: make(map[string]*Nodo)}
}

type ArbolComprimido struct {
	raiz *Nodo
}

func NuevoArbolComprimido() *ArbolComprimido {
	return &ArbolComprimido{raiz: NuevoNodo()}
}

// Funcion para agrega una nueva palabra al Arbol Comprimido.
func (t *ArbolComprimido) Insertar(palabra string) {
	actual := t.raiz
	for palabra != "" {
		encontrado := false
		for clave, nodoHijo := range actual.hijos {
			prefijoComun := prefijoComunMasLargo(palabra, clave)
			if prefijoComun != "" {
				if prefijoComun == clave {
					palabra = palabra[len(prefijoComun):]
					actual = nodoHijo
				} else {
					// Divide el nodo actual en el prefijo comun y la parte restante.
					nuevoNodo := NuevoNodo()
					nuevoNodo.hijos[clave[len(prefijoComun):]] = nodoHijo
					actual.hijos = map[string]*Nodo{prefijoComun: nuevoNodo}
					palabra = palabra[len(prefijoComun):]
				}
				encontrado = true
				break
			}
		}
		if !encontrado {
			actual.hijos[palabra] = NuevoNodo()
			actual.hijos[palabra].esPalabra = true
			break
		}
	}
}

// Funcion parar buscar una palabra en el Arbol Comprimido y devuelve verdadero si la palabra existe.
func (t *ArbolComprimido) Buscar(palabra string) bool {
	actual := t.raiz
	for palabra != "" {
		encontrado := false
		for clave, nodoHijo := range actual.hijos {
			if strings.HasPrefix(palabra, clave) {
				if len(palabra) == len(clave) && nodoHijo.esPalabra {
					return true
				}
				palabra = palabra[len(clave):]
				actual = nodoHijo
				encontrado = true
				break
			}
		}
		if !encontrado {
			return false
		}
	}
	return actual.esPalabra
}

// Funcion para encontrar el prefijo comun mas largo entre dos cadenas
func prefijoComunMasLargo(s1, s2 string) string {
	longitudMinima := min(len(s1), len(s2))
	for i := 0; i < longitudMinima; i++ {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	return s1[:longitudMinima]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	trie := NuevoArbolComprimido()
	trie.Insertar("pan")
	trie.Insertar("palo")
	trie.Insertar("pala")

	fmt.Println(trie.Buscar("pan"))   // verdadero
	fmt.Println(trie.Buscar("palo"))  // verdadero
	fmt.Println(trie.Buscar("pala"))  // verdadero
	fmt.Println(trie.Buscar("pato"))  // falso
	fmt.Println(trie.Buscar("panes")) // falso
}
