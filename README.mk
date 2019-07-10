# git-prompt - Minimal Git prompt for Bash

## Screenshot


## Installing


Here's what I use:

```bash
source ~/common/conf/git-prompt.sh
PS1='\u@\h:\w $(git-prompt)$ '
```


## Example Output

```
[master]      # At master branch with clean repository.
[master *]    # Uncommited changes.
[master ↑]    # One or more commits ahead of remote.
[master *?]   # Changes and untracked files.
[master *!↕]  # Unmerged conflict.
[:initial]    # New git repository with no commits yet.
```


## Branch Name

```
master    Branch name.
:f9a02c   Detached head. (First 6 characters of commit hash.)
:initial  Initialized repository with nothing commited yet.
```


## Status Flags

```
*  Repository has uncommited changes.
?  There are untracked files.
!  There are conflicts.
↑  Local repository is ahead of remote by >= 1 commit.
↓  Local repository is behind remote by >= 1 commit.
↕  Local repository has diverged from remote.
   (It's both ahead and behind.)
```


## Colors

*Red:* There are conflicts that need to be resolved.

*Yellow:* There are changed or untracked files.

*Green:* Working directory is clean.


# License

[MIT license](http://en.wikipedia.org/wiki/MIT_License)


# Author

Ole Martin Bjørndalen
