
# go-json-crud-api

A simple RESTful API for managing movies written in Go, utilizing JSON for data storage and Gorilla Mux for routing.

## Overview

This API allows you to perform CRUD (Create, Read, Update, Delete) operations on movie resources.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Rizz-33/go-json-crud-api.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-json-crud-api
   ```

3. Build and run the application:

   ```bash
   go build
   ./go-json-crud-api
   ```

## Endpoints

### Get all movies

```
GET /movies
```

### Get a single movie

```
GET /movies/{id}
```

### Create a movie

```
POST /movies
```

Sample request body:

```json
{
  "isbn": "978-3-16-148410-0",
  "title": "Movie Title",
  "director": {
    "firstname": "Director's First Name",
    "lastname": "Director's Last Name"
  }
}
```

### Update a movie

```
PUT /movies/{id}
```

Sample request body:

```json
{
  "isbn": "978-3-16-148410-0",
  "title": "Updated Movie Title",
  "director": {
    "firstname": "Updated Director's First Name",
    "lastname": "Updated Director's Last Name"
  }
}
```

### Delete a movie

```
DELETE /movies/{id}
```

## Contributor

- Risini Amarathunga (@Rizz-33)
