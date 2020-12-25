Golang (Gin framework) Rest API with Oauth, JWT and Swagger boilerplate
-

- oauth auth - with google, but you can easy change it on other oauth provider
- authorization with JWT token - token generated after successful google login
- API Docs with swagger available under `/swagger/index.html`
    - update docs with `swag init` command
- Live reload - all changes in files: `*.go, .env and oauth.google.json` are detected and the app is automatically rebuilt

How to run:
- run `docker-compose up`

TODO:
- integrate Sentry https://github.com/getsentry/sentry-go/tree/master/gin
- tests
- Github actions
- Makefile
- add support for roles system
