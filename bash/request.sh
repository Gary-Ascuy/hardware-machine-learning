#!/bin/bash

#
# Source - https://github.com/sararob/ml-talk-demos/blob/master/speech/request.sh
# @sararob - Sara Robinson
# Updated by Gary Ascuy
#

# Record the Audio
read -p "Press [Enter] to start recording the audio..."
sox -d --channels=1 --bits=16 --rate=16000 audio.flac trim 0 5

# Create request JSON
FILENAME="request-"`date +"%s".json`
AUDIO_CONTENT_BASE64=`base64 audio.flac`
cat <<EOF > $FILENAME
{
  "config": {
    "encoding":"FLAC",
    "sampleRateHertz":16000,
    "profanityFilter": true,
    "languageCode": "en-US",
    "speechContexts": {
      "phrases": ['']
    },
    "maxAlternatives": 1
  },
  "audio": {
    "content": "$AUDIO_CONTENT_BASE64"
	}
}
EOF

# Call to Google API
curl -s -X POST -H "Content-Type: application/json" --data-binary @${FILENAME} https://speech.googleapis.com/v1/speech:recognize?key=YOUR_API_KEY
