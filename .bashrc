alias pm='go build -ldflags="-w -s" && upx -9 go_psum &&  cp ./go_psum ~/soft && mv ~/soft/go_psum ~/soft/pstat'
alias run='sudo ./go_psum --name=nginx --show=0'

alias prun='pm && run'
source ~/envs/default/bin/activate