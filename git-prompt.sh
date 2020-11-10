# Git prompt for Bash
# 
# To use add to (or source into) your bash config and add $(git-prompt)
# to PS1, for example.
# 
#     PS1='\u@\h:\w $(git-prompt.py)$ '
# 
# Ole Martin Bjorndalen
# https://github.com/olemb/git-prompt
# 
# License: MIT

function git-prompt() {   
    local branch=''
 
    local oid=''
    local head=''

    local ahead=0
    local behind=0
    local untracked=0
    local conflict=0
    local changed=0


    # Get data.

    # The 'local' statement needs to be on its own line since or it
    # will overwrite $?.
    local status_text=''
    status_text=$(git status --porcelain=v2 --branch 2>/dev/null)
    if [ ! $? -eq 0 ]
    then
        # Not a Git repository (or some error occured).
        return
    fi


    # Parse data.

    local IFS=$'\n'
    for line in $status_text
    do
        if [[ $line =~ ^#[[:space:]]branch\.([a-z.]+)[[:space:]](.+)$ ]];
        then
            local name="${BASH_REMATCH[1]}"
            local value="${BASH_REMATCH[2]}"

            case $name in
                oid)
                    oid=$value
                    ;;
                head)
                    head=$value
                    ;;
                ab)
                    if [[ $line =~ \+[1-9] ]];
                    then
                        ahead=1
                    fi

                    if [[ $line =~ \-[1-9] ]];
                    then
                        behind=1
                    fi
                    ;;
            esac
        else
            case ${line:0:1} in
                \?)
                    untracked=1
                    ;;
                u)
                    conflict=1
                    ;;
                [1-2])
                    changed=1
                    ;;
            esac
        fi
    done


    # Get branch.

    if [ $oid == "(initial)" ]
    then
        branch=":initial"
    elif [ $head == "(detached)" ]
    then
        branch=":"${oid:0:6}
    else
        branch=$head
    fi


    # Add flags.

    local flags=
    local color="green"

    if [ $changed -eq 1 ]
    then
        # If this is '*' it will expand to filenames.
        flags='*'
        color="yellow"
    fi

    if [ $untracked -eq 1 ]
    then
        flags="$flags?"
        color="yellow"
    fi

    if [ $conflict -eq 1 ]
    then
        flags="$flags!"
        color="red"
    fi

    case $ahead$behind in
        11)
            flags="$flags↕"
            ;;
        10)
            flags="$flags↑"
            ;;
        01)
            flags="$flags↓"
            ;;
    esac


    # Build.
    
    if [ -z "$flags" ]
    then
        local text="[$branch$flags]"
    else
        local text="[$branch $flags]"
    fi


    # Add colors.

    case $color in
        green)
            local colorcode="92"
            ;;
        yellow)
            local colorcode="93"
            ;;
        red)
            local colorcode="31"
            ;;
    esac
    echo -ne "\001\e[${colorcode}m\002${text}\001\e[0m\002"
}
