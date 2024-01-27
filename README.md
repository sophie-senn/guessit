# guessit

Design:

```
cmd
    main.go
api
    client.go
        Post guessing to remote server (other player)
        Post update of the word to remote server (other player)
        Get word from the remote server (other player)
        Get a random word from a remote dictionnary
    server.go
        Retrieve guessing from remote client (other player)
    utils.go
    const.go
game
    hmi.go
        Retrieve word from backend (guess.go) and show updated word and image
        Display prompt to catch keyboard inputs and push inputs to backend
    guess.go
        Retrieve keyboard inputs from the player guessing
        Retrieve keyboard inputs from the player proposing the word
        Update word using client/server exchanges and push to HMI
    utils.go
    const.go
```