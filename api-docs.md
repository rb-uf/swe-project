
# Backend API Documentation

There are two types of objects the backend is designed to handle: "subjects" and "reviews".

"Subject" refers to the subject of a review, such as a classroom or location on campus.


## Create Subject

- http request type: `POST`
- url: `/create-subject`
- body (json):
  ```
  {
      "Name": "Example subject"
  }
  ```

## Create Review

- http request type: `POST`
- url: `/create-review`
- body (json):
  ```
  {
      "Subject": "Example subject",
      "Rating": 5,
      "Text": "It's alright.",
      "Author": "My name",
      "AuthorID": 1234
  }
  ```
