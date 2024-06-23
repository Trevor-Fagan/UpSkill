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
The purpose of partitions is to split a topic log into multiple logs. This makes it so reading, writing to,
storing, etc is able to be done across multiple nodes in the cluster. If there is no key associated with a
value being written, the messages will be spread across partitions in a round-robin fashion. If it does have
a key, the key is used to determine which partition to be used in. Kafka uses an internal hash fuction mod
the number of partitions to determine which partition to write to. This makes it so that any messages that
share the same key will always have the same partition number.

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
- confluent kafka topic describe poem                 => shows extra data about the topic
- confluent kafka topic create --partitions 1 poems_1 => create a new topic with 1 partition

You should be able to use the above commands to both produce and consume from the topics in your local terminal.
From the UI, you will also see all of the messages you produce reflected.

## Topics
A topic is Kafka's fundamental unit of event organization. You can think of a topic as a table in a relation databse;
it stores similarly structured data under a given name. Systems often use many topics and some topics can be duplicates
or transformations from other topics. A topic can also be thought of as a log of events; they are append only and
durable. They can only be read by seeking to an offset, they are not indexed. The events in the log are immutable. Events
in topics are configured to expire after a certain amount of time (1 day, 1 week, 1 month, 1 year, etc.) or by size.

## Brokers
Kafka is composed of a network of machines known as brokers. This may be physical hardware, containers running somewhere,
or in the cloud. They are independent machines each running the Kafka broker process. They are intentionally very simple.
They do manage partitions, replications of partitions, and handle read/write requests.

## Replication
Partition data is duplicated across brokers for redundancy. Replicas on brokers other than the primary are known as
follower replicas. The main partition is known as the leader replica. Writes and reads always happen to the leader.
Replication is handled by default by Kafka and the followers will automatically work. You are able to manage the settings
for these clusters to fine tune your durability, but most of the time developers do not handle this type of thing. If one
node in the cluster dies, another one will take over the role and the messages in a topic will still be there.

## Producers
A client application that is used for publishing messages to Kafka. Using the package for your specific language along
with configuration settings for network settings and others, you will create the key value pair that you are going to
send to Kafka. The producer makes the decision on which partition to produce messages to.

## Consumers
A client application that is used for consuming messages from Kafka. Simialr to producers, you will use the standard
Kafka consumer package to subscribe to one or more topics you are reading from. Consumers are significantly more
complicated than producers. Reading a message from a Kafka topic does *not* destroy it. A single instance of a consuming
application will automatically receive all of the messages from all partitions that it is subscribed to. If you create
a separate instance of a consumer, that triggers an automatic rebalancing process that distributes partitions fairly
between the two applications. This process happens anytime you add or remove a consumer group instance. You can have as
many consumers as you have partitions for a topic. If you have any more than this, the consumers will be idle. Each consumer
after the first one will be assigned to a particular partition and read from that.

A consumer group is a set of consumers that cooperate to consume data from some topics. This is how the above-mentioned
coordinate and rebalance when consumers are added/removed.