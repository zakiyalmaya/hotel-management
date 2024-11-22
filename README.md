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
        curl --location 'http://localhost:3000/api/booking' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIyNTg5Nzl9.p0-u5qD6gt-sYPsVGEroZgFAhxcTRexuI92k3ByZ8IA' \
        --data '{
            "guest_id": 2,
            "room_name": "101",
            "check_in": "11-12-2024",
            "check_out": "12-12-2024",
            "payment_method": 1
        }'
        ```

        - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | guest_id | integer | Y | unique identifier for the guest | 
        | room_name | string | Y | name of the room |
        | check_in | string | Y | check in date (format: DD-MM-YYYY) |
        | check_out | string | Y | check out date (Format: DD-MM-YYYY) |
        | payment_method | integer | Y | payment method used (e.g 1=Credit Card, 2=Bank Transfer, 3=Cash) |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |
        | register_number | string | N | unique identifier for the booking transaction |
        
    - `GET /api/booking`: Get detail room reservation using register number.
        ```sh
        curl --location 'http://localhost:3000/api/booking?register_number=46242953-af4a-460c-8cf4-7d93758c3bfa' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIyNTg5Nzl9.p0-u5qD6gt-sYPsVGEroZgFAhxcTRexuI92k3ByZ8IA'
        ```

        - Request Query Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | register_number | string | Y | unique identifier for the booking transaction |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |
        | register_number | string | N | unique identifier for the booking transaction |
        | guest_id | integer | N | unique identifier for the guest |
        | guest_name | string | N | name of the guest |
        | guest_identity | string | N | identification number of the guest (e.g., passport, national ID) |
        | room_name | string | N | name of the room |
        | room_floor | string | N | floor the room is located |
        | room_type | string | N | type of the room |
        | room_status | string | N | status of the room (e.g., Available, Booked, Maintenance, Unavailable) |
        | check_in | string | N | check in date |
        | check_out | string | N | check out date |
        | paid_amount | float | N | the total amount paid by the guest |
        | payment_method | string | N | payment method used (e.g., Credit Card, Cash, Bank Transfer) |
        | payment_status | string | N | status of the payment (e.g., Panding, Completed, Failed, Canceled, Refunded) |
        | additional_request | string | N | any special requests made by the guest (e.g., Extra bed, Early check-in) |
        | created_at | string | N | timestamp when the record was created |

    - `PUT /api/reschedule`: Reschedule a room reservation
        ```sh
        curl --location --request PUT 'http://localhost:3000/api/reschedule' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIyNjIzOTl9.a-WBLlBM4LYirW-tS_Eze-UijaOsaJKGRqFOGOQo3cM' \
        --data '{
            "check_in": "01-12-2024",
            "check_out": "03-12-2024",
            "register_number": "5586a507-5e6b-4857-b9ea-914c754bc70f"
        }'
        ```

        - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | check_in | string | Y | check in date (format: DD-MM-YYYY) |
        | check_out | string | Y | check out date (Format: DD-MM-YYYY) |
        | register_number | string | Y | unique identifier for the booking transaction |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |

    - `PUT /api/payment`: Update payment status.
        ```sh
        curl --location --request PUT 'http://localhost:3000/api/payment' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIxNTM4NzZ9.m1zLH1xfQtiGg53jhTbGcQveEolu5QCdrovhj5WtPzw' \
        --data '{
            "payment_status": 3,
            "register_number": "5586a507-5e6b-4857-b9ea-914c754bc70f"
        }'
        ```

        - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | payment_status | string | Y | status of the payment |
        | register_number | string | Y | unique identifier for the booking transaction |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |

2. **Guest Management**
    - `POST /api/guest`: Register a new guest.
    - `GET /api/guest`: Retrieve guest data using id.
3. **Room Management**
    - `GET /api/rooms`: Retrieve a list of rooms.
        ```sh
        curl --location 'http://localhost:3000/api/rooms?floor=2&status=1' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIyNTAyNDJ9.O-Lf3696NMTrVJxt_UB2kX9ocE_2hGL6FxV8WCNYHXY'
        ```

        - Request Query Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | floor | string | N | floor of the room |
        | status | integer | N | status of the room (e.g., 1=Available, 2=Booked, 3=Maintenance, 4=Unavailable) |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | arrary | Y | response data |
        | id | integer | N | room id |
        | name | string | N | name of the room |
        | floor | integer | N | floor the room is located |
        | type | string | N | type of the room |
        | price | float | N | price of the room per night |
        | status | string | N | status of the room (e.g., Available, Booked, Maintenance, Unavailable) |
        | description | string | N | brief description of the room |
        
    - `POST /api/room`: Register a new room.
        ```sh
        curl --location 'http://localhost:3000/api/room' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIyNTg5Nzl9.p0-u5qD6gt-sYPsVGEroZgFAhxcTRexuI92k3ByZ8IA' \
        --data '{
            "name": "205",
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
        | status | string | Y | status of the room (e.g., Available, Booked, Maintenance, Unavailable) |
        | price | string | Y | price of the room per night |
        | description | string | Y | brief description of the room |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |

    - `GET /api/room`: Get detail room.
        ```sh
        curl --location 'http://localhost:3000/api/room?name=101' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIxNzM4MTd9.-xpqF6z2piPikSpXAuS_rpv34GhfGue7oyyLYyQbz7g'
        ```
        
        - Request Query Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | name | string | Y | name of the room |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |
        | id | integer | N | room id |
        | name | string | N | name of the room |
        | floor | integer | N | floor the room is located |
        | type | string | N | type of the room |
        | price | float | N | price of the room per night |
        | status | string | N | status of the room (e.g., Available, Booked, Maintenance, Unavailable) |
        | description | string | N | brief description of the room |
        
    - `PUT /api/room/:name`: Update room information
        ```sh
        curl --location --request PUT 'http://localhost:3000/api/room/201' \
        --header 'Content-Type: application/json' \
        --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImphbmVkb2UiLCJleHAiOjE3MzIyNTg5Nzl9.p0-u5qD6gt-sYPsVGEroZgFAhxcTRexuI92k3ByZ8IA' \
        --data '{
            "status": 2
        }'
        ```

         - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | status | string | Y | status of the room (e.g., 1=Available, 2=Booked, 3=Maintenance, 4=Unavailable) |

        - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | code | integer | Y | response http status |
        | message | string | Y | response message |
        | data | object | Y | response data |

4. **Payment Processing**
    - `PUT /api/payment`: Update payment status for a booking
5. **User Management**
    - `POST /api/user`: Create a new hotelier profile.
    - `PUT /api/password`: Change password for user account.
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