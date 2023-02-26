
# Backend API Documentation

There are two types of objects the backend is designed to handle: "subjects" and "reviews".

"Subject" refers to the subject of a review, such as a classroom or location on campus.


## Create Subject

- http request type: `POST`
- url: `/create-subject`
- body: (json)
  ```
  {
      "Name": "Example subject"
  }
  ```

## Create Review

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

## Get Subjects

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

## Get Subject by Name

## Delete Subject

## Get Subject Reviews

## Delete Review

## Update Review
