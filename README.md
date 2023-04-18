# Rater-Gator

```
 ____       _                  ____       _             
|  _ \ __ _| |_ ___ _ __      / ___| __ _| |_ ___  _ __ 
| |_) / _` | __/ _ \ '__|____| |  _ / _` | __/ _ \| '__|
|  _ < (_| | ||  __/ | |_____| |_| | (_| | || (_) | |   
|_| \_\__,_|\__\___|_|        \____|\__,_|\__\___/|_|   
                                                        
```

Rater-Gator is a webapp for reviewing things around UF campus (such as buildings, classrooms, and chairs).

Created for CEN3031 Intro to Software Engineering.

Frontend team:
- Devala (Nimai) Griffith :kangaroo: 
- Shane Ferrell :monkey:

Backend team:
- Emmett Kogan :skull:
- Richard Bachmann :)


## Setup Instructions

1. Install node.js: https://nodejs.org/en/download/

2. Install Angular CLI: https://angular.io/guide/setup-local#install-the-angular-cli

3. Install Go: https://go.dev/doc/install

4. Clone the repository, and enter the newly-created directory:
   ```
   git clone https://github.com/rb-uf/swe-project.git
   cd swe-project
   ```
5. From `frontend/rater-gator/`, build the Angular application:
   ```
   ng build
   ```
6. From `backend/`, run:
   ```
   sh run.sh
   ```
   This should automatically download dependencies, compile the Go code, and start the backend.

7. Open a web browser and go to `localhost:3000`.

   You should see the front end.

8. `Ctrl-C` to stop the server on the terminal it was started on.


## Backend Options

The backend program accepts several command-line options:
- `-db _` -- specify database file path (default `datamgr/database.db`)
- `-frontend _` -- frontend dist path (default `../frontend/rater-gator/dist/rater-gator/`)
- `-port _` -- specify port number on which to listen for requests (default `3000`)
- `-help` -- list available command line options
