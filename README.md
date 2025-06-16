This is an event driven program in golang where the publisher publishes a message and the subscriber consumes it. You would need docker installed to run this project.

To run this project:
1. Run this command in your terminal - docker run -d --name rabbitmq -p 5672:5672
2. Next cd to the cmd directory
3. cd into the consumer directory and run the command : go run main.go
4. Do the same for the producer directory too.
5. When both are running, you would observe that in the comsumer terminal, messages will be received.
