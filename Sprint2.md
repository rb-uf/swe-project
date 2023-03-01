# Sprint 2

## Work completed in sprint 2
### Frontend
+ We completed an overhaul of our frontend to be compatible with form submission to the local data structure for list of reviews.
+ We created a review service for accessing reviews.
+ We implemented create functionality for new reviews.
+ We implemented communication with the backend (post review request).
### Backend
+

## Links
+ https://youtu.be/sX0GyUzVErc

## Frontend unit tests (Cypress)
+ Per the guidelines of a TA, created 3 tests in cypress. 
+ One fills out the form completely and submits a review successfully. 
+ The second enters a string into the integer-only rating field and fails. 
+ The third only partially fills out the review but successfully adds it. 

## Backend unit tests
+ 

## Documentation for backend API

There are two types of objects the backend is designed to handle: "subjects" and "reviews".

"Subject" refers to the subject of a review, such as a classroom or location on campus.


### Create Subject

- http request type: `POST`
- url: `/create-subject`
- body: (json)
  ```
  {
      "Name": "Example subject"
  }
  ```

### Create Review

- http request type: `POST`
- url: `/create-review`
- body: (json)
  ```
  {
      "Subject": "Example subject",
      "Rating": 5,
      "Text": "It's alright.",
      "Author": "My name",
      "AuthorID": 1234
  }
  ```

### Get Subjects

- http request type: `GET`
- url: `/get-subjects`
- returns: A json array of subjects:
  ```
  [
    { "Name": "subject1" },
    { "Name": "subject2" },
    { "Name": "subject3" }
  ]
  ```

### Get Subject by Name

### Delete Subject

- http request type: `DELETE`
- url: `/delete-subject`
- body:
```
<subject object JSON, e.g. the ones returned by Get Subjects>
```
- returns: http.StatusOK (200)

### Get Subject Reviews

- http request type: `GET`
- url: `/get-subject-reviews`
- body:
```
<subject object JSON, e.g. the ones returned by Get Subjects>
```
- returns: A json array of reviews

### Delete Review

- http request type: `DELETE`
- url: `/delete-review`
- body:
```
<review object JSON, e.g. the ones returned by Get Subject Reviews>
```
- returns: http.StatusOK (200)

### Update Review

- http request type: `PUT`
- url: `/update-review`
- body:
```
<review object JSON, e.g. the ones returned by Get Subject Reviews>
```
- returns: http.StatusOK (200)
