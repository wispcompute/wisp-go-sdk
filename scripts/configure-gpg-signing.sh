#!/bin/bash
# Create key definition file for gpg batch operation
cat >gpg-key-definition <<EOF
     %echo Generating a RSA key
     Key-Type: RSA
     Key-Length: 4096
     Name-Real: github-gpg-key
     Name-Comment: GPG key used for signing commits with GitHub
     Name-Email: $(git config user.email)
     Expire-Date: 0
     Passphrase: $1
     # Do a commit here, so that we can later print "done" :-)
     %commit
     %echo done
EOF

# Generate a new GPG key
gpg --batch --generate-key gpg-key-definition

# Delete key definition file
rm gpg-key-definition

# Fetch long form ID of GPG key
GPG_LONGFORM_KEY_ID=$(gpg --list-signatures --with-colons | grep 'sig:' | grep 'github-gpg-key' | head -n 1 | cut -d':' -f5)

# Enable gpg signing for repository
git config commit.gpgsign true

# Set the git signing key using the long form id
git config user.signingkey $GPG_LONGFORM_KEY_ID

# Set the IFS variable to an empty string to tell the shell not to perform splitting.
OLD_IFS="$IFS"
IFS=""

# Export public key based on longform id
GPG_PUBLIC_KEY="$(gpg --armor --export $GPG_LONGFORM_KEY_ID)"

# Parse the POST payload for the REST call to GitHub
GITHUB_REST_API_GPG_PAYLOAD=$(echo $GPG_PUBLIC_KEY | sed -E ':a;N;$!ba;s/\r{0,1}\n/\\n/g')

# Restore default shell splitting behaviour
IFS="$OLD_IFS"

# Compile JSON payload
GITHUB_REST_API_JSON_PAYLOAD="{\"name\":\"GitHub GPG Key\",\"armored_public_key\":\"$GITHUB_REST_API_GPG_PAYLOAD\"}"

echo $GITHUB_REST_API_JSON_PAYLOAD

# Add GPG key to account associated with GITHUB_TOKEN (classic or PAT)
curl -L --insecure \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer $2" \
  https://api.github.com/user/gpg_keys \
  -d "$GITHUB_REST_API_JSON_PAYLOAD"

# Clear history to purge passphrase and PAT
history -c && rm ~/.bash_history