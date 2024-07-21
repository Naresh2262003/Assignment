# Parking Lot Management Service

A RESTful service for managing parking lot operations using Go, Echo, and PostgreSQL.

## Project Structure

- **`db/`**: Database-related functionality
  - `db.go`: Database connection and setup.
  - `init_data.go`: Initializes default data in the database.
  
- **`handler/`**: HTTP request handlers
  - `parking_lot_handler.go`: Handles parking and un-parking requests.

- **`models/`**: Data models
  - `parking_lot.go`: Model for the parking lot.
  - `parking_ticket.go`: Model for parking tickets.
  - `spot.go`: Model for parking spots.
  - `tariff.go`: Model for parking tariffs.
  - `vehicle_type.go`: Model for vehicle types.

- **`docker-compose.yml`**: Sets up PostgreSQL and the application.

- **`Dockerfile`**: Build configuration for the application.

- **`main.go`**: Application entry point.

- **`go.mod`** & **`go.sum`**: Go module files.

## Getting Started

1. **Clone the repository:**
   ```sh
   git clone <repository-url>
   cd <repository-directory>

2. **Start services:**
   ```sh
   docker-compose up --build

## Database Initialization
**The database is initialized with default data on startup.**

## API Endpoints
**GET /park/available:** List available spots.
**POST /park:** Park a vehicle.
**POST /unpark:** Unpark a vehicle and calculate the fee.

## License
**MIT License - see LICENSE for details.**