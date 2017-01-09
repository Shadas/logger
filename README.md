# logger
It's a tool for free log by golang.



### How to use

1. Run the order 

    ```
    go get "github.com/Shadas/logger"
    ```

2. Create a config object for drive the logger

   ```
   loggerConfig := logger.LoggerConfig{}
   ```

3. The contents of  LoggerConfig are as shown in config.sample.json

4. Add the function InitLogger to finish the configure of logger.

   ```
   err := logger.InitLogger(loggerConfig)
   if err != nil {
      log.Println(err.Error())
   }

   ...(business code)
   ```

5. Then you can use it by 

	```
    logger.Log(logger.DEBUG, "debug thing1")

	  logger.Log(logger.DEBUG, "debug thing2")

    logger.Log(logger.ERROR, "error thing1")

    logger.Log(logger.INFO, "info thing1")
    ```

6. Do not forget to close the logger in the end.

   ```
   ...
   defer logger.CloseLogger()
   ```



### Features

1. **Log classify.** There are 5 kinds of log can be chosen. They are DEBUG, INFO, WARNING, ERROR, FATAL low-to-high. If you print a high class log, it will be print in the lower logs at the same time.

   The different log message will be written in different files.

2. **Asynchronous write.** The messages will be put in a channel first, one other thread will get it from the channel and write to the file. 

3. **Full write.** If there are still some messages in the channel while the main thread is over, the CloseLogger function will ensure the messages can be written to the file before the program exit.