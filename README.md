# my-go-app

## Test (Local)

`make run-test`

## Run (Docker)

`docker-compose build`
`docker-compose up`
API Document: `http://localhost/swagger/index.html`

## Deployment (Read-only)

* Deployment on Heroku
  * Commit and push codes on `https://github.com/tonyltf/my-go-app`
  * Login Heroku Dashboard, create/select a project
  * In Deploy section, select GitHub and connect to the repo
  * Run Deploy Branch under Manual deploy
  
API Document: `https://my-go-app.herokuapp.com/swagger/index.html`

## Background

### Stack

* Go-Chi is the suggested framework
* SQLite was planned to use originally because it is easy to setup for local hosting.  But then deployed on Heroku and found it is not supported
* Postgres is the selected database because I want to enable transforming different source into the same schema so it will be easier for querying the price / average price
* Didn't plan to use NoSQL because the data insertion looks not very frequently

## Enhancement

### Business

* Better handling of exchange rate not found, I think it is kind of business/product requirement because of a few concerns
  * data in the current database does not reflect the complete history of exchange rate
  * if there is another third party service supports complete rate history by specifying the backdate timestamp, those API maybe costly so I would suggest adding a fallback mechanism to lookup and store when needed
* API to get price with timestammp - if a future time is provided, we only shows the last price. Should have better handling for this

### System

* App structure - I am new to it so didn't organized the folder well
* Services - Split the cron job and web server into 2 different services
* Configure file - Better usage , i.e. for different environment like .env in JavaScript projects
* Monitoring - if the currenct third party API is down or any error happens, i.e. rate-limit reached
* Error handling - Should have centralized error logging / handling so we can do something like alert / notification
* Logging - better loggin mechanism like fluentd
* Data pipeline - more decoupled data pipeline like ELT instead of directly dumping into the database
* Some testing is missing - factory pattern is not tested
