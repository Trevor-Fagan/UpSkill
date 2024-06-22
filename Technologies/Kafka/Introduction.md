## Apache Kafka Introduction

### Learning Material: [Confluent](https://www.youtube.com/playlist?list=PLa7VYi0yPIH0KbnJQcMv5N9iW8HkZHztH)

## What is it?
An event streaming platform used to collect, store, and process real time data streams at scale. Kafka
works through events, which is simply a collection of data at a point in time. Events are stored in Kafka
as a key-value pair. The key in the kafka message is not needed to be unique.

## Hands On - Confluent Cloud
Sign up for an account [here](https://confluent.cloud/home). You can use code *kafka101* for extra free
credits. Create a new topic called "poems" with the default 6 partitions. From here you can create a new
message to this topic from the UI.

### Partitions
Topics are broken down into what is known as partitions. When writing a message to a topic, the message is 
stored in one of the partitions. The partition that it is routed to is determined by the key of the message.

## Confluent CLI
### Setup
You can install the CLI using this command
`curl -sL --http1.1 https://cnfl.io/cli | sh -s -- latest`
Then you need to login using the following
`confluent login --save`
To select your environment, you can use the following with the environment you set up in the UI
`confluent environment use env-9zyr80`
Then select your cluster
`confluent kafka cluster use lkc-p5qp95`
Then you can either use an existing API key and secret
`confluent api-key store --resource lkc-p5qp95`
Or create a new one
`confluent api-key use <API_Key>`

### Useful Commands
- confluent kafka topic list                          => list all topics for the cluster you are connected to
- confluent kafka topic consume --from-beginning poem => consume from topic from the beginning
- confluent kafka topic produce poems --parse-key     => open window to produce key, value pairs to topic

You should be able to use the above commands to both produce and consume from the topics in your local terminal.
From the UI, you will also see all of the messages you produce reflected.

## Topics
A topic is Kafka's fundamental unit of event organization. You can think of a topic as a table in a relation databse;
it stores similarly structured data under a given name. Systems often use many topics and some topics can be duplicates
or transformations from other topics. A topic can also be thought of as a log of events; they are append only and
durable. They can only be read by seeking to an offset, they are not indexed. The events in the log are immutable. Events
in topics are configured to expire after a certain amount of time (1 day, 1 week, 1 month, 1 year, etc.) or by size. 