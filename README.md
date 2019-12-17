# sentimentd

High-performance sentiment analyze service

## Features

* Multi-label classification 

# Operations

## Build binary

```shell script
make build
```

## Train

Firstly you need to create the brain (classifier) for the training.
```shell script
./build/sentimentd brain create skynet "The artificial neural network-based conscious group mind"
```

```shell script
./build/sentimentd train
```

### Apply migrations

```shell script
./build/sentimentd migrate
```

# Credits

Many thanks to [kaggle](kaggle.com) for datasets.

See [Sentiment Analysis: Emotion in Text](https://www.kaggle.com/c/sa-emotions/data).

# TODO

Calculate a time and space complexity 