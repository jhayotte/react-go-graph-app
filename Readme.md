# Microservices Dependency Graph Visualization

This project visualizes the dependencies between microservices in a large company using a React application and a Golang backend. The React frontend uses `react-flow` to build interactive graph nodes, while the Golang backend provides the necessary data.

## Features

- Interactive graph visualization using `react-flow`
- Custom node designs representing different microservices and databases
- Dynamic fetching of nodes and edges data from the Golang backend
- Display traffic information between services

## Installation

### Prerequisites

- [Node.js](https://nodejs.org/en/download/) (v14 or higher)
- [Go](https://golang.org/doc/install) (v1.16 or higher)

### Clone the Repository

```sh
git clone https://github.com/yourusername/microservices-dependency-graph.git
cd microservices-dependency-graph
```

#### Frontend Setup
Navigate to the frontend directory:

```sh
cd frontend
```

Install the dependencies
```sh
npm install
```

Build the React application:
```sh
npm run build
```

#### Backend Setup
Return to the root directory:

```sh
cd ..
```
Run the Golang server:

```sh
go run main.go
```

Run the Application
Open your browser and navigate to http://localhost:8080 to see the application.

## Usage
The application displays a graph where each node represents a microservice.
Edges between nodes represent traffic between the services, with labels indicating the amount of traffic.
Custom icons are used to visually differentiate between different types of services (e.g., authentication, user, billing, and database).

## Project Structure

```csharp
microservices-dependency-graph/
│
├── frontend/            # React frontend application
│   ├── public/          # Public assets
│   │   └── icons/       # Icons for microservices
│   ├── src/             # Source files
│   │   ├── App.js       # Main React component
│   │   ├── App.css      # Custom styles
│   │   └── index.js     # Entry point
│   ├── package.json     # NPM dependencies and scripts
│   └── ...              # Other React-related files
│
├── main.go              # Golang backend server
├── go.mod               # Golang dependencies
└── README.md            # Project README
```

## Adding New Microservices
Add new icons for the microservices in the public/icons directory.
Update the main.go file to include new nodes and edges.
Rebuild the React application and restart the Golang server.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes or improvements.

## License
This project is licensed under the MIT License.