FROM golang:1.22-alpine


# set the current working directory inside the container
WORKDIR /app

#copy go.mod and go.sum files first to take advantage of docker caching
COPY go.mod ./ 

# download all dependecies 
RUN go mod download

#copy the source code into the container 
COPY . . 

# build the go application
RUN go build -o optimus_app

#expose the go app on port 8080
EXPOSE 8080

# run the executable
RUN ["./optimus_app"]