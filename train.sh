#!/bin/bash
#
# Train an audio classifier by piping samples.go
# into this script.

if [ ! -f out_net ]; then
  echo 'Creating network...'
  neurocli new -in net_classifier.txt -out classifier
  neurocli new -in net_rnn.txt -out rnn
  neurocli seq2vec -rnn rnn -outnet classifier -out out_net
  rm rnn classifier
fi

neurocli train \
  -net out_net \
  -cost sigmoidce \
  -batch ${BATCH_SIZE:-8} \
  -adam default \
  -step ${STEP_SIZE:-0.001}
