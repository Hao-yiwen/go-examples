package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WhERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q : %v", name, err)
	}

	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q : %v", name, err)
		}
		albums = append(albums, alb)
	}
	// 这里的 rows.Err() 是用来检查在遍历 rows 的过程中是否发生了延迟（deferred）错误，比如数据库连接中断、解码数据出错等情况。
	// 即使前面 db.Query 和 rows.Scan 都已经做了错误处理，但遍历所有 rows 过程中底层驱动可能还是会遇到其他异步错误，这里要把这些 deferred error 捕获出来。
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q : %v", name, err)
	}
	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d : no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d : %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	// for i := 0; i < 3; i++ {
	// 	defer func() {
	// 		fmt.Println(i) // 每次循环 i 都是新的变量
	// 	}()
	// }

	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DB_ADDR")
	cfg.DBName = os.Getenv("DB_NAME")

	var err error
	// FormatDSN 的作用是根据 cfg 里的配置把数据库连接信息格式化成 MySQL DSN 字符串，比如：
	// "user:password@tcp(127.0.0.1:3306)/dbname"
	// 这样 sql.Open 可以读取这个字符串建立连接。
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		// log.Fatal 会在输出错误日志后直接调用 os.Exit(1) 终止程序，而 fmt 只能输出内容并不会终止程序。
		// 下面用 fmt 输出错误，但不会退出程序，需手动加 os.Exit(1) 实现相同效果。
		fmt.Println(err)
		os.Exit(1)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
	newAlb := Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}

	albId, err := addAlbum(newAlb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albId)
}
