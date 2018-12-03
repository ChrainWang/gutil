# gutil
Easier-to-use go utils
Encapsulated some of the go utils to make it easier to use

## For example:
With go library:
```
if block, err := aes.NewCipher(key); err == nil {
		blockMode := cipher.NewCBCDecrypter(block, iv)
		plaintext = make([]byte, len(encrypted))
		blockMode.CryptBlocks(plaintext, encrypted)
    ...
    ...
    ...
}
```
With gutil it would be:
```
if decrypted, err := gutil.CBCDecrypt(encrypted, iv, key); err == nil {
  ...
  ...
  ...
}
```

## Installation
```
go get github.com/chrainwang/gutil
```

## Usage
Most of the functions are easy-to-use and easy-to-understand

### Concurrent
This is a util for goroutine management
```
task1 := gutil.ConcurrentTask{
	Do: func() interface{} {
		return "Hello Golang"
	},
}
task2 := gutil.ConcurrentTask{
	Do: func() interface{} {
		return "Hello GUTIL"
	},
}
counter := gutil.Concurrent([]*gutil.ConcurrentTask{&task1, &task2})
fmt.Println(task1.GetValue().(string))
fmt.Println(task2.GetValue().(string))
```
The Do function of each task would be run asynchronously, and the `gutil.Concurrent([]*gutil.ConcurrentTask{&task1, &task2})` would block until both of the tasks are completed.  
The counter would be 2 in this case indicating the number of tasks which are completed  
`task.GetValue()` would return the same value that's returned by `Do` function in interface{} type. You need to convert it into it's original type by yourself
