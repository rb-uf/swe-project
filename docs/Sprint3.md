# Sprint3

## Demo Video
Sprint 3 video link: [https://youtu.be/NxfM-YV4Q2A](https://youtu.be/NxfM-YV4Q2A)

## Work Completed
### Frontend
#### Subjects Functionality
The frontend now has functionality to filter reviews by their subject. This is normally the location of the classroom being reviewed, but subjects could also be any item the user wants to leave a review for or read reviews about.

Users can create and submit new subjects in the subjects section. Users may also select multiple subjects from a dropdown and the reviews with the subjects selected will be displayed in the reviews component.

This component uses the API actions "Create Subject", "Get Subject", and "Get Subject Reviews". 

#### Styling
The frontend now displays Angular Material components for styling. The header is now a Material toolbar. The subjects dropdown uses a Material select component with  Material checkboxes. The subject create submission button uses a Material button. Material components have yet to be added to form fields.

### Backend
#### User Authentication
The backend is now capable of authenticating a user and maintaining a session cookie.

The bcrypt library is used to hash the passwords provided by users.
Only the hash is stored in the database, not the plaintext password itself, for when our database is eventually hacked.

A "Cookie Jar" is maintained to keep track of logged-in users.
When a user logs in, a hash of the password they provide is compared to the hash stored in the database.
If they match, a cookie is created.
The cookie is both passed to the client and stored in our cookie jar.
In any HTTP requests requiring authentication, the cookie must be provided so it can be looked up in the cookie jar.
When a logout request is made, the cookie is removed from the cookie jar.

For easier testing with postman and to maintain current front end functionality which does not deal with logging in, the functions used to verify cookies within handler functions are setup to take advantage of conditional compilation, using the real version in the default mode and in the alternate build configuration it replaces it with a dummy function that says the cookie was always verified.

During the next sprint, we intend on working with the frontend team to set up the login page.

#### Improved API Documentation
The API Documentation from last sprint has been reorganized into an easy-to-read table (see below).
The information was also updated to account for API changes, and additional information was added to give insight into how the backend operates.
The documentation does not yet include information on authentication.

#### http-request.sh - A shell script for easy API tests
A shell script for making http requests was created for easier API testing.
Example usage:
```
./http-request.sh POST create-subject '{ "Name": "Marston" }'
```
The script is a thin wrapper around Curl with the output piped through a JSON formatter to make the response easy to read.

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
