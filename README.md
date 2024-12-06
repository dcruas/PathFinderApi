# Pathfinder 2E API

An API for calculating the possible outcomes of dice rolls in Pathfinder 2E based on a given `modifier` and `DC` (Difficulty Class). Built using Go and the Gin framework.

---

## Features

- Calculate outcomes (`critical_failures`, `failures`, `successes`, and `critical_successes`) for a d20 roll.
- Efficient, loop-free computation of results.
- Dockerized for easy deployment.
- Unit tested for reliability.

---

## Table of Contents

- [Getting Started](#getting-started)

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.20 or higher)
- [Docker](https://www.docker.com/get-started) (optional for containerization)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/pathfinder-api.git
   cd pathfinder-api
   ```


2. Install Dependecies:
    ```bash
    go mod tidy
    ```

3. Run the API locally:
    ```bash
    go run main.go
    ```

4. Alternatively, build and run the Docker container:
    ```bash
    docker build -t pathfinder-api .
    docker run --name pathfinderapi -p 8080:8080 pathfinder-api
    ```
    

