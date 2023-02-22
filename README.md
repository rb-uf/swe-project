# Rater-Gator

Rater-Gator is a webapp for reviewing things around UF campus (such as buildings, classrooms, and chairs).

Created for CEN3031 Intro to Software Engineering.

Frontend team:
- Devala (Nimai) Griffith :kangaroo: 
- Richard Bachmann :)

Backend team:
- Emmett Kogan :skull:
- Shane Ferrell :monkey:



## Announcements

- Backend setup process has been simplified.
- The backend is now serving the frontend.



## Setup Instructions

First, open a terminal, clone the repository, and enter the newly-created directory:
```
git clone https://github.com/rb-uf/swe-project.git
cd swe-project
```

### Frontend

1. Install node.js: https://nodejs.org/en/download/

2. Install Angular CLI: https://angular.io/guide/setup-local#install-the-angular-cli

3. Verify everything is working by running the application:
   ```
   cd angular-front-end
   ng serve --open
   ```
   The application should open in your web browser.
   This may take a few moments.

### Backend

1. Install Go: https://go.dev/doc/install

2. Enter the `backend/` directory (from within `swe-project/`) and build the project:
   ```
   cd backend
   go build
   ```
   This should automatically download any required dependencies as well.

3. Start the backend server:
   ```
   go run swe-project/backend
   ```
   Several lines should print out starting with `[GIN-debug]`.

4. Open a web browser and go to `localhost:3000`.

   You should see the front end.

5. From the terminal, ress `Ctrl-C` to stop the server.



## Developer Notes

### Backend

- If you want to import a package created within this project, add `"swe-project/backend/package-name"` to the import list.
  "swe-project/backend" is the name of the go module created for this project.

- gorilla/mux
  - The purpose of gorilla/mux is to take an HTTP request, pick it apart, and try matching it against "descriptions" of HTTP Requests.
  - If the request matches a particular description, the request is given to a "handler" function.
  - You make both the descriptions and the handler functions.
  - A description can involve things like "the requested URL should start with /app" or "only match POST requests".

- The http.ResponseWriter and http.Request stuff comes from net/http, which is part of Go's standard library.
  The only thing gorilla/mux provides is the pattern-matching functionality.
 
- Documentation on Go's http.Request data structure: https://pkg.go.dev/net/http#Request

- JSON 
  - Stands for "JavaScript Object Notation".
  - It is a text-based (string) data format for describing an object or struct.
  - It was inspired by Javascript's syntax for objects, but it's use is not limited to just with Javascript.
    It is widely-used as a general-purpose format for saving and sending data structures.
  - encoding/json is a library for converting between Go structs and JSON strings.
    It is part of Go's standard library.

- Useful parts of "The Go Programming Language":
  - Section 1.7, p19 - A Web Server
  - Section 4.5, p107 - JSON
  - Section 7.7, p191 - The http.Handler Interface


## Ideas

### Data Structures

These are data structures the backend could implement for basic functionality.
They can be modified or added-to later.
"Subject" refers to the subject of a review (for example, a classroom at UF).

subject:
- subject_id (int)
- name (string)
- average_rating (float)
- review_list (list of review_ids)

review:
- review_id (int)
- subject_id (int)
- rating (int)
- text (string)

### API

This is a description of a basic API the backend could implement.
The "create" actions could be implemented with an http POST request, and the "get" actions with an http GET request.
The data could be sent in a JSON format.

Create new subject
- client provides name
- server generates subject_id
- server initializes average_rating to zero
- server initializes review_list as an empty list

Create new review
- client provides subject_id, rating, and text
- server generates review_id
- server adds review_id to subject's review_list
- server recomputes subject's average_rating

Get list of subjects
- server returns a list with a subject_id, name, and average_rating for each subject

Get list of reviews
- client provides subject_id
- server returns a list with a review_id, rating, and text for each review
