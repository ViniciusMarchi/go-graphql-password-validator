FROM golang:1.19.4-bullseye
# create a directory named /app
RUN mkdir /app
# copy project code to /app/ dir
COPY . /app/
WORKDIR /app
# download and install the necessary dependencies
RUN go mod download
# build project
RUN cd server && go build -o /main .
EXPOSE 8080
# run api
CMD ["/main"]