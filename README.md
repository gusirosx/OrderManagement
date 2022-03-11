# OrderManagement

## ``OrderManagement`` Service and Client

- Online retail scenario has a `` OrderManagement`` microservice which is responsible for managing the orders and
 their information. The consumer of that service can add, retrieve, search and update order via that service.

- ``OrderManagement`` service and the consumer of that service are implemented in both ``Go`` and ``Java`` languages.

- This sample shows how you can implement both ``OrderManagement`` service and its consumer.


gRPC offers a diverse set of RPC communication styles for building inter-process
communication between gRPC-based applications. In this chapter, we explored four
main communication patterns. Simple RPC is the most basic one; it is pretty much a
simple request–response style remote procedure invocation. Server-streaming RPC
allows you to send multiple messages from the service to the consumer after the first
invocation of the remote method, while client streaming allows you to send multiple
messages from the client to the service. We delve into the details of how we can
implement each of these patterns using some real-world use cases.
The knowledge you gained in this chapter is quite useful for implementing any gRPC
use case so that you can select the most appropriate communication pattern for your
business. While this chapter gave you a solid understanding of gRPC communication
patterns, the low-level communication details that are transparent to the user were
not covered in this chapter. In the next chapter, we will dive deep into how low-level
communication takes place when we have gRPC-based inter-process communication.
