# Pills reminder service stack

## Generate swagger

Run next command from root directory:

`swag init -g ./apps/mobileapi/main.go -o docs/swagger/mobileapi`

## Building images

Run next command from root directory:

`docker build -t reminder_mobileapi:latest -f apps/mobileapi/Dockerfile .`

`docker build -t reminder_reminderservice:latest -f apps/reminderservice/Dockerfile .`

`docker build -t reminder_pillsservice:latest -f apps/pillsservice/Dockerfile .`

## Run stack

Run next command from `deploy` directory:

`docker stack deploy PILLSREMINDER --compose-file docker-compose.yml`

## Drop stack

`docker stack rm PILLSREMINDER`