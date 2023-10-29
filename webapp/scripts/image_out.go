package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	ID      int    `db:"id"`
	Mime    string `db:"mime"`
	Imgdata []byte `db:"imgdata"`
}

func isValidMime(mime string) bool {
	return mime == "image/jpeg" || mime == "image/png" || mime == "image/gif"
}

func getExtension(mime string) string {
	if mime == "image/jpeg" {
		return "jpg"
	} else if mime == "image/png" {
		return "png"
	} else if mime == "image/gif" {
		return "gif"
	}
	return ""
}

func saveAllImages() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true",
		"root",
		"root",
		"localhost",
		"3306",
		"isuconp",
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	offset := 0
	limit := 100

	for {
		posts := []Post{}
		err := db.Select(&posts, "select `id`, `mime`, `imgdata` from `posts` limit ? offset ?", limit, offset)
		if err != nil {
			log.Printf("DB error: %v", err)
			return
		}

		if len(posts) == 0 {
			break
		}

		for _, post := range posts {
			if isValidMime(post.Mime) {
				filename := fmt.Sprintf("../image/%d.%s", post.ID, getExtension(post.Mime))
				err := os.WriteFile(filename, post.Imgdata, 0644)
				// err := ioutil.WriteFile(filename, post.Imgdata, 0644)
				if err != nil {
					log.Printf("Failed to write file: %v", err)
					continue
				}
				log.Printf("Saved %s", filename)
			}
		}

		offset += limit
	}
}

func main() {
	saveAllImages()
}
