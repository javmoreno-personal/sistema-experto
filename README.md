
## Integrantes
- [x] **Richard ----**
- [x] **Adrian Rocha**
- [x] **Agustin ----**
- [x] **Javier Moreno**



# Sistema experto para la detección de enfermedades 

**Resíratorias** 
- Gripe
- Neumonía
- Bronquitis

**Infecciosas y parasitarias**
- Dengue
- Malaria
- Parasitosis intestinal

**Digestivas**
- Gastritis
- Colitis
- Ulcera

**Mentales**
- Depresión
- Retraso mental
- Anorexia

## Descripción del proyecto

Consiste en crear un sistema experto basado en conocimiento para la detección de enfermedades, el cual se basa en un conjunto de reglas que se le proporcionan al sistema para que este pueda inferir y llegar a una conclusión. Inialmente se le pedira al usuario que ingrese nombre, sexo y edad para poder realizar la inferencia de la enfermedad, posteriormente se le dara una tabla para que seleccione los sintomas que presenta, y finalmente se le mostrara el resultado de la posible enfermedad que tiene y el tratamiento que debe seguir.

## Requisitos

- [x] **Go 1.16**
- [x] **Gin**

Ejemplo JSON de entrada

```json
{
    "nombre": "Richard",
    "sexo": "Masculino",
    "edad": 21,
    "sintomas": [
        "Tos",
        "Fiebre",
        "Dolor de cabeza"
    ]
}
```

Ejemplo JSON de salida

```json
{
    "nombre": "Richard",
    "sexo": "Masculino",
    "edad": 21,
    "enfermedad": "Gripe",
    "tratamiento": "Tomar paracetamol"
}
```

## Se usara un motor de inferencia hacia llamado GRULE RULE ENGINE 

Este nos permitira crear las reglas para que el sistema experto pueda inferir y llegar a una conclusion, este motor ademas evita que limitemos el sistema a un conjunto de if-else, es decir, las conclusiones que se obtienen son mas precisas y no se limitan a un conjunto de reglas.

## Instalación

Para instalar el motor de inferencia GRULE RULE ENGINE se debe ejecutar el siguiente comando:

```bash
go get github.com/hyperjumptech/grule-rule-engine
```

Para instalar el framework Gin se debe ejecutar el siguiente comando:

```bash
go get github.com/gin-gonic/gin
```

## Ejecución

Para ejecutar el programa se debe ejecutar el siguiente comando:

```bash
go run main.go
```

## Referencias

- [x] **[GRULE RULE ENGINE]
- [x] **[GIN]

## Para el despliegue de la aplicacion se utilizara Heroku 

- [x] **[HEROKU]
**Instrucciones para el despliegue de la aplicacion en Heroku**

1. Crear una cuenta en Heroku
2. Crear una nueva aplicacion
3. Seleccionar la opcion de GitHub
4. Seleccionar el repositorio que contiene el proyecto
5. Seleccionar la opcion de Deploy Branch
6. Seleccionar la opcion de View
7. Seleccionar la opcion de Open app
