# rater-gator

## Group Members

### Frontend
- Devala (Nimai) Griffith :kangaroo: 
- Richard Bachmann :)

### Backend
- Emmett Kogan :skull:
- Shane Ferrell :monkey:

## Project Description

### Definite Features
- (Create) Users should be able to add locations, and leave reviews for 
  elements within them. For example, a user should be able to search
  a lecture hall, and then leave a review for the chairs within it.
- (Read) They should also be able to search for existing locations, and look
  at reviews for that location, or leave a new review
- (Create) When looking at locations, they should be able to leave reviews of 
  their own
- (Update) Users should be able to edit their posts/reviews
- (Delete) Users/moderators should be able to delete posts
- (Read) Should be able to filter search by elements, ratings, or reviews;
  e.g. they can query for the room with the best chairs. (For this to work
  better we may want to add specific categories or do some processing of user
  strings)
- Basic login screen and account system (just a name like When2Meet)

### Possible Features
- (Read) Maybe implement a map, where locations/rooms can be displayed, maybe
  an information sidebar is also displayed (including a picture of the location
  if available, and some type of summary of statistics for the ratings)
- Functionality for users to be able to post pictures within their reviews. For
  instance, so that if someone is reviewing a chair they can include a picture for
  other users
- The ability to comment on other people's ratings or reviews
- Sophisticated login screen and account system (actual authentication)

### Possible Components
1. https://github.com/JacobSamro/ngx-spellchecker
2. https://timdeschryver.dev/blog/google-maps-as-an-angular-component#angular

[Postman Link](https://app.getpostman.com/join-team?invite_code=93eeacce0902283fe854b0f8cedadc87&target_code=26326ce816c118a21b7b8f868fab1030)


### Stucture
We plan to utilize REST API for fully-reactive communication

## Instructions

### Angular Front-end

1. Install node.js: https://nodejs.org/en/download/

2. Install Angular CLI: https://angular.io/guide/setup-local#install-the-angular-cli

3. Verify everything is working by running the application:
   ```
   cd angular-front-end
   ng serve --open
   ```
   The application should open in your web browser.
