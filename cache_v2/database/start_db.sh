#!/bin/sh

service mysql start

mysql -uroot mysql -e "update user set plugin='mysql_native_password' where User='root'"

mysql -uroot mysql -e "FLUSH PRIVILEGES"

mysql -uroot mysql -e "EXIT"