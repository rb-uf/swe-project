# Rater-Gator

Rater-Gator is a webapp for reviewing things around UF campus (such as buildings, classrooms, and chairs).

Created for CEN3031 Intro to Software Engineering.

Frontend team:
- Devala (Nimai) Griffith :kangaroo: 
- Shane Ferrell :monkey:

Backend team:
- Emmett Kogan :skull:
- Richard Bachmann :)


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

### Features
- (Create) Users should be able to add locations, and leave reviews for 
  elements within them. For example, a user should be able to search
  a lecture hall, and then leave a review for the chairs within it.
- (Read) They should also be able to search for existing locations, and look
  at reviews for that location, or leave a new review
- (Create) When looking at locations, they should be able to leave reviews of 
  their own
- (Update) Users should be able to edit their posts/reviews
- (Delete) Users/moderators should be able to delete posts
- (Read) Should be able to filter search by elements, ratings, or reviews;
  e.g. they can query for the room with the best chairs. (For this to work
  better we may want to add specific categories or do some processing of user
  strings)
- Basic login screen and account system (just a name like When2Meet)

### Possible Features
- (Read) Maybe implement a map, where locations/rooms can be displayed, maybe
  an information sidebar is also displayed (including a picture of the location
  if available, and some type of summary of statistics for the ratings)
- Functionality for users to be able to post pictures within their reviews. For
  instance, so that if someone is reviewing a chair they can include a picture for
  other users
- The ability to comment on other people's ratings or reviews
- Sophisticated login screen and account system (actual authentication)

### Possible Components
1. https://github.com/JacobSamro/ngx-spellchecker
2. https://timdeschryver.dev/blog/google-maps-as-an-angular-component#angular

[Postman Link](https://app.getpostman.com/join-team?invite_code=93eeacce0902283fe854b0f8cedadc87&target_code=26326ce816c118a21b7b8f868fab1030)

### Stucture
We plan to utilize REST API for fully-reactive communication
