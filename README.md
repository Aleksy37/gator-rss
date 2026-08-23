# Gator-RSS -  A CLI RSS feed aggregator in Go

-   

## Dependencies

-  Go v1.26.2 or later
-  Postgres v15 or later

### Install

Once you have the required dependencies installed:
 `go install github.com/Aleksy37/gator-rss@latest`
 paste the above command into your terminal to install gator-rss

#### Setup

Once youve installed gator- rss youll need to create a config file named .gatorconfig.json with this structure in your home directory

`
{
  "db_url": "",
  "current_user_name": "",
  "current_user_id": ""
}
`

the db_url format is `protocol://username:password@host:port/database?sslmode=disable`

then connect to your postgres instance and create a database called gator, once your database is setup and the config is configured properly use the command `gator-rss migrate` to get the database in the correct state to begin using gator.


#### Usage

Some Commands you can run in the gator-rss CLI:

- **migrate**: this command only needs to be ran when you first install/update gator-rss, it creates the tables we need in our database for the program to work
-  **register**: *usage: register username* create a user 
- **login**: *usage: login username* login as that user
- **reset**: Clears the user list 
- **users**: lists users
- **addfeed** *usage: addfeed name url* register a new feed and add it to the current users list of follows
- **feeds**: shows a list of tracked feeds with name and url and the user who first registered them
- **follow**: *usage: follow <url>* if a feed is already being tracked but you arent following it on the current user use this command to add it to your follow list
- **unfollow**: *usage: unfollow url* stop following a feed 
- **following**: shows the names of the feed the current user is following
- **agg**: *usage: agg interval(eg. 5m15s)* this command polls the feeds in a loop on a user specified interval and should be run in a second terminal window in the background **DO NOT DDOS THE RSS FEEDS**
- **browse** *usage: browse limit* returns the latest posts scrapped by the agg command, user specifies how many posts to return 