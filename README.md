# Starbucks Coffee Shop

A simplified full-stack application for selling coffee, featuring a Go GraphQL backend and a modern React TypeScript frontend with a beautiful, responsive UI.

## Features

- ☕ **Coffee Catalog**: Browse different types of coffee with prices and availability
- 🛒 **Shopping Cart**: Add items to cart, adjust quantities, and checkout
- 💳 **Real-time Inventory**: Automatic inventory updates after purchases
- 🎨 **Modern UI**: Beautiful, responsive design with TailwindCSS
- 🚀 **GraphQL API**: Efficient data fetching with Apollo Client

## Tech Stack

### Backend
- **Go 1.25+**: High-performance backend
- **gqlgen**: GraphQL server implementation
- **CORS Support**: Cross-origin resource sharing enabled

### Frontend
- **React 18**: Modern UI framework
- **TypeScript**: Type-safe development
- **Vite**: Fast build tool and dev server
- **Apollo Client**: GraphQL client
- **TailwindCSS**: Utility-first CSS framework
- **Lucide React**: Beautiful icon library

## Getting Started

### Prerequisites
- Go 1.22 or higher
- Node.js 20 or higher
- npm or yarn

### Development Setup

1. **Navigate to the project:**
```bash
cd ~/src/starbucks
```

2. **Install backend dependencies:**
```bash
go mod download
```

3. **Generate GraphQL code:**
```bash
go run github.com/99designs/gqlgen generate
```

4. **Install frontend dependencies:**
```bash
cd frontend
npm install
cd ..
```

5. **Run in development mode:**

In one terminal, start the backend:
```bash
go run .
```

In another terminal, start the frontend dev server:
```bash
cd frontend
npm run dev
```

The backend will run on `http://localhost:8080` and the frontend on `http://localhost:5173`

## API Usage

### GraphQL Playground

Visit `http://localhost:8080/playground` in your browser to access the GraphQL Playground, where you can interactively explore the API and test queries.

### Queries

**Get all coffees:**
```graphql
query {
  coffees {
    id
    name
    size
    price
    quantityAvailable
    description
  }
}
```

**Get a specific coffee by ID:**
```graphql
query {
  coffee(id: "1") {
    id
    name
    size
    price
    quantityAvailable
    description
  }
}
```

### Mutations

**Purchase coffee:**
```graphql
mutation {
  purchaseCoffee(id: "1", quantity: 2) {
    success
    message
    remainingQuantity
  }
}
```

## Project Structure

```
starbucks/
├── frontend/              # React TypeScript frontend
│   ├── src/
│   │   ├── components/   # React components
│   │   ├── graphql/      # GraphQL queries and mutations
│   │   ├── App.tsx       # Main app component
│   │   ├── main.tsx      # Entry point
│   │   └── types.ts      # TypeScript types
│   ├── public/           # Static assets
│   └── package.json      # Frontend dependencies
├── graph/                # GraphQL schema and resolvers
│   ├── schema.graphqls   # GraphQL schema definition
│   ├── resolver.go       # Resolver implementation
│   └── model/            # Generated models
├── main.go               # Application entry point
├── server.go             # HTTP server and routing
└── gqlgen.yml            # GraphQL code generation config
```

## Building for Production

To build the frontend and backend separately:

```bash
# Build frontend
cd frontend
npm run build

# Build backend
go build -o starbucks .

# Run
./starbucks
```

The application will serve the frontend from `frontend/dist` and be available at `http://localhost:8080`

## Sample Data

The application comes pre-loaded with sample Starbucks coffee products:
- Pike Place Roast (Grande) - $3.95
- Caffè Americano (Grande) - $3.65
- Caffè Latte (Grande) - $4.95
- Cappuccino (Grande) - $4.75
- Caramel Macchiato (Grande) - $5.45

## License

This project is licensed under the MIT License.
