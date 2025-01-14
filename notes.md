# Notes:

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

