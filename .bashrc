alias pm='go build && cp ./go_psum ~/soft && mv ~/soft/go_psum ~/soft/pstat'
alias run='sudo ./go_psum --name=nginx --show=0'

alias prun='pm && run'