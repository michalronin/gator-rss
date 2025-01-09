**GATOR**
Gator is an RSS feed aggregator, written in Go.

**REQUIREMENTS**
* Gator requires PostgreSQL and Go installed to run
* After installing PostgreSQL and Go, you can clone this repository and install Gator using `go install`
* To create a user, run `gator register <username>`

**OTHER COMMANDS**
* `gator login` - log as a different, already existing user
* `gator reset` - reset the state of the program, clearing all stored data
* `gator users` - list existing users
* `gator addfeed` - add an RSS feed for the currently logged in user
* `gator feeds` - list added RSS feeds
* `gator follow` - follow an RSS feed for the currently logged in user
* `gator following` - list RSS feeds followed by the currently logged in user
* `gator unfollow` - unfollow an RSS feed followed by the currently logged in user
* `gator agg` - aggregate posts from followed feeds
* `gator browse` - browse posts aggregated from followed feeds
