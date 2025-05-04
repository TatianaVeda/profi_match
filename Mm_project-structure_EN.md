# Match-Me Project Structure

```
match-me/
│
├── backend/                   # Go backend
│   ├── cmd/                   # Application entry points
│   │   └── server/
│   │       └── main.go        # Server main file
│   │
│   ├── internal/              # Internal application code
│   │   ├── api/               # REST API endpoints
│   │   │   ├── handlers/      # Request handlers
│   │   │   ├── middleware/    # Middleware (JWT, logging)
│   │   │   └── routes.go      # Route registration
│   │   │
│   │   ├── auth/              # Authentication
│   │   │   ├── jwt.go
│   │   │   └── password.go    # Working with bcrypt
│   │   │
│   │   ├── config/            # Application configuration
│   │   │   └── config.go
│   │   │
│   │   ├── database/          # Database operations
│   │   │   ├── migrations/    # DB migrations
│   │   │   └── postgres.go    # PostgreSQL connection
│   │   │
│   │   ├── models/            # Data models
│   │   │   ├── user.go
│   │   │   ├── profile.go
│   │   │   ├── bio.go
│   │   │   ├── location.go
│   │   │   ├── connection.go
│   │   │   └── message.go
│   │   │
│   │   ├── repository/        # Data access layer
│   │   │   ├── user_repo.go
│   │   │   ├── profile_repo.go
│   │   │   ├── connection_repo.go
│   │   │   ├── chat_repo.go
│   │   │   └── recommendation_repo.go
│   │   │
│   │   ├── service/           # Business logic
│   │   │   ├── user_service.go
│   │   │   ├── auth_service.go
│   │   │   ├── profile_service.go
│   │   │   ├── connection_service.go
│   │   │   ├── chat_service.go
│   │   │   └── recommendation_service.go
│   │   │
│   │   ├── utils/             # Helper functions
│   │   │   ├── validator.go
│   │   │   └── location_utils.go
│   │   │
│   │   └── websocket/         # WebSocket implementation for chat
│   │       ├── client.go
│   │       ├── hub.go
│   │       └── message.go
│   │
│   ├── pkg/                   # Public packages for potential reuse
│   │   └── geolocation/
│   │       └── distance.go    # Calculating distances between coordinates
│   │
│   ├── scripts/               # Development and deployment utilities
│   │   ├── seed.go            # Test data generation
│   │   └── reset_db.go        # DB reset/reload
│   │
│   ├── .env                   # Environment variables
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── frontend/                  # React frontend
│   ├── public/
│   │   ├── favicon.ico
│   │   ├── index.html
│   │   └── assets/
│   │       └── placeholder.png # Avatar placeholder
│   │
│   ├── src/
│   │   ├── api/               # API operations
│   │   │   ├── auth.js        # Authentication
│   │   │   ├── profile.js     # Profile
│   │   │   ├── connections.js # Connections
│   │   │   ├── recommendations.js
│   │   │   └── chat.js        # Chat
│   │   │
│   │   ├── components/        # React components
│   │   │   ├── Auth/
│   │   │   │   ├── Login.jsx
│   │   │   │   └── Register.jsx
│   │   │   │
│   │   │   ├── Profile/
│   │   │   │   ├── ProfileForm.jsx
│   │   │   │   ├── PhotoUpload.jsx
│   │   │   │   └── ViewProfile.jsx
│   │   │   │
│   │   │   ├── Recommendations/
│   │   │   │   ├── RecommendationList.jsx
│   │   │   │   └── RecommendationCard.jsx
│   │   │   │
│   │   │   ├── Connections/
│   │   │   │   ├── ConnectionList.jsx
│   │   │   │   └── RequestList.jsx
│   │   │   │
│   │   │   ├── Chat/
│   │   │   │   ├── ChatList.jsx
│   │   │   │   ├── ChatWindow.jsx
│   │   │   │   ├── Message.jsx
│   │   │   │   └── TypingIndicator.jsx
│   │   │   │
│   │   │   └── UI/
│   │   │       ├── Navbar.jsx
│   │   │       ├── Footer.jsx
│   │   │       ├── Button.jsx
│   │   │       └── Modal.jsx
│   │   │
│   │   ├── context/           # React context
│   │   │   ├── AuthContext.jsx
│   │   │   └── ChatContext.jsx
│   │   │
│   │   ├── hooks/             # Custom hooks
│   │   │   ├── useAuth.js
│   │   │   ├── useWebSocket.js
│   │   │   └── useGeolocation.js
│   │   │
│   │   ├── pages/             # Application pages
│   │   │   ├── LoginPage.jsx
│   │   │   ├── RegisterPage.jsx
│   │   │   ├── ProfilePage.jsx
│   │   │   ├── RecommendationsPage.jsx
│   │   │   ├── ConnectionsPage.jsx
│   │   │   └── ChatPage.jsx
│   │   │
│   │   ├── utils/             # Helper functions
│   │   │   ├── validators.js
│   │   │   ├── formatters.js
│   │   │   └── localStorage.js
│   │   │
│   │   ├── App.jsx            # Main component
│   │   ├── index.jsx          # Entry point
│   │   └── router.jsx         # Routing
│   │
│   ├── .env                   # Environment variables
│   ├── package.json
│   ├── package-lock.json
│   └── Dockerfile
│
├── docker-compose.yml         # Docker Compose configuration
├── .gitignore
├── README.md                  # Project documentation
└── docs/                      # Additional documentation
    ├── api.md                 # API documentation
    ├── deployment.md          # Deployment instructions
    └── algorithm.md           # Recommendation algorithm description
```

## Main Components Description

### Backend (Go)

1. **cmd/server** - application entry point
2. **internal** - internal application packages:
   - **api** - REST API and handlers
   - **auth** - authentication and JWT operations
   - **models** - data structures
   - **repository** - database operations (Repository pattern)
   - **service** - business logic (Service pattern)
   - **websocket** - chat implementation via WebSocket

3. **pkg** - public packages that can be reused
4. **scripts** - utilities for test data generation and DB management

### Frontend (React)

1. **src/api** - functions for backend API operations
2. **src/components** - React components by functional groups
3. **src/context** - contexts for global state management
4. **src/hooks** - custom hooks for logic reuse
5. **src/pages** - main application pages

### Project Configuration

1. **docker-compose.yml** - container configuration (Go, React, PostgreSQL)
2. **docs** - project documentation
3. **.env** files - environment variables for different parts of the project

## Database Structure (Main Tables)

- **users** - system users
- **profiles** - profiles with basic information
- **bios** - biographical data for matching algorithm
- **locations** - geographical data of users
- **connections** - established connections between users
- **connection_requests** - connection requests
- **messages** - chat messages
- **recommendations** - recommendations cache 