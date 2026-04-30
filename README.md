# Go-Blog-Aggregator (gator)

A command-line RSS feed aggregator built in Go. Follow your favorite blogs and browse their latest posts directly from the terminal.

## Prerequisites

- [Go](https://golang.org/dl/) (1.21 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

```bash
go install github.com/your-username/Go-Blog-Aggregator@latest

Database Setup
Create a PostgreSQL database for the application:

CREATE DATABASE gator;

Then run the migrations to set up the schema.
```

## Configuration
Go-Blog-Aggregator reads from a config file located at ~/.gatorconfig.json. Create it with the following structure:

{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}

Replace username and password with your PostgreSQL credentials.

## Usage
Register and login
Go-Blog-Aggregator register <username>
Go-Blog-Aggregator login <username>

Add and follow feeds
### Add a new RSS feed
Go-Blog-Aggregator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml

### Follow an existing feed
Go-Blog-Aggregator follow https://blog.boot.dev/index.xml

### See all feeds you follow
Go-Blog-Aggregator following

### Unfollow a feed
Go-Blog-Aggregator unfollow https://blog.boot.dev/index.xml

Aggregate and browse posts
### Start the aggregator (fetches feeds on an interval)
Go-Blog-Aggregator agg 30s

### Browse your latest posts (defaults to 2)
Go-Blog-Aggregator browse

### Browse with a custom limit
Go-Blog-Aggregator browse 10