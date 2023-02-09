## System requirements 
You need to have [Docker](https://www.docker.com) 
installed and open in order to build and run the application.
No additional tools required. 

## How to build and run
1. Create a Docker network:
    ```shell script
    docker network create chairRatings-net
    ```
2. Start the DB:
    ```shell script
    docker run \
      -e POSTGRES_USER=go \
      -e POSTGRES_PASSWORD=your-strong-pass \
      -e POSTGRES_DB=go \
      --name chairRatings-db \
      --net=chairRatings-net \
      postgres:11.5
    ```
3. Build the application image:
    ```shell script
    docker build -t chairratings-app .
    ```
4. Start the application container:
    ```shell script
    docker run -p 8080:8080 \
      -e DB_PASS='your-strong-pass' \
      --net=chairRatings-net chairratings-app
    ```
Access the application via http://localhost:8080

Source: https://github.com/Shpota/go-angular
