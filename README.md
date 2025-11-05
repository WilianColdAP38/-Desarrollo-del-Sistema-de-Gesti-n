# 📘 Sistema de Gestión de Libros Electrónicos

Este proyecto fue desarrollado como parte de la **Etapa 2** del trabajo autónomo de la asignatura de **Programación**, donde se aplica la creación de un sistema de gestión usando el lenguaje **Go (Golang)**.  
El objetivo es construir un sistema sencillo que permita **registrar, listar, buscar y modificar libros y usuarios** desde la consola.

---

## 🧩 Descripción General

El sistema permite gestionar un pequeño catálogo de libros electrónicos y un listado de usuarios registrados.  
Funciona completamente desde la consola y está desarrollado utilizando los conceptos vistos durante las primeras semanas de clases.

**Funciones principales del sistema:**
- Registrar nuevos usuarios y listarlos.
- Registrar libros y mostrarlos en un catálogo.
- Buscar libros por título (coincidencias parciales).
- Modificar la disponibilidad de un libro (disponible / no disponible).
- Navegar entre menús para usuarios y libros.

---

## ⚙️ Tecnologías Utilizadas

- **Lenguaje:** Go (Golang)
- **Versión recomendada:** Go 1.21 o superior  
- **Ejecución en consola:** `go run main.go`

**Paquetes estándar utilizados:**
- `fmt` → para mostrar información en consola.  
- `bufio` y `os` → para leer entradas del usuario.  
- `strings` → para manejar y comparar cadenas de texto.  
- `time` → para registrar fechas de creación o registro.

---

## 🧱 Estructura del Código

Todo el código está implementado dentro de un único archivo `main.go` para mantener la simplicidad y facilitar la revisión del proyecto.

| Fase | Descripción |
|------|--------------|
| **Fase 1** | Definición de estructuras (`Libro`, `Usuario`) |
| **Fase 2** | Creación de la estructura principal `SistemaGestion` |
| **Fase 3** | Implementación de métodos (crear, listar, buscar, modificar) |
| **Fase 4** | Menús de consola e interacción con el usuario |
| **Fase 5** | Pruebas, revisión final y documentación |

---

## 👩‍💻 Ejecución del Programa

1. Asegurarse de tener Go instalado en el sistema.
2. Abrir una terminal en la carpeta del proyecto.
3. Ejecutar el siguiente comando:

```bash
go run main.go
```

## 🚀 Navegación del Sistema

Desde el menú principal se puede acceder a:

- **Gestión de usuarios**
- **Gestión de libros**
- **Salir del sistema**

---

## 🧠 Conceptos Aplicados

Durante el desarrollo se aplicaron los siguientes temas vistos en clase:

- **Funciones** y **métodos** con y sin parámetros.  
- **Condicionales** `if`, `else` y estructuras `switch`.  
- **Bucles** `for` y `for range`.  
- **Structs** para definir tipos personalizados (`Usuario`, `Libro`, `SistemaGestion`).  
- **Slices** para almacenar colecciones dinámicas de datos.  
- **Uso de punteros** para modificar estructuras dentro de los métodos.  
- **Entrada y salida de datos** desde consola con `bufio`.  
- **Uso de paquetes estándar:** `fmt`, `strings`, `time`, `os`.

---

## 💻 Ejemplos de Ejecución

### 🧍‍♂️ Registrar y Listar Usuarios

<img width="693" height="697" alt="image" src="https://github.com/user-attachments/assets/4213e6c4-38cf-4bb0-84dc-a24872499293" />


---

### 📚 Registrar y Listar Libros

<img width="853" height="706" alt="image" src="https://github.com/user-attachments/assets/db0470eb-8a76-46cf-b28d-93905b7290d7" />


---

### 🔍 Buscar Libros por Título

<img width="667" height="55" alt="image" src="https://github.com/user-attachments/assets/7676c706-1ace-458c-ad90-2a069dc77110" />


---

### 🛠️ Modificar Disponibilidad de un Libro

<img width="541" height="197" alt="image" src="https://github.com/user-attachments/assets/7acbf533-6937-4f81-90dc-7ce3e522ce1b" />


---

### 🚪 Salir del Sistema

<img width="308" height="260" alt="image" src="https://github.com/user-attachments/assets/2c8d6c11-ca0b-4cd4-b211-7ea63c343b0b" />



