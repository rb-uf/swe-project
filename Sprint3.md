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
    "ID": 1234
    "Subject": "Example subject",
    "Rating": 5,
    "Text": "It's alright.",
    "Author": "Author's name",
    "AuthorID": 5678
}
```

| Action | HTTP request type | URL | Request body | Return value |
| --- | --- | --- | --- | --- |
| Create subject | `POST` | `/create-subject` | JSON `{ "Name": "Subject name" }` |

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
- body: JSON `{ "Name": "Subject name" }`
- return: http.StatusOK (200)

### Get all of Subject's Reviews
- http request type: `GET`
- url: `/get-subject-reviews`
- body: JSON `{ "Name": "Subject name" }`
- return: JSON array of reviews

### Delete Review
- http request type: `DELETE`
- url: `/delete-review`
- body: JSON `{ "ID": 1234 }`
- return: http.StatusOK (200)

### Update Review
- http request type: `PUT`
- url: `/update-review`
- body: Review JSON
- returns: http.StatusOK (200)
