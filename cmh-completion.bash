#!/bin/bash

_cmh-completion() {
    COMPREPLY=()
    if [ "${#COMP_WORDS[@]}" == "2" ]; then
        COMPREPLY=( $( compgen -W "--args --dry --no-install --prefix --release --static --verbose --version" -- "${COMP_WORDS[COMP_CWORD]}" ))
    fi

    return 0
}

complete -F _cmh-completion cmh
