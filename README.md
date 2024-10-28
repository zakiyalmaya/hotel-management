# Hotel Management System
Hotel Management System (HMS) is an application designed to streamline daily hotel operations, helping hoteliers manage front desk tasks efficiently. The system provides tools for room reservations, managing guest and room data, tracking booking statuses, and more, making it a one-stop solution for handling a hotel's front and back-office functions.

## Features
1. Booking Reservations:
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

5. User Management
    - Manage hotelier profiles, roles, and permission.

## Installation
### Prerequisites
1. Go (latest stable version)
2. SQLite for lightweight database storage
3. Redis for caching and session management

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
1. **Booking Reservation**
    - `POST /api/booking`: Create a new reservation.
        ```sh
        curl --location 'http://localhost:3000/api/room' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Im1pY2tleW1vdXNlIiwiZXhwIjoxNzE4ODUyOTUwfQ.Db2SXux_uGo_AMRF2keS-bQFWAHuJ48p1XIz52G_2ZI' \
        --data '{
            "name": "203",
            "floor": 2,
            "type": "Queen Room",
            "status": 1,
            "price": 500000,
            "description": "A room with a queen-sized bed. May be occupied by one or two people. It also has a small table, air conditioning, television and toilet."
        }'
        ```

        - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | name | string | Y | name of the room | 
        | floor | string | Y | floor where the room is located |
        | type | string | Y | type of the room |
        | status | string | Y | status of the room (e.g Available, Booked, Maintenance, Unavailable) |
        | price | string | Y | price of the room per night |
        | description | string | Y | brief description of the room |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        
    - `GET /api/booking`: Get detail room reservation using register number.
    - `PUT /api/reschedule`: Reschedule a room reservation
2. **Guest Management**
    - `POST /api/guest`: Register a new guest.
    - `GET /api/guest`: Retrieve guest data using id.
3. **Room Management**
    - `GET /api/rooms`: Retrieve a list of rooms.
    - `POST /api/room`: Register a new room.
    - `GET /api/room`: Get detail room.
    - `PUT /api/room/:name`: Update room information
4. **Payment Processing**
    - `PUT /api/payment`: Update payment status for a booking
5. **User Management**
    - `POST /api/user`: Create a new hotelier profile.
    - `PUT /api/user:id`: Update an existing hotelier's profile.
    - `GET /api/user`: Retrieve hotelier information by id.
6. **User Authentication**
    - `POST /auth/login`: Authenticates a user and generates an access and refresh token.
    - `POST /auth/logout`: Logs the user out by invalidating the access token.
    - `POST /auth/refresh`: Issues a new access token when the refresh token is valid and unexpired.

## Technologies
1. **Golang** - Core language for building the application.
2. **Fiber** - High-performance web framework.
3. **SQLite** - Lightweight, serverless database for simple, efficient data storage.
4. **Redis** - In-memory data store, ideal for caching, session management, and handling transient data efficiently.
5. **JWT Token** - JSON Web Tokens for secure, stateless user authentication and authorization across endpoints.