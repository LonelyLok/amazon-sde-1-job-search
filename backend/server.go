package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

type RequestBody struct {
	URL string `json:"url"`
}

type ResponseData struct {
	Body string `json:"body"`
}

func loadEnv(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue // Skip comments and empty lines
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid syntax in line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	return nil
}

func handleProxyGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request: Only Post method is allowed", http.StatusBadRequest)
		return
	}
	var reqBody RequestBody
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	// Decode the JSON body
	err := decoder.Decode(&reqBody)
	if err != nil {
		http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
		return
	}
	url := reqBody.URL

	// fmt.Println("Fetching data from URL:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Expires", "0")
	req.Header.Set("User-Agent", "PostmanRuntime/7.39.0")
	// some how cookies is needed to get consisting data, otherwise the 1 year less experience filter will not work
	cookies := fmt.Sprintf("analytics_id=%s; preferred_locale=en-US", os.Getenv("ANALYTICS_ID"))

	// Parse and add each cookie
	for _, cookieStr := range strings.Split(cookies, "; ") {
		parts := strings.SplitN(cookieStr, "=", 2)
		if len(parts) != 2 {
			continue
		}
		req.AddCookie(&http.Cookie{
			Name:  parts[0],
			Value: parts[1],
		})
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	responseData := ResponseData{Body: string(body)}
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		fmt.Println("Error marshaling response:", err)
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func init() {
	err := loadEnv(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/proxy-get", handleProxyGetRequest)
	fmt.Println("Starting server at port 8080")
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Allow all origins
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	}).Handler(mux)
	http.ListenAndServe(":8080", handler)
}
