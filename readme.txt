# Crear contenido del README en formato Markdown
readme_content = """
#  Sistema Web de Gestión de Libros Electrónicos

##  Proyecto Final - Programación Orientada a Objetos

---

##  Objetivo del Proyecto

Desarrollar una aplicación web basada en arquitectura REST utilizando el lenguaje Go, que permita gestionar usuarios, libros electrónicos y reservas, aplicando conceptos de Programación Orientada a Objetos, concurrencia y serialización JSON.

---

##  Tecnologías Utilizadas

- Lenguaje: **Go**
- Arquitectura: API REST
- Serialización: JSON (encoding/json)
- Concurrencia: sync.Mutex
- Servidor Web: net/http
- Frontend: HTML + CSS + JavaScript (Fetch API)
- Control de versiones: GitHub

---

##  Funcionalidades Principales

###  Gestión de Usuarios
- Crear usuario
- Listar usuarios

###  Gestión de Libros
- Crear libro
- Listar libros
- Control de stock

###  Gestión de Reservas
- Crear reserva validando disponibilidad
- Listar reservas

###  Reportes
- Total de reservas realizadas
- Listado de libros disponibles

---

##  Servicios Web Implementados (8 Endpoints)

| Método | Endpoint |
|--------|----------|
| POST   | /usuarios |
| GET    | /usuarios/listar |
| POST   | /libros |
| GET    | /libros/listar |
| POST   | /reservas |
| GET    | /reservas/listar |
| GET    | /reportes/total |
| GET    | /reportes/libros |

---

##  Serialización JSON

La comunicación cliente-servidor se realiza mediante serialización y deserialización JSON utilizando el paquete:

encoding/json

Ejemplo:
- json.NewEncoder(w).Encode()
- json.NewDecoder(r.Body).Decode()

---

##  Concurrencia

Se implementa control de concurrencia mediante:

sync.Mutex

Esto evita condiciones de carrera cuando múltiples usuarios realizan reservas simultáneamente, protegiendo el acceso a los datos compartidos.

---

##  Arquitectura del Proyecto

biblioteca-go/
│
├── data/
├── handlers/
├── models/
├── web/
├── go.mod
└── main.go

---

##  Visualización del Futuro

Este sistema puede evolucionar hacia:

- Arquitectura de microservicios
- Implementación en la nube (AWS, Azure, GCP)
- Base de datos relacional o NoSQL
- Autenticación con JWT
- Contenedores Docker
- Orquestación con Kubernetes
- Inteligencia Artificial para recomendaciones

---

##  Integración de las 4 Unidades

✔ Unidad 1: Arquitectura modular  
✔ Unidad 2: Programación Orientada a Objetos (Structs, métodos, encapsulación)  
✔ Unidad 3: Concurrencia (Mutex)  
✔ Unidad 4: Servicios Web REST + JSON  

---

##  Demostración

La aplicación permite:

1. Crear usuarios desde la interfaz web.
2. Crear libros con control de stock.
3. Realizar reservas validando disponibilidad.
4. Visualizar reportes en tiempo real.

---

##  Conclusión

Este proyecto permitió integrar conocimientos de programación orientada a objetos, desarrollo de APIs REST, manejo de concurrencia y arquitectura modular en Go. Se fortalecieron habilidades en diseño de software, validación de datos y comunicación cliente-servidor mediante JSON.

---

##  Integrantes

Francisco Piedra Morales  

Fecha: 2026
"""

# Guardar archivo README.md usando pypandoc
import pypandoc

output_path = "/mnt/data/README_Sistema_Gestion_Libros.md"
pypandoc.convert_text(readme_content, 'md', format='md', outputfile=output_path, extra_args=['--standalone'])

output_path
