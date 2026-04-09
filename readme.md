# Series.exe

## Descripción
app para ver y gestionar series vistas

## Tecnologías
- Backend: Go (net/http)
- Base de datos: PostgreSQL
- Frontend: HTML, CSS, JavaScript 
- se importo la libreria de postgres para go

## Cómo correr el backend

1. Instalar PostgreSQL
2. Crear base de datos:

```sql
CREATE DATABASE series_db;
CREATE TABLE series (
    id SERIAL PRIMARY KEY,
    nombre TEXT NOT NULL,
    genero TEXT,
    capitulos INT,
    portada TEXT
);
para ejecutar: go run main.go


los challenges, las screenshot y la reflexion estan en el readne dek fronted 