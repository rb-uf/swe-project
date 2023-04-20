# Sprint 4

## Demo Video
[https://youtu.be/I4ba3PUEHiU](https://youtu.be/I4ba3PUEHiU)

## Work Completed
### Frontend
- Fixed selector to render review list display.
- Added edit review description toggle, edit review display, and edit review HTTP request to server.
- Added component routing to to the app module with default route '/'.
- Added page navigation with component routing.
- Moved new review page to /add-review route. AddReviewComponent created from review component.
- Added delete review functionality.
- Removed log messages for creating subjects and reviews.

### Backend
- Added command-line arguments for setting the database path, frontend path, and port number.
- Added handler functions for fetching data by author, whether it be just their posts or some basic statistics about their post history.
- Added functionality for users being able to like and dislike posts and track "Up" counters
- Added corresponding test functions for each thing implemented, as well as postman test cases.
- Updated README with new setup instructions.
- Added backend command-line options to README.
- Removed added database file to .gitignore because it was causing merge conflicts before.
- Changed print statements from fmt to log.
- Added log message to print contents of http requests.
- Updated API documentation with new API calls.

## Testing (Unit and Functional)
### Frontend
- Cypress
  - Test 1: Add a new location and verify that it has been added.
  - Test 2: Add a review and verify that it has been added.
  - Test 3: Add a second review and verify that it has been added.
  - Test 4: Delete the first review and verify that is has been deleted. 
  - Test 5: Delete the second review and verify that is has been deleted. 
### Backend
- In handler_test.go/handler_auth_test.go:
  - TestCreateSubject
  - TestCreateSubject_NoCookie
  - TestCreateReview
  - TestCreateReview_NoCookie
  - TestGetSubjects
  - TestGetSubjectReviews
  - TestGetReviewsBySubject
  - TestGetReviewsByAuthor
  - TestUpdateReview
  - TestUpdateReview_NoCookie
  - TestUpdateUps
  - TestDeleteSubject
  - TestDeleteSubject_NoCookie (intended to fail)
  - TestDeleteSubject_NotAdmin (intended to fail)
  - TestCreateUser
  - TestLogin
  - TestGetUserStats
- Through Postman:
  - Delete Subject via JSON body
  - Delete Review
  - Get All Subjects
  - Get Subject Reviews Test
  - Get Reviews by Subjects
  - Create Subject Test
  - Create Review Test
  - Update Review
  - Create User
  - Get User Stats
  - Update Ups (increment and decrement)
- In run-tests.sh with http-request.sh:
  - create-subject
  - create-review
  - get-subject-reviews

Note: there are now two classes of backend tests, one that tests cookies and authorization functionality on top of basic
CRUD functionality and another that just tests the basic crud functionality by utilizing conditional compilation for various implementations of test functions and helper functions. To run the new class of tests do `go test ./ -v` and for the previous versions use `go test -tags alt ./ -v`

## API Documentation

There are two types of objects the backend is designed to handle: "subjects" and "reviews".
"Subject" refers to the subject of a review, such as a classroom or location on campus.

### Subject JSON
```
{
    "Name": "Example Subject"
}
```

### Review JSON
```
{
    "ID": 1234
    "Subject": "Example subject",
    "Rating": 5,
    "Text": "It's alright.",
    "Author": "Author's name",
    "AuthorID": 5678
}
```

### Available HTTP Requests
| Action | Method | URL | Body | Return value |
| --- | --- | --- | --- | --- |
| Create User | `POST` | `/sign-up` | `{ "Username": "name", "Password": "12345" }` | http.StatusCreated |
| Login | `POST` | `/login` | `{ "Username": "name", "Password": "12345" }` | New cookie |
| Logout | `POST` | `/logout` | None (user cookie should be in header) | http.StatusOK |
| Get User Stats | `POST` | `/get-user-stats` | `{ "User": "name" }` | `{ "Posts": 10, "TotalScore": 22 }` |
| Create subject | `POST` | `/create-subject` | `{ "Name": "Subject name" }` | JSON of the new subject |
| Create Review | `POST` | `/create-review` | Review JSON, but without ID | JSON of the new review |
| Get Subjects | `GET` | `/get-subjects` | N/A | JSON array of Subjects |
| Delete Subject | `POST` | `/delete-subject` | `{ "Name": "Subject name" }` | http.StatusOK (200) |
| Get all of Subject's Reviews | `POST` | `/get-subject-reviews` | `{ "Name": "Subject name", "MaxReviews": 10000 }` | JSON array of reviews |
| Delete Review | `POST` | `/delete-review` | Review JSON | http.StatusOK (200) |
| Update Review | `PUT` | `/update-review` | `{ "ID": int, "Text": string }` | http.StatusOK (200) |
| Get Subjects by Author | `POST` | `/get-subjects-by-author` | `{ "Author": <string> }` | http.StatusOK |
| Update Review Ups | `PUT` | `/update-ups` | `{ "ReviewID": <ID number>, "Up": <negative for down, 0 or positive for up> }` | Review JSON |

### Additional Information
Subjects are referenced by their Name attribute and reviews are referenced by their ID.
However, when creating a new review, the an ID should not be supplied in the request.
GORM is responsible for generating a unique ID when the review is added to the database.

Technically, there should be no more than one subject of a given name, but the backend does not currently check for this when processing a create-subject request.
Adding multiple subjects with the same name does not seem to cause any problems (yet).
