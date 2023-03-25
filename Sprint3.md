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

| Action | HTTP request | URL | Request body | Return value |
| --- | --- | --- | --- | --- |
| Create subject | `POST` | `/create-subject` | `{ "Name": "Subject name" }` | |
| Create Review | `POST` | `/create-review` | Review JSON | |
| Get Subjects | `GET` | `/get-subjects` | | JSON array of Subjects |
| Delete Subject | `DELETE` | `/delete-subject` | `{ "Name": "Subject name" }` | http.StatusOK (200) |
| Get all of Subject's Reviews | `GET` | `/get-subject-reviews` | `{ "Name": "Subject name" }` | JSON array of reviews |
| Delete Review | `DELETE` | `/delete-review` | `{ "ID": 1234 }` | http.StatusOK (200) |
| Update Review | `PUT` | `/update-review` | Review JSON | http.StatusOK (200) |
