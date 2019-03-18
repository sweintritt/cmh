#!/bin/bash

_cmh-completion() {
    COMPREPLY=()
    COMPREPLY=( $( compgen -W "--args --dry --no-install --prefix --release --static --verbose --version" -- "${COMP_WORDS[COMP_CWORD]}" ))

    return 0
}

complete -F _cmh-completion cmh
