# Go Api Boilerplate

Basic Api Template
## What's Included

- [x] Versioning
- [x] Pagination
- [x] Configuration
- [x] Logging
- [x] Error Handling
- [x] Validation
- [x] Health Check
- [x] Data Seed
- [x] DB Migration
- [x] Run & Manage via CLI Command
- [x] Makefile
- [x] Linter
- [x] Unit Test
- [x] Docker & Docker Compose
- [x] Integration Test sample

## Get Started

Below are the steps if you want to run locally without docker

1. Set required environment variable

    ENV=local

2. Set configuration

    Change config/local.yaml configuration value properly
    and make sure can connect to blank MySQL database properly

3. Run migration

    > make migrate

4. Add seed data
    
    > make seed

5. Run app
    > make run

6. Open the browser 

    Visit the url
    > http://localhost:8080/v1/users

## Get Started Using Docker Compose

Below are the steps if you want to run locally with docker & docker compose

1. Build docker image

    Generate the docker imgae by this command:
    > docker build --rm -t starterkitapi -f dockerfile.api .
                                                                      
2. Run with docker compose
    
    Copy the sample.env to .env and adjust it then run the docker compose with this command:
    > docker-compose up                
                                                                               
3. Open the browser 

    Visit the url
    > http://localhost:8080/v1/users
                                                                    
4. Quit & Cleanup
    
    Click Ctrl+C to quit in console then run these commands below to clean up all things: 
    > docker-compose rm -v


