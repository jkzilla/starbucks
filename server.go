package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"starbucks/graph"
	"starbucks/graph/model"
)

const defaultPort = "8080"

// corsMiddleware adds CORS headers to allow frontend and Apollo Studio access
func corsMiddleware(next http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",              // Local development
			"http://localhost:8080",              // Production
			"https://studio.apollographql.com",   // Apollo Studio
		},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true, // Required for cookies with SameSite=None; Secure
	})
	return c.Handler(next)
}

// spaHandler implements the http.Handler interface for serving a SPA
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	// Check if file exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// Other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// File exists, serve it
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// startServer runs the GraphQL API server
func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create resolver with sample data
	resolver := graph.NewResolver()

	// Add sample coffees - Starbucks classics
	resolver.AddCoffee(&model.Coffee{
		ID:                "1",
		Name:              "Pike Place Roast",
		Size:              "Grande",
		Price:             3.95,
		QuantityAvailable: 20,
		Description:       strPtr("Smooth and balanced, our signature medium roast coffee"),
	})
	resolver.AddCoffee(&model.Coffee{
		ID:                "2",
		Name:              "Caffè Americano",
		Size:              "Grande",
		Price:             3.65,
		QuantityAvailable: 25,
		Description:       strPtr("Espresso shots topped with hot water for a rich, full-bodied coffee"),
	})
	resolver.AddCoffee(&model.Coffee{
		ID:                "3",
		Name:              "Caffè Latte",
		Size:              "Grande",
		Price:             4.95,
		QuantityAvailable: 15,
		Description:       strPtr("Rich espresso balanced with steamed milk and a light layer of foam"),
	})
	resolver.AddCoffee(&model.Coffee{
		ID:                "4",
		Name:              "Cappuccino",
		Size:              "Grande",
		Price:             4.75,
		QuantityAvailable: 18,
		Description:       strPtr("Dark, rich espresso with steamed milk and a deep layer of foam"),
	})
	resolver.AddCoffee(&model.Coffee{
		ID:                "5",
		Name:              "Caramel Macchiato",
		Size:              "Grande",
		Price:             5.45,
		QuantityAvailable: 12,
		Description:       strPtr("Freshly steamed milk with vanilla syrup, espresso, and caramel drizzle"),
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Create a custom mux to handle routing properly
	mux := http.NewServeMux()
	
	// GraphQL endpoints (register first for priority)
	mux.Handle("/graphql", corsMiddleware(srv))
	mux.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	// Serve static files from frontend/dist
	spa := spaHandler{staticPath: "frontend/dist", indexPath: "index.html"}
	mux.Handle("/", spa)

	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("GraphQL playground: http://localhost:%s/playground", port)
	log.Printf("Frontend: http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func strPtr(s string) *string {
	return &s
}
