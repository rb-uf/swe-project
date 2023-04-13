# Sprint 4

## Demo Video

## Work Completed
### Frontend

### Backend
- Added command-line arguments for setting the database path, frontend path, and port number.

## Testing (Unit and Functional)
### Frontend
- Cypress
  - Test 1: Input a location and a full review with the correct location. Verify that all the inputs were added correctly, then check for errors. No errors show up. 
  - Test 2: Input a location and a full review with an incorrect location. Verify that all the inputs were added correctly, then check for errors. 1 error shows up.
  - Test 3: Submit a black review. No errors show up.
### Backend
- In handler_test.go/handler_auth_test.go:
  - TestCreateSubject
  - TestCreateSubject_NoCookie
  - TestCreateReview
  - TestCreateReview_NoCookie
  - TestGetSubjects
  - TestGetSubjectReviews
  - TestGetReviewsBySubject
  - TestUpdateReview
  - TestUpdateReview_NoCookie
  - TestDeleteSubject
  - TestDeleteSubject_NoCookie (intended to fail)
  - TestDeleteSubject_NotAdmin (intended to fail)
  - TestCreateUser
  - TestLogin
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
| Create subject | `POST` | `/create-subject` | `{ "Name": "Subject name" }` | JSON of the new subject |
| Create Review | `POST` | `/create-review` | Review JSON, but without ID | JSON of the new review |
| Get Subjects | `GET` | `/get-subjects` | N/A | JSON array of Subjects |
| Delete Subject | `DELETE` | `/delete-subject` | `{ "Name": "Subject name" }` | http.StatusOK (200) |
| Get all of Subject's Reviews | `GET` | `/get-subject-reviews` | `{ "Name": "Subject name" }` | JSON array of reviews |
| Delete Review | `DELETE` | `/delete-review` | `{ "ID": 1234 }` | http.StatusOK (200) |
| Update Review | `PUT` | `/update-review` | Review JSON | http.StatusOK (200) |

### Additional Information
Subjects are referenced by their Name attribute and reviews are referenced by their ID.
However, when creating a new review, the an ID should not be supplied in the request.
GORM is responsible for generating a unique ID when the review is added to the database.

Technically, there should be no more than one subject of a given name, but the backend does not currently check for this when processing a create-subject request.
Adding multiple subjects with the same name does not seem to cause any problems (yet).
