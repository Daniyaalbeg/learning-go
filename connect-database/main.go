package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *pgx.Conn

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	var err error
	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// Ping db
	pingErr := db.Ping(context.Background())

	if pingErr != nil {
		log.Fatal("Ping failed", pingErr)
	}

	// var title string
	// var artist string
	// // err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	// err = db.QueryRow(context.Background(), "select title, artist from album").Scan(&title, &artist)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(title, artist)

	albums, err := albumsByArtistName("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumById(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}

func albumsByArtistName(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query(context.Background(), "SELECT * FROM album WHERE artist = $1", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist 1 %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist 2 %p: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist 3 %p: %v", name, err)
	}

	return albums, nil
}

func albumById(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow(context.Background(), "SELECT * FROM album WHERE id=$1", id)

	if err := row.Scan(&alb.ID, &alb.Artist, &alb.Title, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumById %p: %v", id, err)
		}
		return alb, fmt.Errorf("albumById %p: %v", id, err)
	}

	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	var id int64
	row := db.QueryRow(context.Background(), "INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", alb.Title, alb.Artist, alb.Price)

	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("addAlbum %v", err)
	}

	return id, nil
}
