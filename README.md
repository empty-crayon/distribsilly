# Distribsilly

A Silly implementation of distributed key value database.

Distribsilly is a lightweight and straightforward implementation of a distributed key-value database.

### Why "Distribsilly"?

The name reflects the project's purpose: a *silly* exploration of distributed data storage using basic principles.

---

## Features

1. **Static Sharding**  
   Distribsilly implements static sharding by dividing the hash of a key by the number of shards (`hash(key) % num_shards`). This approach, while simple, is surprisingly robust for certain distributed system use cases.  
   
2. **BoltDB Backend**  
   Each shard uses [BoltDB](https://github.com/etcd-io/bbolt) as its underlying storage mechanism, ensuring a fast, reliable, and embedded key-value store for our database.

3. **Resharding for Powers of Two (Planned)**  
   Resharding is often a complex operation in distributed systems. Distribsilly plans to demonstrate an elegant and efficient resharding technique when the number of shards is a power of two. This approach ensures minimal data movement and straightforward implementation.

---

## Usage

Details for installation and usage will be added as the project progresses.

---

## TODO

- **Implement Resharding**  
  Add support for dynamic resharding when the number of shards changes, particularly focusing on the case where the number of shards is a power of two.

- **Automated Testing**  
  Develop comprehensive test cases.

- **Basic Performance Testing**  
  Benchmark reads and writes to analyze and improve the database's performance. (Details and subparts to be added later.)

- **Implement Replication**  
  Add support for data replication across shards to enhance reliability and fault tolerance. (Details and subparts to be added later.)

---

## Contributing

Contributions to Distribsilly are welcome! Feel free to submit issues or pull requests as we refine and expand the project.

---

Stay tuned for updates as Distribsilly grows from a "silly" project into a useful tool for understanding distributed systems!

