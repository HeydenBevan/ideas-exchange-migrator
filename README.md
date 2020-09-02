# ideas-exchange-migrator
Migration Utility to get Ideas out of the Community Forum

# Current Design:
Everything here is still "work in progress" and no decisions have been made. Will be spit balling everything from here on out. You have been warned.

Will be looking at having 3 services that work in tandem:
1. The `Reader`
2. The `Migrator`
3. The `Communicator`

There will need to be a persistence Layer between them so that the `Reader` can pass Ideas into a "queue" for the `Migrator` to start working. Thinking an SQS layer.

Current idea is to:
1. Map out the Idea Exchange data we will need.
2. The `Reader` will make a `GET` request to the Khoros Community API to fetch Ideas.
3. Transform data into a "Universal" format
4. Call the `Communicator` server to start a job
5. `Communicator` server will add it into a Queue and return a `201` to the `Reader`
6. Every 2 hours, the `Migrator` will check the queue for items to migrate
7. It will transform the data into the `DTO` requirements
8. `Migrator` will then call the destination service to take in the new items
9. `Migrator` will then call the `Communicator` to alert that the Job is done
10. `Communicator` will then handle communications on the Forum (IE: Post a Status Update on the Migrated Items on the Khoros Community Exchange)

# Tech Decisions
Will be using [Go (Golang)](https://golang.org/) as the tech stack of choice, the reasons being are:
1. It's lightweight.
2. Easy to learn.
3. Has less reliance on 3rd party packages, strongly advise using the stdlib wherever possible.
4. It's my favourite language and the one I'm most comfortable, should I lose innovation tokens for JS? No way.

We like the `GOPATH` in this house, take a look if you're not familiar: [How to Write Go Code (with GOPATH)](https://golang.org/doc/gopath_code.html)

# Repo structure!
| Directory | Sub Dir | Purpose |
| --- | --- | --- |
| `common` | n/a | Holds the packages that can be used by all 3 services |
| `communicator` | n/a | The Communicator server that manages all comms logic |
| `migrator` | n/a | The Migrator service that handles the DTO's |
| `reader` | n/a | The Reader client that will be talking to the Forum |
