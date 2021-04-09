#!/bin/bash

echo "Which site would you like to sync certificates for?"

SITE_OPTIONS=("lab" "content")

# Present the user with a list of containers they can add the certs to
PS3="Pick an option: "
select OPT in "${SITE_OPTIONS[@]}" "quit";
do 
  SELECT_MESSAGE="Checking if the $OPT site exists and is running..."

  case "$REPLY" in
    1) echo "$SELECT_MESSAGE"; break;;
    2) echo "$SELECT_MESSAGE"; break ;;
    $((${#SITE_OPTIONS[@]}+1))) echo "Goodbye!"; break;;
    *) echo "Invalid option. Try another one."; continue;;
  esac
done

# Handle user's selection
if [ $OPT == "quit" ]; then
  exit 0
else
  # Check whether the required Docker container exists.
  EXISTS=$(docker ps -q -f name=${OPT}_web)

  # If the required Docker container does not exist exit early.
  if [ -z "$EXISTS" ]; then
    echo "The container ${OPT}_web does not exist. Goodbye!"
    exit 0
  fi

  # Check whether the required Docker container is running.
  RUNNING=$(docker container inspect -f '{{.State.Running}}' ${OPT}_web)

  # If the required Docker container is not running exit early.
  if [ $RUNNING == "true" ]; then
    echo "The container ${OPT}_web exists and is running!"
  else
    echo "The container ${OPT}_web exists but is not running, please start it and try again."
    exit 0
  fi

fi

# Set variables for certificate stored and target location
CERT_PATH=config/caddy-data/caddy/pki/authorities/local/
CA_PATH=/etc/pki/ca-trust/source/anchors/

# Copy Caddy-generated certificates into the Docker container
for CERT in root intermediate
do
  docker cp ${CERT_PATH}${CERT}.crt ${OPT}_web:/${CA_PATH}/${CERT}.crt
  docker exec -d ${OPT}_web chmod 644 /${CA_PATH}/${CERT}.crt
done

# Add local certificates to the trusted CA roots bundle
docker exec -d ${OPT}_web update-ca-trust