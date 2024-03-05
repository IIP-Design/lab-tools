# Post Contact DynamoDB Upload

## Background

This script is used to upload mission contact information from a CSV file into a DynamoDB table. It is a companion to the script found in the adjacent directory `usembassy-scraper` and accepts the CSV generated as the output of that script as its input. It was developed in support of the [travel alerts project](https://github.com/IIP-Design/travel-alerts) which required such a centralized list of contact information for all posts.

This simple script completes three steps:

1. Checks whether the required DynamoDB table exists
1. Reads a provided CSV file creating a `PostItem` out of each row in the table
1. Uploads each `PostItem` as an entry in the specified DynamoDB table

Note: While the script is agnostic regarding this point, we recommend that the table uses the `country` value as the hash key and the `name` as the sorting key. This facilitates the common operation of narrowing the list of posts down by country.

## Running the Script

This script is written using the Go programming language. A user with [Go installed and configured](https://go.dev/doc/install) on their system can compile the script by running the `go build` command from within this directory. This will compile the script into a binary file called `uploadcontacts`. The binary can then be executed by running `./uploadcontacts` from within this directory.

Note, you must have valid AWS credentials granting you permission to interact with DynamoDB in order to successfully execute this script.

## Configurations

There are three settings that can be adjusted on this script. These values are defined as constants in the `main.go` file and are as follows:

### AWS_REGION

Unsurprisingly, `AWS_REGION` defines which region the given DynamoDB table is deployed to. Out of the box it defaults to `us-east-1`.

### CSV_FILE

The script expects the presence of a CSV file called `post-contacts.csv` within this directory. This is the file that the scripts reads and converts into items in DynamoDB. This CSV file should contain the four columns `name`, `country`, `website`, and `contact` in that order.

### TABLE

This is the name of the DynamoDB table where the post contact items should be saved.
