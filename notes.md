# Notes:

The static sharding and resharding techniques, particularly with the "power of 2 modulo" approach, are commonly used in distributed systems, especially for partitioning large datasets. Here's an overview of their advantages and disadvantages:

### **Advantages of Static Sharding via Power of 2 Modulo Technique**

1. **Simplicity and Predictability**:
   - The formula for determining which shard a piece of data will go to is simple: `hash(data) % (2^k)`, where `k` is the number of shards.
   - This makes the system easy to implement and understand, and shard allocations are deterministic.

2. **Balanced Load**:
   - Since the number of shards is a power of 2, the hash space is evenly divided, which often leads to a balanced distribution of data across shards.
   - This can reduce the risk of certain shards becoming hotspots, especially in a relatively evenly distributed dataset.

3. **Efficient Lookup and Routing**:
   - The modulo operation with a power of 2 is computationally inexpensive and fast. This efficiency makes routing queries to the correct shard relatively quick.

4. **Scalability**:
   - Adding a new shard (i.e., doubling the number of shards) is straightforward in the power of 2 system, and the system can scale by adjusting the modulo to accommodate more shards as needed.

### **Disadvantages of Static Sharding via Power of 2 Modulo Technique**

1. **Resharding Challenges (Data Movement)**:
   - When the number of shards is increased or decreased, a significant amount of data may need to be redistributed. This can cause a large overhead and downtime for resharding, which might lead to performance degradation or increased operational costs.
   - Static sharding doesn't handle dynamic changes in data volume well, and the distribution may become skewed if the data grows unevenly.

2. **Poor Flexibility for Uneven Load**:
   - If the data has an uneven distribution (e.g., some keys are much more frequent than others), the power-of-2 modulo approach may result in poor load balancing. Some shards may receive much more data or queries than others, leading to hotspots.
   - Static sharding requires manual intervention or changes to deal with this type of imbalance.

3. **Complexity with Changing Access Patterns**:
   - In systems where access patterns change frequently (e.g., new types of queries or changes in workload), a static sharding approach may not adapt well. This can make it difficult to optimize resource usage without performing major resharding operations.

4. **Potential Overhead with Larger Power of 2**:
   - As the number of shards grows, maintaining a large number of shards can become operationally complex. Systems with a large number of shards might require additional layers of infrastructure to manage routing and data consistency.

5. **Risk of Underutilization of Resources**:
   - When using a power of 2 modulo approach, if a large number of shards are created but not enough data is available to fill them, some resources (e.g., hardware, storage) may be underutilized.

### **Conclusion**

- **Static sharding with the power of 2 modulo technique** is simple to implement and scales predictably. However, it comes with challenges related to data redistribution (resharding) and potential load imbalances. While it can be efficient for certain types of systems, its limitations around flexibility and data growth mean it may not be suitable for highly dynamic or large-scale distributed systems without additional management techniques like dynamic sharding.

## Sharding:
What is Sharding?
Redistribute your key value pairs into different servers/locations.

Static Sharding:
Data resides in a particular shard and its location doesnt change ever.

Determine which key lives in which shard?
Since this is static sharding we can create a static config file which has all the info 


## What is a Transaction in BoltDB?

In BoltDB, a transaction represents a single, atomic operation or a set of operations on the database. Transactions ensure that:

    Consistency: All changes within a transaction either succeed together or fail together (no partial updates).
    Isolation: Multiple readers can access the database concurrently, but only one writer can modify it at any given time.

This design makes BoltDB fast and reliable while avoiding conflicts between multiple operations.

Why Transactions in BoltDB?

Transactions in BoltDB handle operations in a safe and atomic way. Without them, there’d be a risk of:

    Inconsistent data writes.
    Conflicts between concurrent read and write operations.

Example: db.Update(func(tx *bbolt.Tx) error {...})

Here’s a breakdown of the snippet:

err = db.Update(func(tx *bbolt.Tx) error {
	// Get or create a bucket named "MyBucket"
	_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
	return err // Commit the transaction if no error occurs
})

    db.Update:
        Opens a write transaction.
        The function you pass to db.Update defines all the changes you want to make.

    tx *bbolt.Tx:
        The transaction object, representing your access to the database during the operation.

    Creating a Bucket:
        A bucket is like a table or namespace in SQL databases.
        The tx.CreateBucketIfNotExists method checks if the bucket exists; if not, it creates it.

    Returning an Error:
        If the function returns nil, BoltDB commits the transaction (saving changes).
        If it returns an error, BoltDB rolls back the transaction (discarding changes).

