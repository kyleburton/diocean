# http://askubuntu.com/questions/95211/how-do-i-set-up-bash-completion-for-command-arguments
function _diocean_completion () {
  COMPREPLY=()
  local cur="${COMP_WORDS[COMP_CWORD]}"
  local prev="${COMP_WORDS[COMP_CWORD-1]}"

  local completions="$(diocean -cmplt ${COMP_LINE})"
  COMPREPLY=( $(compgen -W "$completions" -- "$cur") )
}

complete -F _diocean_completion diocean

