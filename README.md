## CatCMS-Go

Backend Golang microservice for CatCMS. Designed to run as a single backend for any number of CatCMS frontends. Designed to be hosted with Google Cloud Run, but it could be hosted easily in any container-based offering.

## Database

Currently designed to work with Google Datastore NoSQL. Chosen because it is cheap and covers the basic requirements of storage needed for this service.

## Local Setup

- Create a service account key 
  - https://cloud.google.com/docs/authentication/getting-started
  - Save JSON file to local machine
- Create a `.env` file in root

```
GOOGLE_APPLICATION_CREDENTIALS="${path to file with JSON key}"
APP_ID=${GUID unique to app instance}
JWT_SECRET=${32 character string}
BASIC_AUTH_USERNAME=${username}
BASIC_AUTH_PASSWORD=${password}
```

## Run Dev

- Start the service with `go run .`
- Service runs at `localhost:8080`