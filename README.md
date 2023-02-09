# Rater-Gator

Rater-Gator is an app for reviewing things (buildings, classrooms, chairs, etc.) around UF campus.

Created for CEN3031 Intro to Software Engineering.

Frontend team:
- Devala (Nimai) Griffith :kangaroo: 
- Richard Bachmann :)

Backend team:
- Emmett Kogan :skull:
- Shane Ferrell :monkey:

## Setup Instructions

### Frontend

1. Install node.js: https://nodejs.org/en/download/

2. Install Angular CLI: https://angular.io/guide/setup-local#install-the-angular-cli

3. Verify everything is working by running the application:
   ```
   cd angular-front-end
   ng serve --open
   ```
   The application should open in your web browser. This may take a few moments.

### Backend

## Internal Design



### Data Structures

subject:
- subject_id (int)
- name (string)
- average_rating (float)
- review_list (list of review_ids)

review:
- review_id (int)
- subject_id (int)
- rating (int)
- text (string)

### API

Create new subject
- client provides name
- server generates subject_id
- server initializes average_rating to zero
- server initializes review_list as an empty list

Create new review
- client provides subject_id, rating, and text
- server generates review_id
- server adds review_id to subject's review_list
- server recomputes subject's average_rating
