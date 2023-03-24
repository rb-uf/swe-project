# Sprint3

## Demo Video

## Work Completed
### Frontend
-
### Backend
-

## Unit Tests
### Frontend
-
### Backend
-

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
    "Subject": "Example subject",
    "Rating": 5,
    "Text": "It's alright.",
    "Author": "Author's name",
    "AuthorID": 1234
}
```

### Create Subject
- http request type: `POST`
- url: `/create-subject`
- body: Subject JSON

### Create Review
- http request type: `POST`
- url: `/create-review`
- body: Review JSON

### Get Subjects
Returns a json array of subjects.
- http request type: `GET`
- url: `/get-subjects`
- return: JSON array of Subjects

### Delete Subject
- http request type: `DELETE`
- url: `/delete-subject`
- body: Subject JSON
- return: http.StatusOK (200)

### Get Subject Reviews
- http request type: `GET`
- url: `/get-subject-reviews`
- body: Subject JSON
- return: JSON array of reviews

### Delete Review
- http request type: `DELETE`
- url: `/delete-review`
- body: Review JSON
- return: http.StatusOK (200)

### Update Review
- http request type: `PUT`
- url: `/update-review`
- body: Review JSON
- returns: http.StatusOK (200)
