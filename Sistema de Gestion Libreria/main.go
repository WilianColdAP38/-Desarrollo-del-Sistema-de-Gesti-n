package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Fase 1: Definición de estructuras principales

// Aqui se crea un struct llamado libro el cual guarda la información básica de un libro electrónico.
type Libro struct {
	ISBN       string
	Titulo     string
	Autor      string
	Categoria  string
	Disponible bool
}

// Aqui se crea un struct llamado usuario el cual representa a un usuario dentro del sistema.
type Usuario struct {
	ID            int
	Nombre        string
	Email         string
	FechaRegistro time.Time
}

// Fase 2: Estructura de Gestión y Composición

// Aquí creo una estructura que va a manejar todo el sistema.
// Guarda los usuarios, los libros y un contador para asignar IDs nuevos.
type SistemaGestion struct {
	Usuarios   []Usuario
	Libros     []Libro
	NextUserID int // este número se usa para ir dando IDs automáticos a los usuarios nuevos
}

// Fase 3: Implementación de Funcionalidades

// Este método se encarga de crear un nuevo usuario dentro del sistema.
// Primero se validan los datos ingresados antes de registrarlo.
func (s *SistemaGestion) CrearUsuario(nombre, email string) {
	// Se eliminan posibles espacios en blanco al inicio o al final
	nombre = strings.TrimSpace(nombre)
	email = strings.TrimSpace(email)

	// Si los campos están vacíos, no se realiza el registro
	if nombre == "" || email == "" {
		fmt.Println("Debe ingresar un nombre y un correo válido.")
		return
	}

	// Se verifica que el correo contenga el símbolo '@' como validación básica
	if !strings.Contains(email, "@") {
		fmt.Println("El correo no es válido.")
		return
	}
	// Se crea la estructura Usuario, asignándole el ID correspondiente y la fecha actual.
	nuevoUsuario := Usuario{
		ID:            s.NextUserID,
		Nombre:        nombre,
		Email:         email,
		FechaRegistro: time.Now(), // Utiliza el paquete 'time' que fue importado.
	}

	// Se agrega el nuevo usuario al slice de usuarios usando append().
	s.Usuarios = append(s.Usuarios, nuevoUsuario)

	// Se incrementa el contador de IDs para el siguiente registro.
	s.NextUserID++

	fmt.Printf("\n✅ Usuario '%s' registrado con ID: %d.\n", nombre, nuevoUsuario.ID)
}

// ListarUsuarios: Muestra en consola todos los usuarios registrados en el sistema.
// Recorre el slice de usuarios utilizando un bucle for range.
func (s SistemaGestion) ListarUsuarios() {
	// Se verifica si el slice está vacío. Si no hay usuarios, se muestra un mensaje y se finaliza la función.
	if len(s.Usuarios) == 0 {
		fmt.Println("\nNo hay usuarios registrados en el sistema.")
		return
	}

	// Se imprime el encabezado del listado.
	fmt.Println("\n--- LISTADO DE USUARIOS ---")

	// Se recorre el slice de usuarios e imprime la información de cada registro.
	for _, usuario := range s.Usuarios {
		// Se usa el método Format() del paquete time para mostrar la fecha con un formato legible.
		fmt.Printf("ID: %d | Nombre: %s | Email: %s | Registro: %s\n",
			usuario.ID, usuario.Nombre, usuario.Email, usuario.FechaRegistro.Format("02-01-2006 15:04"))
	}
}

// CrearLibro: Permite registrar un nuevo libro en el sistema.
// Se validan los datos ingresados antes de agregarlo al slice de libros.
func (s *SistemaGestion) CrearLibro(isbn, titulo, autor, categoria string) {
	// Se eliminan posibles espacios en blanco en los datos ingresados.
	isbn = strings.TrimSpace(isbn)
	titulo = strings.TrimSpace(titulo)
	autor = strings.TrimSpace(autor)
	categoria = strings.TrimSpace(categoria)

	// Se verifica que los campos obligatorios no estén vacíos.
	if isbn == "" || titulo == "" || autor == "" {
		fmt.Println("El ISBN, el título y el autor son obligatorios.")
		return
	}

	// Se crea la estructura Libro con los datos proporcionados.
	nuevoLibro := Libro{
		ISBN:       isbn,
		Titulo:     titulo,
		Autor:      autor,
		Categoria:  categoria,
		Disponible: true, // Por defecto, el libro se marca como disponible.
	}

	// Se agrega el nuevo libro al slice de libros usando append().
	s.Libros = append(s.Libros, nuevoLibro)

	fmt.Printf("\nLibro '%s' agregado correctamente al sistema.\n", titulo)
}

// ListarLibros: Muestra en consola todos los libros registrados en el sistema.
func (s SistemaGestion) ListarLibros() {
	// Se verifica si el slice está vacío. Si no hay libros, se muestra un mensaje y se finaliza la función.
	if len(s.Libros) == 0 {
		fmt.Println("\nNo hay libros registrados en el sistema.")
		return
	}

	// Se imprime el encabezado del listado.
	fmt.Println("\n--- LISTADO DE LIBROS ---")

	// Se recorre el slice de libros e imprime la información de cada uno.
	for _, libro := range s.Libros {
		// Se define el estado de disponibilidad del libro.
		estado := "Disponible"
		if !libro.Disponible {
			estado = "No disponible"
		}

		// Se muestra la información formateada de cada libro.
		fmt.Printf("ISBN: %s | Título: %s | Autor: %s | Categoría: %s | Estado: %s\n",
			libro.ISBN, libro.Titulo, libro.Autor, libro.Categoria, estado)
	}
}

