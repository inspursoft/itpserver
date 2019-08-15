#!/bin/sh

USER=root   #username
PASS=123    #password
/usr/bin/expect << EOF
set timeout 5
spawn /usr/bin/ssh-copy-id -i $HOME/.ssh/id_rsa.pub $USER@$1
expect {
        "*yes/no*" { send "yes\r";exp_continue }
        "password:" { send "$PASS\n";exp_continue }
        eof { exit }
        }
EOF
