#
# Git Prompt for Fish.
#
#
#
function git_prompt
    set -l text (command git status --porcelain=v2 --branch ^/dev/null)
    if [ $status -ne 0 ] >/dev/null
        return
    end

    set -l branch_oid
    set -l branch_head

    # Flags.
    set -l changes ''
    set -l untracked ''
    set -l conflict ''
    set -l ahead ''
    set -l behind ''

    for line in $text
        set -l words (string split " " $line)

        switch $words[1]
            # Why doesn't this work if the '#' case is last?
            case '#'
                switch $words[2]
                    case 'branch.oid'
                        set oid $words[3]
                    case 'branch.head'
                        set head $words[3]
                    case 'branch.ab'
                        if [ $words[3] != "+0" ]
                            set ahead '↑'
                        end
                        if [ $words[4] != "-0" ]
                            set behind '↓'
                        end
                end

            case 'u'
                set conflict '!'
            case '1' '2'
                set changes '*'
            case '?'
                set untracked '?'
        end
    end


    if [ $oid = "(initial)" ]
        set head ":initial"
    else if [ $head = "(detached)" ]
        set head ':'(string sub -l 6 $oid)
    end


    set -l color green

    if [ $conflict != '' ]
        set color red
    else if [ $changes$untracked != '' ]
        set color yellow
    end

    set -l flags $changes$untracked$conflict$ahead$behind

    if [ $flags != "" ]
       set flags " $flags"
    end

    set_color $color
    echo "[$head$flags]"

    set_color normal

end

git_prompt
