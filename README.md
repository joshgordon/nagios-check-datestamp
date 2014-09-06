# Nagios Check Datestamp
This is a nagios plugin I wrote to check the datestamp of a file (written in seconds since unix epoch, with a newline at the end) and make sure it's reasonably close to the current time. (Warning and critical levels are specified on the command line.) 

It's written in go, which means that a binary can be compiled and the SUID bit set. 

Usage: 
    check_datestamp /path/to/file <warning seconds> <critical seconds> 

For instance: 
    check_datestamp /var/local/weekly_backup_success.txt 605000 1210000

From within nagios: 
    check_datestamp!/var/local/weekly_backup_success.txt!605000!1210000

will set warning if the datestamp is just over a week old, and critical if the datestamp is a bit over 2 weeks old. 

The use case for this is to monitor some backup scripts running on a server - to make sure that they're running when they're supposed to, and to make sure we're not going too long without a backup. 

## Example Output
    OK - 49103 seconds (0 days, 13:38:23) since last run 
    WARNING - 135457 seconds (1 days, 13:37:37) since last run 
    CRITICAL - 399477 seconds (4 days, 14:57:57) since last run 

# Installation
Copy a binary of `check_datestamp` (appropriate for your platform) to `/usr/lib/nagios/plugins/check_datestamp` Copy `datestamp.cfg` to your plugins config directory (in my instance, it's `/etc/nagios-plugins/config/`), restart nagios, and set up the files that need monitoring as services. 

## NRPE
I have had success using a variant of this (that I wrote for checking puppet's last run time) through NRPE. I'll leave this part as an excersie for the reader. 
