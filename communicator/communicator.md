# communicator

The `Communicator` is the "intermediary" between the `Reader` and the `Migrator`. It is responsible for maintaing persistence, job cleanup, and calling relevant services on different job states.

Any calls back to the Forum will be kept here.
