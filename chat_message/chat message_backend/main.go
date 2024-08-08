package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
	// var err error
	// db, err = sql.Open("postgres", "user=youruser dbname=yourdb password=yourpassword sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	searchPath := os.Getenv("DB_SEARCH_PATH")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", host, port, user, password, dbname, searchPath)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}

}

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	r := gin.Default()

	server := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "newMessage", func(s socketio.Conn, msg string) {
		var message Message
		err := db.QueryRow("INSERT INTO messages(text) VALUES($1) RETURNING id, text", msg).Scan(&message.ID, &message.Text)
		if err != nil {
			log.Println(err)
			return
		}
		server.BroadcastToNamespace("/", "newMessage", message)
	})

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"} // Replace with your frontend's origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization", "Accept"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))

	r.GET("/socket.io/", gin.WrapH(server))
	r.POST("/socket.io/", gin.WrapH(server))

	r.GET("/messages", getMessages)
	r.POST("/messages", createMessage)

	r.Run(":8080")
}

func getMessages(c *gin.Context) {
	rows, err := db.Query("SELECT id, text FROM messages")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	messages := []Message{}
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Text); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		messages = append(messages, m)
	}
	c.JSON(http.StatusOK, messages)
}

func createMessage(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.QueryRow("INSERT INTO messages(text) VALUES($1) RETURNING id", msg.Text).Scan(&msg.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, msg)
}