// BuscarLibrosPorTitulo: Permite buscar libros que coincidan parcialmente con el título ingresado.
// Utiliza el paquete 'strings' para realizar la búsqueda sin diferenciar mayúsculas o minúsculas.
func (s SistemaGestion) BuscarLibrosPorTitulo(busqueda string) {
	busqueda = strings.ToLower(strings.TrimSpace(busqueda))

	// Se verifica si el parámetro de búsqueda está vacío.
	if busqueda == "" {
		fmt.Println("\nDebe ingresar una palabra o frase para realizar la búsqueda.")
		return
	}

	// Se utiliza para controlar si se encontraron coincidencias.
	encontrado := false

	// Se recorre el slice de libros verificando si el título contiene la palabra buscada.
	for _, libro := range s.Libros {
		if strings.Contains(strings.ToLower(libro.Titulo), busqueda) {
			fmt.Printf("ISBN: %s | Título: %s | Autor: %s | Categoría: %s\n",
				libro.ISBN, libro.Titulo, libro.Autor, libro.Categoria)
			encontrado = true
		}
	}

	// Si no se encontraron coincidencias, se muestra el siguiente mensaje.
	if !encontrado {
		fmt.Println("\nNo se encontraron libros con ese título.")
	}
}

// ModificarDisponibilidadLibro: Cambia el estado de disponibilidad de un libro según su ISBN.
// Recorre el slice de libros y actualiza el campo 'Disponible' si encuentra una coincidencia.
func (s *SistemaGestion) ModificarDisponibilidadLibro(isbn string, disponible bool) {
	isbn = strings.TrimSpace(isbn)
	if isbn == "" {
		fmt.Println("\nDebe ingresar un ISBN válido.")
		return
	}

	// Se recorre el slice de libros en busca del ISBN ingresado.
	for i := range s.Libros {
		if s.Libros[i].ISBN == isbn {
			// Se actualiza el campo 'Disponible' con el nuevo valor.
			s.Libros[i].Disponible = disponible

			// Se muestra un mensaje de confirmación con el estado actualizado.
			estado := "disponible"
			if !disponible {
				estado = "no disponible"
			}
			fmt.Printf("\nEl libro con ISBN %s ahora está marcado como %s.\n", isbn, estado)
			return
		}
	}

	// Si no se encontró el libro con el ISBN indicado, se muestra un mensaje informativo.
	fmt.Println("\nNo se encontró un libro con el ISBN ingresado.")
}

// Fase 4: Interfaz de Consola y Control de Flujo

// leerEntrada: Permite leer una línea de texto desde la consola y devolverla como string.
func leerEntrada(mensaje string) string {
	fmt.Print(mensaje)
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// main: Punto de inicio del programa. Muestra el menú principal y controla el flujo general del sistema.
func main() {
	sistema := SistemaGestion{}

	for {
		fmt.Println("\n===== MENÚ PRINCIPAL =====")
		fmt.Println("1. Gestión de usuarios")
		fmt.Println("2. Gestión de libros")
		fmt.Println("3. Salir")

		opcion := leerEntrada("Seleccione una opción: ")

		switch opcion {
		case "1":
			menuUsuarios(&sistema)
		case "2":
			menuLibros(&sistema)
		case "3":
			fmt.Println("\nSaliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}
	}
}

// menuUsuarios: Controla las opciones relacionadas con los usuarios.
func menuUsuarios(s *SistemaGestion) {
	for {
		fmt.Println("\n--- MENÚ DE USUARIOS ---")
		fmt.Println("1. Registrar nuevo usuario")
		fmt.Println("2. Mostrar lista de usuarios")
		fmt.Println("3. Volver al menú principal")

		opcion := leerEntrada("Seleccione una opción: ")

		switch opcion {
		case "1":
			nombre := leerEntrada("Ingrese el nombre del usuario: ")
			email := leerEntrada("Ingrese el correo del usuario: ")
			s.CrearUsuario(nombre, email)
		case "2":
			s.ListarUsuarios()
		case "3":
			return
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}
	}
}

// menuLibros: Controla las opciones relacionadas con los libros del sistema.
func menuLibros(s *SistemaGestion) {
	for {
		fmt.Println("\n--- MENÚ DE LIBROS ---")
		fmt.Println("1. Registrar nuevo libro")
		fmt.Println("2. Mostrar catálogo de libros")
		fmt.Println("3. Buscar libro por título")
		fmt.Println("4. Modificar disponibilidad de un libro")
		fmt.Println("5. Volver al menú principal")

		opcion := leerEntrada("Seleccione una opción: ")

		switch opcion {
		case "1":
			isbn := leerEntrada("Ingrese el ISBN del libro: ")
			titulo := leerEntrada("Ingrese el título del libro: ")
			autor := leerEntrada("Ingrese el autor del libro: ")
			categoria := leerEntrada("Ingrese la categoría del libro: ")
			s.CrearLibro(isbn, titulo, autor, categoria)
		case "2":
			s.ListarLibros()
		case "3":
			busqueda := leerEntrada("Ingrese el título o palabra clave: ")
			s.BuscarLibrosPorTitulo(busqueda)
		case "4":
			isbn := leerEntrada("Ingrese el ISBN del libro: ")
			opcionDisp := leerEntrada("¿Desea marcarlo como disponible? (s/n): ")
			disponible := strings.ToLower(opcionDisp) == "s"
			s.ModificarDisponibilidadLibro(isbn, disponible)
		case "5":
			return
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}
	}
}
