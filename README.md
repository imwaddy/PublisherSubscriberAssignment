# GO RabbitMQ Assignment

The assignment is a solution to the problem statement mentioned in the problem.txt file.

<hr>

## Project Structure

```
- PublisherSubscriberAssignment

  -- PublisherModules: For sending messages
    --- publisher.go

  -- ConsumerModules: For receiving messages & storing in database
    --- helpers: Require library helpers for project
       --- confighelpers: Configurations related services
       --- dbhelpers: DB connection related services
       --- loghelpers: Log related services

    --- models: Database related models

    --- services: Database related models

    --- consumer.go: Application entry point
    --- config.json: Consumer related configurations
```

<hr>

## Running in Development
```
git clone https://github.com/imwaddy/PublisherSubscriberAssignment
cd PublisherSubscriberAssignment
```

Note: 1. Please verify the configurations in ```config.json``` before you start.
      2. Log file is going to create in ```PublisherSubscriberAssignment/ConsumerModules/``` named ```consumer.log```
### ConsumerModules - Listens to incoming message and stores in database
```
cd ConsumerModules
go mod tidy
go run consumer.go
```

### PublisherModules - Send message to Consumer
```
cd PublisherModules
go run publisher.go
```

Note: Make sure the Consumer is running before you run the Publisher.

<hr>

## Testing

A test case is available for database service in ```PublisherSubscriberAssignment/ConsumerModules/services``` named ```hotelService_test.go```.<br>
To run the test case:

```
cd PublisherSubscriberAssignment/ConsumerModules/services
go test hotelService_test.go hotelService.go
```

<hr>
