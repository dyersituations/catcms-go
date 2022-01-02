## CatCMS-Go

Backend Golang microservice for CatCMS. Designed to run as a single backend for any number of CatCMS frontends. Designed to be hosted with Google Cloud Run, but it could be hosted easily in any container-based offering.

## Database

Currently designed to work with Google Datastore NoSQL. Chosen because it is cheap and covers the basic requirements of storage needed for this microservice.

## Local

- Run to get credentials for local dev: `gcloud auth application-default login`
- Create a `.env` file in root

```
JWT_SECRET=${32 character string}
BASIC_AUTH_USERNAME=${username}
BASIC_AUTH_PASSWORD=${password}
```

## Dev

- Start the service with `go run .`
- Service runs at `localhost:8080`

## Deployment

- Create a service account
  - https://cloud.google.com/community/tutorials/cicd-cloud-run-github-actions
  - Use JSON key below for `GCP_CREDENTIALS`
- Use `production` branch for deployment
- Create the following GitHub Actions secrets

```
GCP_APP_NAME=${name of app, ie `catcms`}
GCP_CREDENTIALS=${content of JSON key}
GCP_EMAIL=${email of service account}
GCP_PROJECT_ID=${ID of GCP project}
GCP_REGION=${GCP region, ie `us-west2`}
JWT_SECRET=${random 32 character string}
BASIC_AUTH_USERNAME=${username}
BASIC_AUTH_PASSWORD=${password}
```

## Verify

- Show deployed services: `gcloud run services list`
