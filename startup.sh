#!/bin/sh

/adduserserver --cert="false" &
/usr/sbin/sshd -D
