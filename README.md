<div align="center">
<h1>I finished basic program usages.</h1>
<p>I don't use a fork because I can't open issue and for a while I'll need it. </p>
<p>Here is original repository</p>
  
<a href="https://github.com/X3ric/octobot/tree/main" target="_blank" rel="noopener noreferrer">
  Visit OctoBot Repository on GitHub
</a>


# GoctoBot

Tired of OctoCat hogging the spotlight?

### Installation

<pre>
git clone git@github.com:0bvim/goctobot.git
# or
gh repo clone 0bvim/goctobot
cd goctobot
make
</pre>

### Usage

<pre>
./goctobot &lt;command&gt; [username]
</pre>

<pre>
follow [username] — Follow all followers of the specified user.
unfollow — Unfollow who do not follow back.
following — Shows count of users you follow.
followers — Shows count of your followers.
status - Show both, follower and following
</pre>

<br>

### Allow and Deny list.
To add a username to allow or deny list, you should add in json format in file `userlist.json` under the following path:
*`internal/app/model/userlist.json`* marking as `deny` or `allow` like this:

</div>

```json
{
  "user1": "Deny",
  "user2": "Allow",
  "user3": "Allow"
}
```

* Allow list -> When you run unfollow command you can have a file in repository with usernames that you don't want
  to unfollow even if them don't follow you back. Like Torvalds, Thompson and so forth.
* Deny list -> When you run follow [username] command and you don't want to follow someone is just put name in this file too.

<div align="center">

### Coming soon
Command to add user to allow or deny list file.

</div>
