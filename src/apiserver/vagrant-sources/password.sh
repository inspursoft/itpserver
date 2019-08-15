#!/bin/sh
/usr/bin/expect << EOF
spawn sudo passwd root
expect {
        "Enter new UNIX password:" { send "123\r";exp_continue }
        "Retype new UNIX password:" { send "123\r";exp_continue }
        eof { exit }
        }
EOF