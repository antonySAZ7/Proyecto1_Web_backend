package handlers

import (
	"api-series/db"
	"api-series/modelo"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.DB.Query("SELECT id, nombre, genero, capitulos, portada, rating FROM series")
	if err != nil {
		http.Error(w, "Error DB", 500)
		return
	}
	defer rows.Close()

	var seriesList []modelo.Series

	for rows.Next() {
		var s modelo.Series
		err := rows.Scan(&s.Id, &s.Nombre, &s.Genero, &s.Capitulos, &s.Portada, &s.Rating)
		if err != nil {
			http.Error(w, "Error leyendo DB", 500)
			return
		}
		seriesList = append(seriesList, s)
	}

	json.NewEncoder(w).Encode(seriesList)
}

func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ := strconv.Atoi(idStr)

	var s modelo.Series

	err := db.DB.QueryRow(
		"SELECT id, nombre, genero, capitulos, portada, rating FROM series WHERE id=$1",
		id,
	).Scan(&s.Id, &s.Nombre, &s.Genero, &s.Capitulos, &s.Portada, &s.Rating)

	if err != nil {
		http.Error(w, "No encontrada", 404)
		return
	}

	json.NewEncoder(w).Encode(s)
}
func CreateSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var s modelo.Series

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if s.Nombre == "" {
		http.Error(w, "Nombre es requerido", http.StatusBadRequest)
		return
	}

	err = db.DB.QueryRow(
		"INSERT INTO series (nombre, genero, capitulos, portada, rating) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		s.Nombre,
		s.Genero,
		s.Capitulos,
		s.Portada,
		s.Rating,
	).Scan(&s.Id)

	if err != nil {
		http.Error(w, "Error insertando en DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(s)
}

func UpdateSeries(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ := strconv.Atoi(idStr)

	var s modelo.Series
	json.NewDecoder(r.Body).Decode(&s)

	_, err := db.DB.Exec(
		"UPDATE series SET nombre=$1, genero=$2, capitulos=$3, portada=$4, rating=$5 WHERE id=$6",
		s.Nombre, s.Genero, s.Capitulos, s.Portada, s.Rating, id,
	)

	if err != nil {
		http.Error(w, "Error actualizando", 500)
		return
	}

	s.Id = id
	json.NewEncoder(w).Encode(s)
}

func DeleteSeries(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/series/")
	id, _ := strconv.Atoi(idStr)

	_, err := db.DB.Exec("DELETE FROM series WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Error eliminando", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
