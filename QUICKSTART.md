# Starbucks Coffee Shop - Quick Start Guide

## What's Been Created

A simplified Starbucks coffee shop application based on the egg app architecture, featuring:

- **Backend**: Go 1.25 + GraphQL (gqlgen)
- **Frontend**: React 18 + TypeScript + Vite + TailwindCSS
- **Features**: Coffee catalog, shopping cart, simple checkout (no authentication, no Signal)
- **Sample Data**: 5 Starbucks coffee products pre-loaded

## Project Structure

```
~/src/starbucks/
├── graph/                    # GraphQL schema and resolvers
│   ├── schema.graphqls      # Coffee product schema
│   ├── resolver.go          # Business logic
│   └── schema.resolvers.go  # Generated resolver code
├── frontend/                 # React frontend
│   ├── src/
│   │   ├── components/      # CoffeeCard, ShoppingCart, Header
│   │   ├── graphql/         # GraphQL queries
│   │   ├── App.tsx          # Main app
│   │   └── types.ts         # TypeScript types
│   └── package.json
├── main.go                   # Entry point
├── server.go                 # HTTP server
└── README.md                 # Full documentation
```

## Running the Application

### Option 1: Development Mode (Recommended)

**Terminal 1 - Backend:**
```bash
cd ~/src/starbucks
go run .
```

**Terminal 2 - Frontend:**
```bash
cd ~/src/starbucks/frontend
npm run dev
```

Then open:
- Frontend: http://localhost:5173
- GraphQL Playground: http://localhost:8080/playground

### Option 2: Production Build

```bash
cd ~/src/starbucks/frontend
npm run build
cd ..
go run .
```

Then open: http://localhost:8080

## Key Differences from Egg App

### Removed:
- ❌ Signal messaging integration
- ❌ Authentication system
- ❌ Payment method selection (cash/card)
- ❌ Pickup time scheduling
- ❌ Customer phone number collection

### Simplified:
- ✅ Basic coffee products (5 Starbucks classics)
- ✅ Simple shopping cart
- ✅ One-click checkout
- ✅ Real-time inventory updates
- ✅ Clean, modern UI with Starbucks green theme

## Testing the App

1. **Browse Coffee**: View 5 coffee products on the home page
2. **Add to Cart**: Click "Add to Cart" on any coffee
3. **View Cart**: Click the cart icon (top right)
4. **Adjust Quantities**: Use +/- buttons in cart
5. **Checkout**: Click "Checkout" button
6. **Verify**: Inventory updates automatically

## GraphQL API Examples

### Get All Coffees
```graphql
query {
  coffees {
    id
    name
    size
    price
    quantityAvailable
  }
}
```

### Purchase Coffee
```graphql
mutation {
  purchaseCoffee(id: "1", quantity: 2) {
    success
    message
    remainingQuantity
  }
}
```

## Next Steps

- Customize coffee products in `server.go`
- Modify UI colors/theme in components
- Add more features (sizes, customizations, etc.)
- Deploy using Docker or cloud platforms

## Troubleshooting

**Backend won't start:**
```bash
cd ~/src/starbucks
go mod tidy
go run github.com/99designs/gqlgen generate
```

**Frontend errors:**
```bash
cd ~/src/starbucks/frontend
rm -rf node_modules package-lock.json
npm install
```

**Port already in use:**
```bash
# Change backend port
PORT=3000 go run .

# Frontend will auto-proxy to backend
```

## Architecture Highlights

- **No Database**: In-memory storage (resets on restart)
- **CORS Enabled**: Frontend can call backend API
- **Type-Safe**: TypeScript + Go type checking
- **Modern Stack**: Latest versions of all dependencies
- **Responsive**: Works on mobile, tablet, desktop

Enjoy your Starbucks coffee shop! ☕
