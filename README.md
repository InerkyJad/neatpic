```
         /$$   /$$                       /$$     /$$$$$$$  /$$          
        | $$$ | $$                      | $$    | $$__  $$|__/          
        | $$$$| $$  /$$$$$$   /$$$$$$  /$$$$$$  | $$  \ $$ /$$  /$$$$$$$
        | $$ $$ $$ /$$__  $$ |____  $$|_  $$_/  | $$$$$$$/| $$ /$$_____/
        | $$  $$$$| $$$$$$$$  /$$$$$$$  | $$    | $$____/ | $$| $$      
        | $$\  $$$| $$_____/ /$$__  $$  | $$ /$$| $$      | $$| $$      
        | $$ \  $$|  $$$$$$$|  $$$$$$$  |  $$$$/| $$      | $$|  $$$$$$$
        |__/  \__/ \_______/ \_______/   \___/  |__/      |__/ \_______/
```

# NeatPic
NeatPic is a CLI tool that allows you to manipulate images using a variety of commands. With NeatPic, you can blur, change the brightness, contrast, gamma, and saturation of an image, as well as flip, grayscale, invert, reduce, resize, restore, rotate, and sharpen it.

# Usage
To use NeatPic, simply run `neatpic [command]` followed
by the appropriate flags and arguments for the command you want to use.
For example, to blur an image, you would run `neatpic blur <image_path> [flags]`. To see the full list of available commands and their flags, run `neatpic --help`.


# Building
```bash
$ git clone https://github.com/InerkyJad/neatpic.git
$ cd neatpic
$ go build
```


# Available Commands (table)
| Command    | Description                                      | Flags                         |
|------------|--------------------------------------------------|-------------------------------|
| blur       | Blur the image                                   | --value -v                    |
| brightness | Change the brightness of the image               | --value -v                    |
| completion | Generate the autocompletion script for the shell | none                          |
| contrast   | Change the contrast of the image                 | none                          |
| corp       | Corp the image                                   | --width -W --height -H        |
| flip       | Flip the image                                   | --horizontal -h --vertical -v |
| gamma      | Change the gamma of the image                    | --value -v                    |
| grayscale  | Grayscale the image                              | none                          |
| help       | Help about any command                           | --help -h                     |
| invert     | Invert the image                                 | none                          |
| reduce     | Reduce the image size                            | --value -v                    |
| resize     | Resize the image                                 | --width -W --height -H        |
| restore    | Restore the image                                | none                          |
| rotate     | Rotate the image                                 | --angle -a                    |
| saturation | Change the saturation of the image               | --value -v                    |
| sharpen    | Sharpen the image                                | --value -v                    |




# License
NeatPic is released under the BSD-3-Clause License. See the (LICENSE)[LICENSE] file for more information.