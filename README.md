# Steps to reproduce

1. Setup venv
2. `docker-compose up localstack` # Remeber to set the LOCALSTACK API KEY since this is using lambda layers
3. `samlocal build`
4. `samlocal build --config-env local --guided`
5. `docker-compose up go`

### When running the 5 step it will invoke the lambda, so the lambda container will be created. Analyse the logs in Docker dashboard
### With `localstack-pro:latest` it occurs an error of permission denied 
### With `localstack-pro:1.4` it works as expected. Saving the url content to a pdf file at `/tmp/12345.pdf` 