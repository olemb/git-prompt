#!/usr/bin/env python3
"""
git-prompt

Ole Martin Bjorndalen
https://github.com/olemb/git-prompt

License: MIT
"""
#
# To use just add $(git-prompt.py) to PS1, for example.
# 
#    PS1='\u@\h:\w $(git-prompt.py)$ '
#

import os

green = '92'
yellow = '93'
red = '31'


def get_status():
    return list(os.popen('git status --porcelain=v2 --branch 2>/dev/null'))


def parse_status(lines):
    oid = ''
    head = ''
    status = {'ahead': False,
              'behind': False,
              'untracked': False,
              'conflict': False,
              'changed': False}

    for line in lines:
        words = line.split()
        char = line[:1]

        if char == '#':
            _, name, *args = line.split()
            if name == 'branch.oid':
                oid = args[0]
            elif name == 'branch.head':
                head = args[0]
            elif name == 'branch.ab':
                status['ahead'] = args[0] != '+0'
                status['behind'] = args[1] != '-0'
        elif char == '?':
            status['untracked'] = True
        elif char == 'u':
            status['conflict'] = True
        elif char.isdigit():
            status['changed'] = True
 
    if oid == '(initial)':
        branch = 'initial'
    elif head == '(detached)':
        branch = oid[:6]
    else:
        branch = head

    return branch, status


def format_status(branch, status):
    flags = ''
    color = green

    if status['changed']:
        flags += '*'
        color = yellow
    
    if status['untracked']:
        flags += '?'
        color = yellow

    if status['conflict']:
        flags += '!'
        color = red

    if status['ahead'] and status['behind']:
        flags += '↕'
    elif status['ahead']:
        flags += '↑'
    elif status['behind']:
        flags += '↓'

    if flags != '':
        flags = ' ' + flags

    return f'\033[{color}m[{branch}{flags}]\033[0m'


lines = get_status()
if lines != []:
    branch, status = parse_status(lines)
    print(format_status(branch, status), end='')
