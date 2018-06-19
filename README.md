In order to deploy this app to aws lambda the application has to be build first

GOOS=linux go build main.go

The resulting main executable needs to be compressed to a main.zip which can then be uploaded to lambda
