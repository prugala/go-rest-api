Golang (Gin framework) Rest API with Oauth, JWT and Swagger boilerplate
-

- oauth auth - with google, but you can easy change it on other oauth provider
- authorization with JWT token - token generated after successful google login
- API Docs with swagger available under `/swagger/index.html`
    - update docs with `swag init` command

How to run:
- run `cp oauth.google.json.dist oauth.google.json`
- run `cp .env.dist .env`
- type google app data to file: `oauth.google.json`
- modify `.env` file
- run `go run .`

TODO:
- integrate Sentry https://github.com/getsentry/sentry-go/tree/master/gin
- tests
- Github actions
- Makefile
- docker-compose
- add support for roles system
