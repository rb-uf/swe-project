# Rater-Gator

Rater-Gator is a webapp for reviewing things around UF campus (such as buildings, classrooms, and chairs).

Created for CEN3031 Intro to Software Engineering.

Frontend team:
- Devala (Nimai) Griffith :kangaroo: 
- Shane Ferrell :monkey:

Backend team:
- Emmett Kogan :skull:
- Richard Bachmann :)


## Setup Instructions

First, open a terminal, clone the repository, and enter the newly-created directory:
```
git clone https://github.com/rb-uf/swe-project.git
cd swe-project
```
### Frontend

1. Install node.js: https://nodejs.org/en/download/

2. Install Angular CLI: https://angular.io/guide/setup-local#install-the-angular-cli

3. Verify everything is working by running the application:
   ```
   cd angular-front-end
   ng serve --open
   ```
   The application should open in your web browser.
   This may take a few moments.

### Backend

1. Install Go: https://go.dev/doc/install

2. Enter the `backend/` directory (from within `swe-project/`) and build the project:
   ```
   cd backend
   go build
   ```
   This should automatically download any required dependencies as well.

3. Start the backend server:
   ```
   go run swe-project/backend
   ```

4. Open a web browser and go to `localhost:3000`.

   You should see the front end.

5. From the terminal, press `Ctrl-C` to stop the server.

