## CatCMS-Go

Backend Golang microservice for CatCMS. Designed to run as a single backend for any number of CatCMS frontends. Designed to be hosted with Google Cloud Run, but it could be hosted easily in any container-based offering.

## Database

Currently designed to work with Google Datastore NoSQL. Chosen because it is cheap and covers the basic requirements of storage needed for this microservice.

## Local

- Create a service account key 
  - https://cloud.google.com/community/tutorials/cicd-cloud-run-github-actions
  - Save JSON file to repo root as `key.json`
- Create a `.env` file in root

```
GOOGLE_APPLICATION_CREDENTIALS="${path to file with JSON key}"
JWT_SECRET=${32 character string}
BASIC_AUTH_USERNAME=${username}
BASIC_AUTH_PASSWORD=${password}
```

## Dev

- Start the service with `go run .`
- Service runs at `localhost:8080`

## Deployment

- https://cloud.google.com/community/tutorials/cicd-cloud-run-github-actions
- `production` branch used for deployment
- Create the following GitHub Actions secrets

```
GCP_APP_NAME=${name of app}
GCP_CREDENTIALS=${content of JSON key}
GCP_EMAIL=${email from service account}
GCP_PROJECT_ID=${ID of GCP project}
GCP_REGION=${GCP region, for example `us-west2`}
JWT_SECRET=${32 character string}
BASIC_AUTH_USERNAME=${username}
BASIC_AUTH_PASSWORD=${password}
```

## Verify
- `gcloud run services list`
