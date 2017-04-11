# audioclass

This is a test to see how well an RNN can classify sounds based on raw PCM waveforms. I will be using the [audioset](https://github.com/unixpickle/audioset) package for labelled training and evaluation data.

# Initial results

Initially, the model overfit horribly and still did horribly on the training set.

First, two baselines. With constant network outputs, the baseline cost on the training set was 13.2. On the evaluation set, it was 13.9.

After 93K iterations of training, the training cost got down to an average of 9.83. However, the same model on the evaluation set had an average cost of 14.8, which was much worse than the baseline. Thus, training took an incredibly long time and was unfruitful.
