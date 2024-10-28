# Hotel Management System
Hotel Management System (HMS) is an application designed to streamline daily hotel operations, helping hoteliers manage front desk tasks efficiently. The system provides tools for room reservations, managing guest and room data, tracking booking statuses, and more, making it a one-stop solution for handling a hotel's front and back-office functions.

## Features
1. Room Reservations:
    - Create, update, and manage room bookings.
    - Check room availability based on specified dates.

2. Guest Management:
    - Store and manage guest details (personal information, booking history, etc.).
    - Quickly retrieve guest data for check-in and check-out.

3. Room Data Management:
    - Track room types, amenities, rates, and availability.
    - Configure different room categories and pricing models.

4. Payment Processing:
    - Track booking payments and payment statuses.
    - Generate invoices and link payments with specific reservations.

## Installation
### Prerequisites
1. Go (latest stable version)
2. SQLite for lightweight database storage

### Steps
1. Clone the repository:
```bash
git clone https://github.com/zakiyalmaya/hotel-management.git
cd hotel-management
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up SQLite database:
- A database file (e.g., hotel_management.db) will be automatically created if it doesn't exist.
- No additional configuration is necessary for SQLite.

4. Run the application:
```bash
go run main.go
```

## Usage
Once the application is running, access the API endpoints to perform different hotel management operations. Use tools like Postman or cURL to test various features, such as booking a room or retrieving guest data.

## Configuration
Configuration options, such as the server port, can be found in config.json. Customize these as needed.
Example config.json:
```sh
{
  "server": {
    "port": ":8080"
  },
  "database": {
    "file": "./hotel_management.db"
  }
}
```

## API Endpoints
1. **Room Reservation**
    - POST /api/rooms/reserve - Create a new reservation.
    - GET /api/rooms/available - Check room availability.
2. **Guest Management**
    - POST /api/guests/create - Register a new guest.
    - GET /api/guests/:id - Retrieve guest data.
3. **Room Management**
    - GET /api/rooms - Retrieve a list of rooms.
    - PUT /api/rooms/:id/update - Update room information.
4. **Payment Processing**
    - POST /api/payments/create - Record a payment.
    - GET /api/payments/:booking_id - Check payment status for a booking.

## Technologies
1. **Golang** - Core language for building the application.
2. **Fiber** - High-performance web framework.
3. **SQLite** - Lightweight, serverless database for simple, efficient data storage.