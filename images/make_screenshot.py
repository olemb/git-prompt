#!/usr/bin/env python
from blessings import Terminal

t = Terminal()
G = t.green
R = t.red
Y = t.yellow
N = t.normal
grey = t.grey


script = f"""\


    ~/src/super-ai {G}[main]{N}> git switch -c bugfix
    ~/src/super-ai {G}[bugfix]{N}> vi ai_matrix.js
    ~/src/super-ai {Y}[bugfix *]{N}> git commit -a -m "Matrix was upside down"
    ~/src/super-ai {G}[bugfix]{N}> git switch main
    ~/src/super-ai {G}[main]{N}> git merge bugfix
    ~/src/super-ai {G}[main ↑]{N}> git push
    ~/src/super-ai {G}[main]{N}>


    {G}[main]{N}       at main branch with clean working directory
    {Y}[main ?]{N}     one or more untracked files
    {Y}[main *]{N}     one or more uncommited changes
    {G}[main ↑]{N}     one or more commits ahead of remote
    {R}[main *!]{N}    unmerged conflict
    {G}[bugfix]{N}     at branch called bugfix
    {G}[:initial]{N}   initial branch (nothing commited yet)
    {G}[:f9a820]{N}    specific commit checked out


"""

print(script, end='')
