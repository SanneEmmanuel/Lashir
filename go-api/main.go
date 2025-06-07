package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow all WebSocket origins
}

func main() {
	r := gin.Default()

	// ===== CORS MIDDLEWARE =====
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No content for preflight
			return
		}
		c.Next()
	})

	// ===== WEBSOCKET (Real-time) =====
	r.GET("/stream", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatal("WebSocket upgrade failed:", err)
		}
		defer ws.Close()

		for {
			_, audioChunk, err := ws.ReadMessage()
			if err != nil {
				break
			}

			// Call Python pitch detection
			cmd := exec.Command("python3", "../python-audio/detect.py", string(audioChunk))
			pitches, err := cmd.Output()
			if err != nil {
				ws.WriteJSON(gin.H{"error": "Pitch detection failed"})
				continue
			}

			ws.WriteJSON(gin.H{
				"pitches": string(pitches),
				"solfa":   convertToSolfa(string(pitches)),
			})
		}
	})

	// ===== FILE UPLOAD =====
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("audio")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No audio file uploaded"})
			return
		}

		// Save the file temporarily
		filePath := "/tmp/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "File save failed"})
			return
		}

		// Detect pitches
		cmd := exec.Command("python3", "../python-audio/detect.py", filePath)
		pitches, err := cmd.Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Pitch detection failed"})
			return
		}

		// Generate sheet music
		cmd = exec.Command("python3", "../python-music21/notate.py", string(pitches))
		pdfPath, err := cmd.Output()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Notation generation failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"pdf_url": string(pdfPath),
		})
	})

	r.Run(":8080")
}

// ===== HELPER FUNCTIONS =====
func convertToSolfa(note string) string {
	solfa := map[string]string{
		"C": "Do", "D": "Re", "E": "Mi",
		"F": "Fa", "G": "Sol", "A": "La", "B": "Ti",
	}
	if len(note) > 0 {
		return solfa[note[:1]] // Extract base note (e.g., "C" from "C4")
	}
	return ""
}
