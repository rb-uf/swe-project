
# Backend API Documentation

There are two types of objects the backend is designed to handle: "subjects" and "reviews".

## Subjects

The subject of a review, such as a classroom or location on campus.

To create a new subject, the frontend should send the backend an http request of the form:
- http request type: `POST`
- url: `/create-subject`
- json body:
  ```
  {
      "Name": "<name-of-the-new-subject>"
  }
  ```



## Reviews

- *Review*
- *Subject* (the subject of a review, such as a classroom or location on campus)



## 
