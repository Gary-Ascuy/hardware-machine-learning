#!/bin/bash

#
# Source - https://github.com/sararob/ml-talk-demos/blob/master/speech/request.sh
# @sararob - Sara Robinson
# Updated by Gary Ascuy
# 
# Create Create Google API KEY
# https://support.google.com/googleapi/answer/6158862?hl=en
# !do not forget setup GOOGLE_API_KEY env variable  
#
# Install Sox (Sound eXchange) & Manual
# http://sox.sourceforge.net/
# http://www.thegeekstuff.com/2009/05/sound-exchange-sox-15-examples-to-manipulate-audio-files
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
    "languageCode": "es-BO",
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
curl -s -X POST -H "Content-Type: application/json" --data-binary @${FILENAME} https://speech.googleapis.com/v1/speech:recognize?key=$GOOGLE_API_KEY
