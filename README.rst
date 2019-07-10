git-prompt - Minimal Git prompt for Bash
========================================

Screenshot
----------




Installing
----------

Just add the Here's what I use::

    source ~/common/conf/git-prompt.sh
    PS1='\u@\h:\w $(git-prompt)$ '


Example Output
--------------

::

    [master]      # At master branch with clean repository.
    [master *]    # Uncommited changes.
    [master ↑]    # One or more commits ahead of remote.
    [master *?]   # Changes and untracked files.
    [master *!↕]  # Unmerged conflict.
    [:initial]    # New git repository with no commits yet.


Branch Name
-----------

::

    master    Branch name.
    :f9a02c   Detached head. (First 6 characters of commit hash.)
    :initial  Initialized repository with nothing commited yet.


Status Flags
^^^^^^^^^^^^

::

    *  Repository has uncommited changes.
    ?  There are untracked files.
    !  There are conflicts.
    ↑  Local repository is ahead of remote by >= 1 commit.
    ↓  Local repository is behind remote by >= 1 commit.
    ↕  Local repository has diverged from remote.
       (It's both ahead and behind.)


Colors
------

Green
    Working directory is clean but may be out of sync with remote.

Yellow
    There are changed or untracked files.

Red
    There are conflicts that need to be resolved.


License
-------

`MIT license <http://en.wikipedia.org/wiki/MIT_License>`_.


Contact
-------

Ole Martin Bjorndalen - ombdalen@gmail.com

