# Car Management Tool

Small application to manage fleet of vehicles.

Features planned at the moment:

- Add vehicles to the fleet
- Add refuellings to vehicleds to track fuel consumption

## Build and run

The software consists of a golang backend and a vuejs frontend.

### Backend

The backend is located in the `./backend/` subdirectory. In this directory call

- `go get ./...` to install dependencies
- `make local_test_env` to start a local development environment
- `make local_server` to start the backend server.

To run tests, run
- `go get -t ./...` to install testing dependencies, too
- `go generate ./...` to generate mocks (note: if this command fails, rerun it and it should succeed)
- `make test_unit` or `make test_all` to run unit tests only or execute all tests (unittests + integration tests).


### Frontend

The frontend is located in the `./frontend/` subdirectory. In this directory call

- `yarn install` to install dependencies and
- `yarn serve` to run the development server.

To run unit tests use `yarn test:unit`.
