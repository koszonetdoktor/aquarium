1. Start the conainer
```docker run -p 27017:27017 --name mongodb -v $(pwd)/data/db:/data/db -d mongo``