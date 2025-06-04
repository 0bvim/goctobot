# GoctoBot

GoctoBot is a command-line interface (CLI) tool to manage your followers and the people you follow on GitHub. Tired of OctoCat hogging the spotlight? GoctoBot is here to help automate some common social interactions on the platform.

This project is based on the original [OctoBot repository](https://github.com/X3ric/octobot/tree/main).

## Features

* **List Counts**: Shows the number of followers you have and the number of users you are following.
* **Follow Back**: Follows users who follow you, but you don't follow back yet.
* **Unfollow**: Unfollows users who don't follow you back.
* **Allow and Deny List**: Allows you to specify users who should never be unfollowed (allow list) or who should never be followed (deny list) through a JSON configuration file.

## Installation

To get started with GoctoBot, you need to clone the repository and build the project.

```bash
# Clone the repository
gh repo clone 0bvim/goctobot

# Or use git clone
git clone git@github.com:0bvim/goctobot.git

# Enter the project directory
cd goctobot

# Compile the executable
make
```


## Configuration

Before using GoctoBot, you need to configure your GitHub Personal Access Token.

1.  Create a file named `.env` in the project's root directory. GoctoBot automatically loads environment variables from this file.
2.  Add your token to the `.env` file as follows:

    ```env
    PERSONAL_GITHUB_TOKEN="your_token_here"
    ```

    GoctoBot requires the `PERSONAL_GITHUB_TOKEN` environment variable to authenticate with the GitHub API. You can generate a new Personal Access Token in your GitHub account's developer settings.

## Usage

The main executable is `goctobot`. Commands are run as follows:

```bash
./goctobot <command>
```


### Available Commands

* `list`: Displays your current follower and following counts.
    ```bash
    ./goctobot list
    ```

* `follow`: Follows all users who follow you but you don't follow back.
    ```bash
    ./goctobot follow
    ```

* `unfollow`: Unfollows users you follow who do not follow you back.
    ```bash
    ./goctobot unfollow
    ```

### Allow and Deny Lists

You can prevent GoctoBot from unfollowing important accounts (like Linus Torvalds) or following unwanted accounts. To do this, create a `userlist.json` file at the following path: `internal/app/model/userlist.json`.

Add the usernames in the following JSON format:

```json
{
  "user1": "Deny",
  "user2": "Allow",
  "user3": "Allow"
}
```


* **Allow**: Users on this list will not be unfollowed by the `unfollow` command, even if they don't follow you back.
* **Deny**: Users on this list will not be followed by the `follow` command.

## Dependency Maintenance

This project uses [Dependabot](https://docs.github.com/en/code-security/dependabot) to keep the Go Modules dependencies up to date. Checks are performed weekly.
