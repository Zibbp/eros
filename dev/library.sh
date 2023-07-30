#!/bin/bash

log_file="/tmp/ceres_$RANDOM.log"  # Please replace this with the correct path
status="success"
hostname=$HOSTNAME  # Get the hostname
script_name=$(basename ${BASH_SOURCE[1]})  # Get the name of the script that sourced the library
add_timestamp=true  # Set this to false if you don't want timestamps in the log file

# Function to be executed on ERR and EXIT signals
function handler() {
    # Check if the script exited due to an error
    if [[ $? -ne 0 ]]; then
        status="failed"
    fi

    # Close the connections to the tee process and the log file
    exec 1>&3 2>&4

    # Send the log file and status to the endpoint
    curl -X POST -H "Content-Type: multipart/form-data" -F "name=$script_name" -F "hostname=$hostname" -F "status=$status" -F file=@$log_file http://localhost:4000/api/v1/report

    # Clean up the log file
    rm "$log_file"
}

# Function to start logging
function start_logging() {
    if $add_timestamp; then
        # Redirect stdout ( > ) and stderr (2>&1) to a file and prepend a timestamp to each line
        exec > >(awk '{ print strftime("[%Y-%m-%d %H:%M:%S]"), $0 }' | tee -a "$log_file") 2>&1
    else
        # Redirect stdout ( > ) and stderr (2>&1) to a file without prepending a timestamp
        exec > >(tee -a "$log_file") 2>&1
    fi

    # Trap ERR and EXIT signals
    trap 'handler' ERR EXIT
}

# Store stdout and stderr
exec 3>&1
exec 4>&2

# Start logging
start_logging