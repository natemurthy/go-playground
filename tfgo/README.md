# tfgo

Sandbox for Tensorflow Go

## Installation

https://github.com/galeone/tfgo#installation

```
go get github.com/galeone/tfgo
```

https://github.com/galeone/tfgo#tensorflow-installation

https://www.tensorflow.org/install/lang_c#macos

Currently installed:

```
curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.11.0.tar.gz" | sudo tar -C /usr/local -xz

sudo update_dyld_shared_cache

nate:local$ pwd
/usr/local

nate:local$ find . -type d -name "*tensor*"
./include/tensorflow
```

Blocker: https://github.com/galeone/tfgo/issues/79
