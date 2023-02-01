# Infrastructure

It is a driven layer (driven by core/domain).

There we implement adapters of the core layer and provide additional functionality to our domain.
For example database, cache, mailing service, notification service, message queue etc.

Adapters should ensure that we can change our tech stack at any time. For example replace RabbitMQ by Kafka as a message queue.